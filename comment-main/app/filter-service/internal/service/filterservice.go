package service

import (
	"context"

	pb "comment-main/api/filterService/v1"
)

type FilterServiceService struct {
	pb.UnimplementedFilterServiceServer
}

func NewFilterServiceService() *FilterServiceService {
	return &FilterServiceService{}
}

func (s *FilterServiceService) CreateFilterService(ctx context.Context, req *pb.CreateFilterServiceRequest) (*pb.CreateFilterServiceReply, error) {
	return &pb.CreateFilterServiceReply{}, nil
}
func (s *FilterServiceService) UpdateFilterService(ctx context.Context, req *pb.UpdateFilterServiceRequest) (*pb.UpdateFilterServiceReply, error) {
	return &pb.UpdateFilterServiceReply{}, nil
}
func (s *FilterServiceService) DeleteFilterService(ctx context.Context, req *pb.DeleteFilterServiceRequest) (*pb.DeleteFilterServiceReply, error) {
	return &pb.DeleteFilterServiceReply{}, nil
}
func (s *FilterServiceService) GetFilterService(ctx context.Context, req *pb.GetFilterServiceRequest) (*pb.GetFilterServiceReply, error) {
	return &pb.GetFilterServiceReply{}, nil
}
func (s *FilterServiceService) ListFilterService(ctx context.Context, req *pb.ListFilterServiceRequest) (*pb.ListFilterServiceReply, error) {
	return &pb.ListFilterServiceReply{}, nil
}
