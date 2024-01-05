// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package link

import (
	"context"
	"fmt"
	hcpclient "github.com/hashicorp/consul/agent/hcp/client"
	"github.com/hashicorp/consul/agent/hcp/config"

	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/consul/internal/controller"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
)

type HCPClientFn func(resourceID, clientID, clientSecret string) (hcpclient.Client, error)

func NewHCPClient(resourceID, clientID, clientSecret string) (hcpclient.Client, error) {
	hcpClient, err := hcpclient.NewClient(config.CloudConfig{
		ResourceID:   resourceID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthURL:      "https://auth.idp.hcp.dev",
		Hostname:     "api.hcp.dev",
	})
	if err != nil {
		return nil, err
	}
	return hcpClient, nil
}

func LinkController(resourceApisEnabled bool, overrideResourceApisEnabledCheck bool, hcpClientFn HCPClientFn) *controller.Controller {
	return controller.NewController("link", pbhcp.LinkType).
		WithReconciler(&linkReconciler{
			resourceApisEnabled:              resourceApisEnabled,
			overrideResourceApisEnabledCheck: overrideResourceApisEnabledCheck,
			hcpClientFn:                      hcpClientFn,
		})
}

type linkReconciler struct {
	resourceApisEnabled              bool
	overrideResourceApisEnabledCheck bool
	hcpClientFn                      HCPClientFn
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

	hcpClient, err := r.hcpClientFn(link.ResourceId, link.ClientId, link.ClientSecret)
	if err != nil {
		rt.Logger.Error("error creating HCP Client", "error", err)
		return err
	}

	cluster, err := hcpClient.GetCluster(ctx)
	fmt.Println("cluster:", cluster)

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
