// Code generated by protoc-gen-grpc-inmem. DO NOT EDIT.

package pbserverdiscovery

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ ServerDiscoveryServiceClient = &InmemServerDiscoveryServiceClient{}

// InmemServerDiscoveryServiceClient implements the ServerDiscoveryServiceClient interface by directly
// calling methods on the server implementation. This avoids unnecessary serialization
// and deserialization when the client and server are executing in the same process.
// The caveat to this performance optimization is that not all standard network
// based gRPC server options and client dialing options will be supported. The
// second caveat is that the caller needs to take care to ensure they don't modify
// shared data.
type InmemServerDiscoveryServiceClient struct {
	srv ServerDiscoveryServiceServer
}

func NewInmemServerDiscoveryServiceClient(srv ServerDiscoveryServiceServer) (ServerDiscoveryServiceClient, error) {
	return &InmemServerDiscoveryServiceClient{
		srv: srv,
	}, nil
}

func (c *InmemServerDiscoveryServiceClient) WatchServers(ctx context.Context, in *WatchServersRequest, opts ...grpc.CallOption) (ServerDiscoveryService_WatchServersClient, error) {
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}
	st := newInmemStreamFromServer[*WatchServersResponse](ctx)
	go func() {
		st.CloseServerStream(c.srv.WatchServers(in, st))
	}()

	return st, nil
}

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ ServerDiscoveryServiceClient = CloningServerDiscoveryServiceClient{}

// CloningServerDiscoveryServiceClient implements the ServerDiscoveryServiceClient interface by wrapping
// another implementation and copying all protobuf messages that pass through the client.
// This is mainly useful to wrap the InmemServerDiscoveryServiceClient client to insulate users of that
// client from having to care about potential immutability of data they receive or having
// the server implementation mutate their internal memory.
type CloningServerDiscoveryServiceClient struct {
	ServerDiscoveryServiceClient
}

func NewCloningServerDiscoveryServiceClient(client ServerDiscoveryServiceClient) CloningServerDiscoveryServiceClient {
	return CloningServerDiscoveryServiceClient{
		ServerDiscoveryServiceClient: client,
	}
}

func (c CloningServerDiscoveryServiceClient) WatchServers(ctx context.Context, in *WatchServersRequest, opts ...grpc.CallOption) (ServerDiscoveryService_WatchServersClient, error) {
	in = proto.Clone(in).(*WatchServersRequest)
	st, err := c.ServerDiscoveryServiceClient.WatchServers(ctx, in)
	if err != nil {
		return nil, err
	}

	return newCloningStream(st), nil
}
