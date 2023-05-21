package service

import (
	"context"

	pb "comment-main/api/commentService/v1"
)

type CommentServiceService struct {
	pb.UnimplementedCommentServiceServer
}

func NewCommentServiceService() *CommentServiceService {
	return &CommentServiceService{}
}

func (s *CommentServiceService) CreateCommentService(ctx context.Context, req *pb.CreateCommentServiceRequest) (*pb.CreateCommentServiceReply, error) {
	return &pb.CreateCommentServiceReply{}, nil
}
func (s *CommentServiceService) UpdateCommentService(ctx context.Context, req *pb.UpdateCommentServiceRequest) (*pb.UpdateCommentServiceReply, error) {
	return &pb.UpdateCommentServiceReply{}, nil
}
func (s *CommentServiceService) DeleteCommentService(ctx context.Context, req *pb.DeleteCommentServiceRequest) (*pb.DeleteCommentServiceReply, error) {
	return &pb.DeleteCommentServiceReply{}, nil
}
func (s *CommentServiceService) GetCommentService(ctx context.Context, req *pb.GetCommentServiceRequest) (*pb.GetCommentServiceReply, error) {
	return &pb.GetCommentServiceReply{}, nil
}
func (s *CommentServiceService) ListCommentService(ctx context.Context, req *pb.ListCommentServiceRequest) (*pb.ListCommentServiceReply, error) {
	return &pb.ListCommentServiceReply{}, nil
}
