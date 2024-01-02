// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package link

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	svctest "github.com/hashicorp/consul/agent/grpc-external/services/resource/testing"
	"github.com/hashicorp/consul/agent/hcp/config"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/hcp/internal/types"
	"github.com/hashicorp/consul/internal/resource/resourcetest"
	rtest "github.com/hashicorp/consul/internal/resource/resourcetest"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/sdk/testutil"
)

type controllerSuite struct {
	suite.Suite

	ctx    context.Context
	client *rtest.Client
	rt     controller.Runtime

	ctl       linkReconciler
	tenancies []*pbresource.Tenancy
}

func (suite *controllerSuite) SetupTest() {
	suite.ctx = testutil.TestContext(suite.T())
	suite.tenancies = resourcetest.TestTenancies()
	client := svctest.NewResourceServiceBuilder().
		WithRegisterFns(types.Register).
		WithTenancies(suite.tenancies...).
		Run(suite.T())

	suite.rt = controller.Runtime{
		Client: client,
		Logger: testutil.Logger(suite.T()),
	}
	suite.client = rtest.NewClient(client)
}

func TestLinkController(t *testing.T) {
	suite.Run(t, new(controllerSuite))
}

func (suite *controllerSuite) deleteResourceFunc(id *pbresource.ID) func() {
	return func() {
		suite.client.MustDelete(suite.T(), id)
	}
}

func (suite *controllerSuite) TestController_Ok() {
	// Run the controller manager
	mgr := controller.NewManager(suite.client, suite.rt.Logger)
	mgr.Register(LinkController(false, false, config.CloudConfig{}))
	mgr.SetRaftLeader(true)
	go mgr.Run(suite.ctx)

	linkData := &pbhcp.Link{
		ClientId:     "abc",
		ClientSecret: "abc",
		ResourceId:   "abc",
	}

	link := rtest.Resource(pbhcp.LinkType, "global").
		WithData(suite.T(), linkData).
		Write(suite.T(), suite.client)

	suite.T().Cleanup(suite.deleteResourceFunc(link.Id))

	suite.client.WaitForStatusCondition(suite.T(), link.Id, StatusKey, ConditionLinked(linkData.ResourceId))
}

func (suite *controllerSuite) TestController_Initialize() {
	// Run the controller manager with a configured link
	mgr := controller.NewManager(suite.client, suite.rt.Logger)

	cloudCfg := config.CloudConfig{
		ClientID:     "client-id-abc",
		ClientSecret: "client-secret-abc",
		ResourceID:   "resource-id-abc",
	}
	mgr.Register(LinkController(false, false, cloudCfg))
	mgr.SetRaftLeader(true)
	go mgr.Run(suite.ctx)

	// Wait for link to be created by initializer
	id := &pbresource.ID{
		Type: pbhcp.LinkType,
		Name: types.LinkName,
	}
	suite.T().Cleanup(suite.deleteResourceFunc(id))
	r := suite.client.WaitForResourceExists(suite.T(), id)

	// Check that created link has expected values
	var link pbhcp.Link
	err := r.Data.UnmarshalTo(&link)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), cloudCfg.ResourceID, link.ResourceId)
	require.Equal(suite.T(), cloudCfg.ClientID, link.ClientId)
	require.Equal(suite.T(), cloudCfg.ClientSecret, link.ClientSecret)

	// Wait for link to be connected successfully
	suite.client.WaitForStatusCondition(suite.T(), id, StatusKey, ConditionLinked(link.ResourceId))
}

func (suite *controllerSuite) TestControllerResourceApisEnabled_LinkDisabled() {
	// Run the controller manager
	mgr := controller.NewManager(suite.client, suite.rt.Logger)
	mgr.Register(LinkController(true, false, config.CloudConfig{}))
	mgr.SetRaftLeader(true)
	go mgr.Run(suite.ctx)

	linkData := &pbhcp.Link{
		ClientId:     "abc",
		ClientSecret: "abc",
		ResourceId:   "abc",
	}
	// The controller is currently a no-op, so there is nothing to test other than making sure we do not panic
	link := rtest.Resource(pbhcp.LinkType, "global").
		WithData(suite.T(), linkData).
		Write(suite.T(), suite.client)

	suite.T().Cleanup(suite.deleteResourceFunc(link.Id))

	suite.client.WaitForStatusCondition(suite.T(), link.Id, StatusKey, ConditionDisabled)
}

func (suite *controllerSuite) TestControllerResourceApisEnabledWithOverride_LinkNotDisabled() {
	// Run the controller manager
	mgr := controller.NewManager(suite.client, suite.rt.Logger)
	mgr.Register(LinkController(true, true, config.CloudConfig{}))
	mgr.SetRaftLeader(true)
	go mgr.Run(suite.ctx)

	linkData := &pbhcp.Link{
		ClientId:     "abc",
		ClientSecret: "abc",
		ResourceId:   "abc",
	}
	// The controller is currently a no-op, so there is nothing to test other than making sure we do not panic
	link := rtest.Resource(pbhcp.LinkType, "global").
		WithData(suite.T(), linkData).
		Write(suite.T(), suite.client)

	suite.T().Cleanup(suite.deleteResourceFunc(link.Id))

	suite.client.WaitForStatusCondition(suite.T(), link.Id, StatusKey, ConditionLinked(linkData.ResourceId))
}
