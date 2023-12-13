package pbserverdiscovery

import (
	"context"
	"errors"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

const (
	streamBufferSize = 16
)

var (
	InvalidStreamTypeError  = errors.New("Received unexpected type to send over the stream")
	HeadersAlreadySentError = errors.New("cannot set more headers after they have already been sent")
)

type inmemStreamFromServer[T proto.Message] struct {
	ctx          context.Context
	ch           chan T
	headersReady chan struct{}
	headers      metadata.MD
	err          error
	cloneOnRecv  bool
}

func newInmemStreamFromServer[T proto.Message](ctx context.Context) *inmemStreamFromServer[T] {
	return &inmemStreamFromServer[T]{
		ctx:          ctx,
		ch:           make(chan T, streamBufferSize),
		headersReady: make(chan struct{}),
	}
}

func (stream *inmemStreamFromServer[T]) Recv() (T, error) {
	var zero T
	select {
	case <-stream.ctx.Done():
		return zero, stream.ctx.Err()
	case event, more := <-stream.ch:
		if !more {
			return zero, stream.err
		}

		return event, nil
	}
}

func (stream *inmemStreamFromServer[T]) Send(msg T) error {
	// ensure that the headers get sent out first.
	if !stream.headersSent() {
		stream.SendHeader(nil)
	}

	select {
	case <-stream.ctx.Done():
		return stream.ctx.Err()
	case stream.ch <- msg:
		return nil
	}
}

func (stream *inmemStreamFromServer[T]) Header() (metadata.MD, error) {
	if err := stream.waitForHeaders(); err != nil {
		return nil, err
	}
	return stream.headers, nil
}

func (stream *inmemStreamFromServer[T]) waitForHeaders() error {
	select {
	case <-stream.ctx.Done():
		return stream.ctx.Err()
	case <-stream.headersReady:
		return nil
	}
}

func (stream *inmemStreamFromServer[T]) headersSent() bool {
	select {
	case <-stream.headersReady:
		return true
	default:
		return false
	}
}

func (stream *inmemStreamFromServer[T]) SetHeader(md metadata.MD) error {
	if stream.headersSent() {
		return HeadersAlreadySentError
	}

	for k, v := range md {
		stream.headers.Append(k, v...)
	}

	return nil
}

func (stream *inmemStreamFromServer[T]) SendHeader(md metadata.MD) error {
	if err := stream.SetHeader(md); err != nil {
		return err
	}

	// close our sentinel channel to record that the headers are now available for reading.
	close(stream.headersReady)
	return nil
}

func (stream *inmemStreamFromServer[T]) Trailer() metadata.MD {
	// trailers are unsupported
	return nil
}

func (stream *inmemStreamFromServer[T]) SetTrailer(_ metadata.MD) {
}

func (stream *inmemStreamFromServer[T]) CloseSend() error {
	// This will never be called because the inmem server stream
	// never had a sending channel open anyways.
	return errors.ErrUnsupported
}

func (stream *inmemStreamFromServer[T]) Context() context.Context {
	return stream.ctx
}

func (stream *inmemStreamFromServer[T]) SendMsg(m interface{}) error {
	// There is no need for inmem streams to ever use this method
	return errors.ErrUnsupported
}

func (_ *inmemStreamFromServer[T]) RecvMsg(m interface{}) error {
	// There is no need for inmem streams to ever use this method
	return errors.ErrUnsupported
}

func (stream *inmemStreamFromServer[T]) CloseServerStream(err error) {
	stream.err = err
	close(stream.ch)
}

type ServerStream[T proto.Message] interface {
	Recv() (T, error)
	grpc.ClientStream
}

type cloningStream[T proto.Message] struct {
	ServerStream[T]
}

func newCloningStream[T proto.Message](stream ServerStream[T]) cloningStream[T] {
	return cloningStream[T]{ServerStream: stream}
}

func (st cloningStream[T]) Recv() (T, error) {
	var zero T
	val, err := st.ServerStream.Recv()
	if err != nil {
		return zero, err
	}

	return proto.Clone(val).(T), nil
}
