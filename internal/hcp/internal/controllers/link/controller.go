// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package link

import (
	"context"
	"strings"

	"github.com/hashicorp/consul/agent/hcp/config"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/storage"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/hcp/internal/types"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
)

const (
	MetadataSourceKey    = "source"
	MetadataSourceConfig = "config"
)

func LinkController(resourceApisEnabled bool, overrideResourceApisEnabledCheck bool, cfg config.CloudConfig) *controller.Controller {
	return controller.NewController("link", pbhcp.LinkType).
		WithReconciler(&linkReconciler{
			resourceApisEnabled:              resourceApisEnabled,
			overrideResourceApisEnabledCheck: overrideResourceApisEnabledCheck,
		}).
		WithInitializer(&linkInitializer{
			cloudConfig: cfg,
		})
}

type linkReconciler struct {
	resourceApisEnabled              bool
	overrideResourceApisEnabledCheck bool
}

func (r *linkReconciler) Reconcile(ctx context.Context, rt controller.Runtime, req controller.Request) error {
	// The runtime is passed by value so replacing it here for the remainder of this
	// reconciliation request processing will not affect future invocations.
	rt.Logger = rt.Logger.With("resource-id", req.ID, "controller", StatusKey)

	rt.Logger.Trace("reconciling link")

	rsp, err := rt.Client.Read(ctx, &pbresource.ReadRequest{Id: req.ID})
	switch {
	case status.Code(err) == codes.NotFound:
		rt.Logger.Trace("link has been deleted")
		return nil
	case err != nil:
		rt.Logger.Error("the resource service has returned an unexpected error", "error", err)
		return err
	}

	res := rsp.Resource
	var link pbhcp.Link
	if err := res.Data.UnmarshalTo(&link); err != nil {
		rt.Logger.Error("error unmarshalling link data", "error", err)
		return err
	}

	var newStatus *pbresource.Status
	if r.resourceApisEnabled && !r.overrideResourceApisEnabledCheck {
		newStatus = &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions:         []*pbresource.Condition{ConditionDisabled},
		}
	} else {
		newStatus = &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions:         []*pbresource.Condition{ConditionLinked(link.ResourceId)},
		}
	}

	if resource.EqualStatus(res.Status[StatusKey], newStatus, false) {
		return nil
	}
	_, err = rt.Client.WriteStatus(ctx, &pbresource.WriteStatusRequest{
		Id:     res.Id,
		Key:    StatusKey,
		Status: newStatus,
	})

	if err != nil {
		return err
	}

	return nil
}

type linkInitializer struct {
	cloudConfig config.CloudConfig
}

func (i *linkInitializer) Initialize(ctx context.Context, rt controller.Runtime) error {
	if !i.cloudConfig.IsConfigured() {
		return nil
	}

	// Construct a link resource to reflect the configuration
	data, err := anypb.New(&pbhcp.Link{
		ResourceId:   i.cloudConfig.ResourceID,
		ClientId:     i.cloudConfig.ClientID,
		ClientSecret: i.cloudConfig.ClientSecret,
	})
	if err != nil {
		return err
	}

	// Create the link resource for a configuration-based link
	_, err = rt.Client.Write(ctx,
		&pbresource.WriteRequest{
			Resource: &pbresource.Resource{
				Id: &pbresource.ID{
					Name: types.LinkName,
					Type: pbhcp.LinkType,
				},
				Metadata: map[string]string{
					MetadataSourceKey: MetadataSourceConfig,
				},
				Data: data,
			},
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), storage.ErrWrongUid.Error()) {
			// Ignore wrong UID errors, which indicates the link already exists
			return nil
		}
		return err
	}

	return nil
}
