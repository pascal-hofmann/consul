// Code generated by protoc-gen-resource-types. DO NOT EDIT.

package hcpv1

import (
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName = "hcp"
	Version   = "v1"

	LinkKind           = "Link"
	TelemetryStateKind = "TelemetryState"
)

var (
	LinkType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         LinkKind,
	}

	TelemetryStateType = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: Version,
		Kind:         TelemetryStateKind,
	}
)
