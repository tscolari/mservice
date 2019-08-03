package transports

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/tscolari/mservice/pkg/endpoints"
	"github.com/tscolari/mservice/pkg/pb"
	"golang.org/x/net/context"
)

type taxServer struct {
	add grpctransport.Handler
	sub grpctransport.Handler
}

// NewTaxGRPCServer make the Tax endpoint available	as a gRPC TaxServer
func NewTaxGRPCServer(endpoints endpoints.Tax) pb.TaxServer {
	return &taxServer{
		add: grpctransport.NewServer(
			endpoints.AddEndpoint,
			decodeGRPCAddRequest,
			encodeGRPCAddResponse,
		),
		sub: grpctransport.NewServer(
			endpoints.SubEndpoint,
			decodeGRPCSubRequest,
			encodeGRPCSubResponse,
		),
	}
}

func (s *taxServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	_, rep, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*pb.AddReply), nil
}

func (s *taxServer) Sub(ctx context.Context, req *pb.SubRequest) (*pb.SubReply, error) {
	_, rep, err := s.sub.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep.(*pb.SubReply), nil
}

func decodeGRPCAddRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddRequest)
	return endpoints.AddRequest{Value: float64(req.Value)}, nil
}

func decodeGRPCSubRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SubRequest)
	return endpoints.SubRequest{Value: float64(req.Value)}, nil
}

func encodeGRPCAddResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.AddResponse)
	return &pb.AddReply{Value: float32(resp.Value), Err: toStr(resp.Err)}, nil
}

func encodeGRPCSubResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SubResponse)
	return &pb.SubReply{Value: float32(resp.Value), Err: toStr(resp.Err)}, nil
}

func toStr(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
