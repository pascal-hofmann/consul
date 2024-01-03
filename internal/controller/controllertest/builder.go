package controllertest

import (
	svc "github.com/hashicorp/consul/agent/grpc-external/services/resource"
	svctest "github.com/hashicorp/consul/agent/grpc-external/services/resource/testing"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/sdk/testutil"
)

type Builder struct {
	serviceBuilder        *svctest.Builder
	controllerRegisterFns []func(*controller.Manager)
}

func NewControllerTestBuilder() *Builder {
	return &Builder{
		// disable cloning because we will enable it after passing the non-cloning variant
		// to the controller manager.
		serviceBuilder: svctest.NewResourceServiceBuilder().WithCloningDisabled(),
	}
}

func (b *Builder) WithV2Tenancy(useV2Tenancy bool) *Builder {
	b.serviceBuilder = b.serviceBuilder.WithV2Tenancy(useV2Tenancy)
	return b
}

func (b *Builder) Registry() resource.Registry {
	return b.serviceBuilder.Registry()
}

func (b *Builder) WithResourceRegisterFns(registerFns ...func(resource.Registry)) *Builder {
	b.serviceBuilder = b.serviceBuilder.WithRegisterFns(registerFns...)
	return b
}

func (b *Builder) WithControllerRegisterFns(registerFns ...func(*controller.Manager)) *Builder {
	for _, registerFn := range registerFns {
		b.controllerRegisterFns = append(b.controllerRegisterFns, registerFn)
	}
	return b
}

func (b *Builder) WithACLResolver(aclResolver svc.ACLResolver) *Builder {
	b.serviceBuilder = b.serviceBuilder.WithACLResolver(aclResolver)
	return b
}

func (b *Builder) WithTenancies(tenancies ...*pbresource.Tenancy) *Builder {
	b.serviceBuilder = b.serviceBuilder.WithTenancies(tenancies...)
	return b
}

func (b *Builder) Run(t testutil.TestingTB) pbresource.ResourceServiceClient {
	t.Helper()

	ctx := testutil.TestContext(t)

	client := b.serviceBuilder.Run(t)

	mgr := controller.NewManager(client, testutil.Logger(t))
	for _, register := range b.controllerRegisterFns {
		register(mgr)
	}

	mgr.SetRaftLeader(true)
	go mgr.Run(ctx)

	// auto-clone messages going through the client so that test
	// code is free to modify objects in place without cloning themselves.
	return pbresource.NewCloningResourceServiceClient(client)
}
