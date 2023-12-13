// Code generated by protoc-gen-grpc-inmem. DO NOT EDIT.

package pbdataplane

import (
	"context"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ DataplaneServiceClient = &InmemDataplaneServiceClient{}

// InmemDataplaneServiceClient implements the DataplaneServiceClient interface by directly
// calling methods on the server implementation. This avoids unnecessary serialization
// and deserialization when the client and server are executing in the same process.
// The caveat to this performance optimization is that not all standard network
// based gRPC server options and client dialing options will be supported. The
// second caveat is that the caller needs to take care to ensure they don't modify
// shared data.
type InmemDataplaneServiceClient struct {
	srv DataplaneServiceServer
}

func NewInmemDataplaneServiceClient(srv DataplaneServiceServer) (DataplaneServiceClient, error) {
	return &InmemDataplaneServiceClient{
		srv: srv,
	}, nil
}

func (c *InmemDataplaneServiceClient) GetSupportedDataplaneFeatures(ctx context.Context, in *GetSupportedDataplaneFeaturesRequest, opts ...grpc.CallOption) (*GetSupportedDataplaneFeaturesResponse, error) {
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.srv.GetSupportedDataplaneFeatures(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *InmemDataplaneServiceClient) GetEnvoyBootstrapParams(ctx context.Context, in *GetEnvoyBootstrapParamsRequest, opts ...grpc.CallOption) (*GetEnvoyBootstrapParamsResponse, error) {
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.srv.GetEnvoyBootstrapParams(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// compile-time check to ensure that the generator is implementing all
// of the grpc client interfaces methods.
var _ DataplaneServiceClient = CloningDataplaneServiceClient{}

// CloningDataplaneServiceClient implements the DataplaneServiceClient interface by wrapping
// another implementation and copying all protobuf messages that pass through the client.
// This is mainly useful to wrap the InmemDataplaneServiceClient client to insulate users of that
// client from having to care about potential immutability of data they receive or having
// the server implementation mutate their internal memory.
type CloningDataplaneServiceClient struct {
	DataplaneServiceClient
}

func NewCloningDataplaneServiceClient(client DataplaneServiceClient) CloningDataplaneServiceClient {
	return CloningDataplaneServiceClient{
		DataplaneServiceClient: client,
	}
}

func (c CloningDataplaneServiceClient) GetSupportedDataplaneFeatures(ctx context.Context, in *GetSupportedDataplaneFeaturesRequest, opts ...grpc.CallOption) (*GetSupportedDataplaneFeaturesResponse, error) {
	in = proto.Clone(in).(*GetSupportedDataplaneFeaturesRequest)
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.DataplaneServiceClient.GetSupportedDataplaneFeatures(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*GetSupportedDataplaneFeaturesResponse), nil
}

func (c CloningDataplaneServiceClient) GetEnvoyBootstrapParams(ctx context.Context, in *GetEnvoyBootstrapParamsRequest, opts ...grpc.CallOption) (*GetEnvoyBootstrapParamsResponse, error) {
	in = proto.Clone(in).(*GetEnvoyBootstrapParamsRequest)
	md, found := metadata.FromOutgoingContext(ctx)
	if found {
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	out, err := c.DataplaneServiceClient.GetEnvoyBootstrapParams(ctx, in)
	if err != nil {
		return nil, err
	}

	return proto.Clone(out).(*GetEnvoyBootstrapParamsResponse), nil
}
