package consul

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/hashicorp/consul/acl"
	aclgrpc "github.com/hashicorp/consul/agent/grpc-external/services/acl"
	"github.com/hashicorp/consul/agent/grpc-external/services/connectca"
	"github.com/hashicorp/consul/agent/grpc-external/services/dataplane"
	"github.com/hashicorp/consul/agent/grpc-external/services/peerstream"
	resourcegrpc "github.com/hashicorp/consul/agent/grpc-external/services/resource"
	"github.com/hashicorp/consul/agent/grpc-external/services/serverdiscovery"
	"github.com/hashicorp/consul/agent/grpc-internal/services/subscribe"
	"github.com/hashicorp/consul/agent/rpc/operator"
	"github.com/hashicorp/consul/agent/rpc/peering"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/tenancy"
	"github.com/hashicorp/consul/lib/stringslice"
	"github.com/hashicorp/consul/logging"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/proto/private/pbsubscribe"
)

func (s *Server) registerResourceServiceServer(typeRegistry resource.Registry, resolver resourcegrpc.ACLResolver, registrars ...grpc.ServiceRegistrar) error {
	if s.raftStorageBackend == nil {
		return fmt.Errorf("raft storage backend cannot be nil")
	}

	var tenancyBridge resourcegrpc.TenancyBridge
	if s.useV2Tenancy {
		tenancyBridge = tenancy.NewV2TenancyBridge().WithClient(
			pbresource.NewResourceServiceClient(s.insecureUnsafeGRPCChan),
		)
	} else {
		tenancyBridge = NewV1TenancyBridge(s)
	}

	srv := resourcegrpc.NewServer(resourcegrpc.Config{
		Registry:      typeRegistry,
		Backend:       s.raftStorageBackend,
		ACLResolver:   resolver,
		Logger:        s.loggers.Named(logging.GRPCAPI).Named(logging.Resource),
		TenancyBridge: tenancyBridge,
		UseV2Tenancy:  s.useV2Tenancy,
	})

	for _, reg := range registrars {
		pbresource.RegisterResourceServiceServer(reg, srv)
	}
	return nil
}

