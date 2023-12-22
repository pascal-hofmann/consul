package telemetrystate

import (
	"context"

	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/resource"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TelemetryStateController(resourceApisEnabled bool, overrideResourceApisEnabledCheck bool) *controller.Controller {
	return controller.NewController("telemetrystate", pbhcp.TelemetryStateType).
		WithReconciler(&stateReconciler{
			resourceApisEnabled:              resourceApisEnabled,
			overrideResourceApisEnabledCheck: overrideResourceApisEnabledCheck,
		})
}

type stateReconciler struct {
	resourceApisEnabled              bool
	overrideResourceApisEnabledCheck bool
}

func (r *stateReconciler) Reconcile(ctx context.Context, rt controller.Runtime, req controller.Request) error {
	rt.Logger = rt.Logger.With("resource-id", req.ID, "controller", "consul.io/hcp-telemetry-state")

	rt.Logger.Trace("reconciling telemetry")

	rsp, err := rt.Client.Read(ctx, &pbresource.ReadRequest{Id: req.ID})
	switch {
	case status.Code(err) == codes.NotFound:
		rt.Logger.Trace("telemetry has been deleted")
		return nil
	case err != nil:
		rt.Logger.Error("the resource service has returned an unexpected error", "error", err)
		return err
	}

	res := rsp.Resource
	var telemetry pbhcp.TelemetryState
	if err := res.Data.UnmarshalTo(&telemetry); err != nil {
		rt.Logger.Error("error unmarshalling telemetry data", "error", err)
		return err
	}

	var newStatus *pbresource.Status
	if r.resourceApisEnabled && !r.overrideResourceApisEnabledCheck {
		newStatus = &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions: []*pbresource.Condition{{
				Type:    "configured",
				State:   pbresource.Condition_STATE_FALSE,
				Reason:  "DISABLED",
				Message: "telemetry is disabled because resource-apis are enabled",
			}},
		}
	} else {
		newStatus = &pbresource.Status{
			ObservedGeneration: res.Generation,
			Conditions: []*pbresource.Condition{{
				Type:    "configured",
				State:   pbresource.Condition_STATE_TRUE,
				Reason:  "SUCCESS",
				Message: "successfully configured telemetry",
			}},
		}
	}

	statusKey := "consul.io/hcp-telemetry-state"
	if resource.EqualStatus(res.Status[statusKey], newStatus, false) {
		return nil
	}

	_, err = rt.Client.WriteStatus(ctx, &pbresource.WriteStatusRequest{
		Id:     res.Id,
		Key:    statusKey,
		Status: newStatus,
	})

	if err != nil {
		return status.Errorf(status.Code(err), "error writing status: %v", err)
	}

	return nil
}
