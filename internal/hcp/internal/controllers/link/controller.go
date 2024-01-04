// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package link

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/consul/agent/hcp/bootstrap"
	"github.com/hashicorp/consul/agent/hcp/client"
	"github.com/hashicorp/consul/agent/hcp/config"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/consul/internal/controller"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
)

type Dependencies struct {
	ResourceApisEnabled              bool
	OverrideResourceApisEnabledCheck bool
	DataDir                          string

	testHCPClient client.Client
}

func LinkController(deps Dependencies) *controller.Controller {
	return controller.NewController("link", pbhcp.LinkType).
		WithReconciler(&linkReconciler{
			resourceApisEnabled:              deps.ResourceApisEnabled,
			overrideResourceApisEnabledCheck: deps.OverrideResourceApisEnabledCheck,
			dataDir:                          deps.DataDir,
			testHCPClient:                    deps.testHCPClient,
		})
}

type linkReconciler struct {
	resourceApisEnabled              bool
	overrideResourceApisEnabledCheck bool
	dataDir                          string

	testHCPClient client.Client
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

	// Check if linking feature is disabled
	if r.resourceApisEnabled && !r.overrideResourceApisEnabledCheck {
		disabledStatus := &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions:         []*pbresource.Condition{ConditionDisabled},
		}
		updateStatus(ctx, rt, res, disabledStatus)
		return nil
	}

	// Configure an HCP client from the resource
	var hcpClient client.Client
	if r.testHCPClient != nil {
		hcpClient = r.testHCPClient
	} else {
		var err error
		hcpClient, err = client.NewClient(config.CloudConfig{
			ResourceID:   link.ResourceId,
			ClientID:     link.ClientId,
			ClientSecret: link.ClientSecret,
		})
		if err != nil {
			rt.Logger.Error("error creating HCP client", "error", err)
			newStatus := &pbresource.Status{
				ObservedGeneration: res.Generation,
				Conditions:         []*pbresource.Condition{ConditionFailed},
			}
			updateStatus(ctx, rt, res, newStatus)
			return nil // return nil to avoid continuous requeueing
		}
	}

	// Load the management token
	tokenCtx, tokenCancel := context.WithTimeout(ctx, 5*time.Minute)
	defer tokenCancel()
	_, err = bootstrap.LoadManagementToken(tokenCtx, rt.Logger, hcpClient, r.dataDir)
	if err != nil {
		var conditions []*pbresource.Condition
		switch {
		case errors.Is(err, bootstrap.ErrFetchBootstrapUnauthorized):
			conditions = append(conditions, ConditionUnauthorized)
		case errors.Is(err, bootstrap.ErrFetchBootstrapForbidden):
			conditions = append(conditions, ConditionForbidden)
		default:
			conditions = append(conditions, ConditionFailed)
		}
		newStatus := &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions:         conditions,
		}
		updateStatus(ctx, rt, res, newStatus)
		return nil // return nil to avoid continuous requeueing
	}

	successStatus := &pbresource.Status{
		ObservedGeneration: res.Generation,
		Conditions:         []*pbresource.Condition{ConditionLinked(link.ResourceId)},
	}
	updateStatus(ctx, rt, res, successStatus)

	return nil
}

func updateStatus(ctx context.Context, rt controller.Runtime, res *pbresource.Resource, newStatus *pbresource.Status) {
	if resource.EqualStatus(res.Status[StatusKey], newStatus, false) {
		return
	}
	_, err := rt.Client.WriteStatus(ctx, &pbresource.WriteStatusRequest{
		Id:     res.Id,
		Key:    StatusKey,
		Status: newStatus,
	})
	if err != nil {
		rt.Logger.Error("error updating link resource status", "error", err)
		return
	}
}
