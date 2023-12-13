// Code generated by protoc-gen-grpc-inmem. DO NOT EDIT.

package pbdns

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ DNSServiceClient = &InmemDNSServiceClient{}

// InmemDNSServiceClient implements the DNSServiceClient interface by directly
// calling methods on the server implementation. This avoids unnecessary serialization
// and deserialization when the client and server are executing in the same process.
// The caveat to this performance optimization is that not all standard network
// based gRPC server options and client dialing options will be supported. The
// second caveat is that the caller needs to take care to ensure they don't modify
// shared data.
type InmemDNSServiceClient struct {
	srv DNSServiceServer
}

func NewInmemDNSServiceClient(srv DNSServiceServer) (DNSServiceClient, error) {
	return &InmemDNSServiceClient{
		srv: srv,
	}, nil
}

func (c *InmemDNSServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.srv.Query(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ DNSServiceClient = CloningDNSServiceClient{}

// CloningDNSServiceClient implements the DNSServiceClient interface by wrapping
// another implementation and copying all protobuf messages that pass through the client.
// This is mainly useful to wrap the InmemDNSServiceClient client to insulate users of that
// client from having to care about potential immutability of data they receive or having
// the server implementation mutate their internal memory.
type CloningDNSServiceClient struct {
	DNSServiceClient
}

func NewCloningDNSServiceClient(client DNSServiceClient) CloningDNSServiceClient {
	return CloningDNSServiceClient{
		DNSServiceClient: client,
	}
}

func (c CloningDNSServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	in = proto.Clone(in).(*QueryRequest)
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.DNSServiceClient.Query(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*QueryResponse), nil
}
