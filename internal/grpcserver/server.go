package grpcserver

import (
	"context"

	"fault-tolerant-kv/api/proto"
	"fault-tolerant-kv/internal/kv"
)

type Server struct {
	proto.UnimplementedKVServiceServer

	store  *kv.Store
	nodeID int
}

func NewServer(store *kv.Store, nodeID int) *Server {
	return &Server{
		store:  store,
		nodeID: nodeID,
	}
}

func (s *Server) Set(
	ctx context.Context,
	req *proto.SetRequest,
) (*proto.SetResponse, error) {

	s.store.Set(req.GetKey(), req.GetValue())

	return &proto.SetResponse{
		Success: true,
	}, nil
}

func (s *Server) Get(
	ctx context.Context,
	req *proto.GetRequest,
) (*proto.GetResponse, error) {

	value, ok := s.store.Get(req.GetKey())

	return &proto.GetResponse{
		Value: value,
		Found: ok,
	}, nil
}

func (s *Server) Delete(
	ctx context.Context,
	req *proto.DeleteRequest,
) (*proto.DeleteResponse, error) {

	s.store.Delete(req.GetKey())

	return &proto.DeleteResponse{
		Success: true,
	}, nil
}

func (s *Server) Replicate(
	ctx context.Context,
	req *proto.ReplicateRequest,
) (*proto.ReplicateResponse, error) {

	switch req.GetOperation() {
	case "SET":
		s.store.Set(
			req.GetKey(),
			req.GetValue(),
		)

	case "DELETE":
		s.store.Delete(req.GetKey())

	default:
		return &proto.ReplicateResponse{
			Success: false,
		}, nil
	}

	return &proto.ReplicateResponse{
		Success: true,
	}, nil
}

func (s *Server) NodeStatus(
	ctx context.Context,
	req *proto.NodeStatusRequest,
) (*proto.NodeStatusResponse, error) {

	return &proto.NodeStatusResponse{
		NodeId: int32(s.nodeID),
		Alive:  true,
	}, nil
}
