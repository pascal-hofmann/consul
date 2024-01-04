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

	LinkedMessageFormat                = "Successfully linked to cluster '%s'"
	DisabledResourceAPIsEnabledMessage = "Link is disabled because resource-apis are enabled"
	FailedMessage                      = "Failed to link due to unexpected error"
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
)

func ConditionLinked(resourceId string) *pbresource.Condition {
	return &pbresource.Condition{
		Type:    StatusLinked,
		State:   pbresource.Condition_STATE_TRUE,
		Reason:  LinkedReason,
		Message: fmt.Sprintf(LinkedMessageFormat, resourceId),
	}
}