func (s *Server) registerACLServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := aclgrpc.NewServer(aclgrpc.Config{
		ACLsEnabled: s.config.ACLsEnabled,
		ForwardRPC: func(info structs.RPCInfo, fn func(*grpc.ClientConn) error) (bool, error) {
			return s.ForwardGRPC(s.grpcConnPool, info, fn)
		},
		InPrimaryDatacenter: s.InPrimaryDatacenter(),
		LoadAuthMethod: func(methodName string, entMeta *acl.EnterpriseMeta) (*structs.ACLAuthMethod, aclgrpc.Validator, error) {
			return s.loadAuthMethod(methodName, entMeta)
		},
		LocalTokensEnabled:        s.LocalTokensEnabled,
		Logger:                    s.loggers.Named(logging.GRPCAPI).Named(logging.ACL),
		NewLogin:                  func() aclgrpc.Login { return s.aclLogin() },
		NewTokenWriter:            func() aclgrpc.TokenWriter { return s.aclTokenWriter() },
		PrimaryDatacenter:         s.config.PrimaryDatacenter,
		ValidateEnterpriseRequest: s.validateEnterpriseRequest,
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}

func (s *Server) registerPeerStreamServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) (*peerstream.Server, error) {
	srv := peerstream.NewServer(peerstream.Config{
		Backend:        s.peeringBackend,
		GetStore:       func() peerstream.StateStore { return s.FSM().State() },
		Logger:         s.loggers.Named(logging.GRPCAPI).Named(logging.PeerStream),
		ACLResolver:    s.ACLResolver,
		Datacenter:     s.config.Datacenter,
		ConnectEnabled: s.config.ConnectEnabled,
		ForwardRPC: func(info structs.RPCInfo, fn func(*grpc.ClientConn) error) (bool, error) {
			// Only forward the request if the dc in the request matches the server's datacenter.
			if info.RequestDatacenter() != "" && info.RequestDatacenter() != config.Datacenter {
				return false, fmt.Errorf("requests to generate peering tokens cannot be forwarded to remote datacenters")
			}
			return s.ForwardGRPC(s.grpcConnPool, info, fn)
		},
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return srv, nil
}

func (s *Server) registerPeeringServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	if s.peeringBackend == nil {
		panic("peeringBackend is required during construction")
	}

	srv := peering.NewServer(peering.Config{
		Backend: s.peeringBackend,
		Tracker: s.peerStreamServer.Tracker,
		Logger:  s.loggers.Named(logging.GRPCAPI).Named(logging.Peering),
		ForwardRPC: func(info structs.RPCInfo, fn func(*grpc.ClientConn) error) (bool, error) {
			// Only forward the request if the dc in the request matches the server's datacenter.
			if info.RequestDatacenter() != "" && info.RequestDatacenter() != config.Datacenter {
				return false, fmt.Errorf("requests to generate peering tokens cannot be forwarded to remote datacenters")
			}
			return s.ForwardGRPC(s.grpcConnPool, info, fn)
		},
		Datacenter:     config.Datacenter,
		ConnectEnabled: config.ConnectEnabled,
		PeeringEnabled: config.PeeringEnabled,
		Locality:       config.Locality,
		FSMServer:      s,
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}

func (s *Server) registerOperatorServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := operator.NewServer(operator.Config{
		Backend: NewOperatorBackend(s),
		Logger:  deps.Logger.Named("grpc-api.operator"),
		ForwardRPC: func(info structs.RPCInfo, fn func(*grpc.ClientConn) error) (bool, error) {
			// Only forward the request if the dc in the request matches the server's datacenter.
			if info.RequestDatacenter() != "" && info.RequestDatacenter() != config.Datacenter {
				return false, fmt.Errorf("requests to transfer leader cannot be forwarded to remote datacenters")
			}
			return s.ForwardGRPC(s.grpcConnPool, info, fn)
		},
		Datacenter: config.Datacenter,
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}

func (s *Server) registerStreamSubscriptionServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := subscribe.NewServer(
		&subscribeBackend{srv: s, connPool: deps.GRPCConnPool},
		s.loggers.Named(logging.GRPCAPI).Named("subscription"),
	)

	for _, reg := range registrars {
		pbsubscribe.RegisterStateChangeSubscriptionServer(reg, srv)
	}

	return nil
}

func (s *Server) registerConnectCAServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := connectca.NewServer(connectca.Config{
		Publisher:   s.publisher,
		GetStore:    func() connectca.StateStore { return s.FSM().State() },
		Logger:      s.loggers.Named(logging.GRPCAPI).Named(logging.ConnectCA),
		ACLResolver: s.ACLResolver,
		CAManager:   s.caManager,
		ForwardRPC: func(info structs.RPCInfo, fn func(*grpc.ClientConn) error) (bool, error) {
			return s.ForwardGRPC(s.grpcConnPool, info, fn)
		},
		ConnectEnabled: s.config.ConnectEnabled,
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}

func (s *Server) registerDataplaneServer(config *Config, deps Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := dataplane.NewServer(dataplane.Config{
		GetStore:          func() dataplane.StateStore { return s.FSM().State() },
		Logger:            s.loggers.Named(logging.GRPCAPI).Named(logging.Dataplane),
		ACLResolver:       s.ACLResolver,
		Datacenter:        s.config.Datacenter,
		EnableV2:          stringslice.Contains(deps.Experiments, CatalogResourceExperimentName),
		ResourceAPIClient: pbresource.NewResourceServiceClient(s.insecureSafeGRPCChan),
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}

func (s *Server) registerServerDiscoveryServer(_ *Config, _ Deps, registrars ...grpc.ServiceRegistrar) error {
	srv := serverdiscovery.NewServer(serverdiscovery.Config{
		Publisher:   s.publisher,
		ACLResolver: s.ACLResolver,
		Logger:      s.loggers.Named(logging.GRPCAPI).Named(logging.ServerDiscovery),
	})

	for _, reg := range registrars {
		srv.Register(reg)
	}

	return nil
}
