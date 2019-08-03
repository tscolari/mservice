package transports

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/tscolari/mservice/pkg/endpoints"
	"github.com/tscolari/mservice/pkg/pb"
	"github.com/tscolari/mservice/pkg/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewTaxGRPCClient returns an TaxService backed by a gRPC server at the other end
// of the conn.
func NewTaxGRPCClient(conn *grpc.ClientConn) services.Tax {
	var addEndpoint endpoint.Endpoint
	addEndpoint = grpctransport.NewClient(
		conn,
		"pb.Tax",
		"Add",
		encodeGRPCAddRequest,
		decodeGRPCAddResponse,
		pb.AddReply{},
	).Endpoint()

	var subEndpoint endpoint.Endpoint
	subEndpoint = grpctransport.NewClient(
		conn,
		"pb.Tax",
		"Sub",
		encodeGRPCSubRequest,
		decodeGRPCSubResponse,
		pb.SubReply{},
	).Endpoint()

	return endpoints.Tax{
		AddEndpoint: addEndpoint,
		SubEndpoint: subEndpoint,
	}
}

func encodeGRPCAddRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.AddRequest)
	return &pb.AddRequest{Value: float32(req.Value)}, nil
}

func encodeGRPCSubRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.SubRequest)
	return &pb.SubRequest{Value: float32(req.Value)}, nil
}

func decodeGRPCAddResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddReply)
	return endpoints.AddResponse{Value: float64(reply.Value), Err: toErr(reply.Err)}, nil
}

func decodeGRPCSubResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SubReply)
	return endpoints.SubResponse{Value: float64(reply.Value), Err: toErr(reply.Err)}, nil
}

func toErr(org string) error {
	var err error
	if org != "" {
		err = errors.New(org)
	}

	return err
}
