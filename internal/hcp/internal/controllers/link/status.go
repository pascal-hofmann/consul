// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package link

import (
	"fmt"

	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	StatusKey = "consul.io/hcp/link"

	StatusLinked                         = "linked"
	LinkedReason                         = "SUCCESS"
	DisabledReasonV2ResourcesUnsupported = "DISABLED_V2_RESOURCES_UNSUPPORTED"
	FailedReason                         = "FAILED"
	UnauthorizedReason                   = "UNAUTHORIZED"
	ForbiddenReason                      = "FORBIDDEN"

	LinkedMessageFormat                = "Successfully linked to cluster '%s'"
	DisabledResourceAPIsEnabledMessage = "Link is disabled because resource-apis are enabled"
	FailedMessage                      = "Failed to link due to unexpected error"
	UnauthorizedMessage                = "Access denied, check client_id and client_secret"
	ForbiddenMessage                   = "Access denied, check the resource_id"
)

var (
	ConditionDisabled = &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_FALSE,
		Reason:  DisabledReasonV2ResourcesUnsupported,
		Message: DisabledResourceAPIsEnabledMessage,
	}
	ConditionFailed = &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_FALSE,
		Reason:  FailedReason,
		Message: FailedMessage,
	}
	ConditionUnauthorized = &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_FALSE,
		Reason:  UnauthorizedReason,
		Message: UnauthorizedMessage,
	}
	ConditionForbidden = &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_FALSE,
		Reason:  ForbiddenReason,
		Message: ForbiddenMessage,
	}
)

func ConditionLinked(resourceId string) *pbresource.Condition {
	return &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_TRUE,
		Reason:  LinkedReason,
		Message: fmt.Sprintf(LinkedMessageFormat, resourceId),
	}
}
