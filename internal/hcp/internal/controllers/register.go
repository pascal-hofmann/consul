// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package controllers

import (
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/hcp/internal/controllers/link"
)

type Dependencies = link.Dependencies

func Register(mgr *controller.Manager, deps Dependencies) {
	mgr.Register(link.LinkController(
		link.Dependencies{
			ResourceApisEnabled:              deps.ResourceApisEnabled,
			OverrideResourceApisEnabledCheck: deps.OverrideResourceApisEnabledCheck,
			DataDir:                          deps.DataDir,
		},
	))
}
