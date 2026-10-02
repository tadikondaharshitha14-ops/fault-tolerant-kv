package grpcserver

import (
	"context"

	"fault-tolerant-kv/api/proto"
	"fault-tolerant-kv/internal/kv"
)

type Server struct {
	proto.UnimplementedKVServiceServer
	store *kv.Store
}

func NewServer(store *kv.Store) *Server {
	return &Server{
		store: store,
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
