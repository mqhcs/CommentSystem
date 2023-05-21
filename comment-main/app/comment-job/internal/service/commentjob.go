package service

import (
	"context"

	pb "comment-main/api/commentJob/v1"
)

type CommentJobService struct {
	pb.UnimplementedCommentJobServer
}

func NewCommentJobService() *CommentJobService {
	return &CommentJobService{}
}

func (s *CommentJobService) CreateCommentJob(ctx context.Context, req *pb.CreateCommentJobRequest) (*pb.CreateCommentJobReply, error) {
	return &pb.CreateCommentJobReply{}, nil
}
func (s *CommentJobService) UpdateCommentJob(ctx context.Context, req *pb.UpdateCommentJobRequest) (*pb.UpdateCommentJobReply, error) {
	return &pb.UpdateCommentJobReply{}, nil
}
func (s *CommentJobService) DeleteCommentJob(ctx context.Context, req *pb.DeleteCommentJobRequest) (*pb.DeleteCommentJobReply, error) {
	return &pb.DeleteCommentJobReply{}, nil
}
func (s *CommentJobService) GetCommentJob(ctx context.Context, req *pb.GetCommentJobRequest) (*pb.GetCommentJobReply, error) {
	return &pb.GetCommentJobReply{}, nil
}
func (s *CommentJobService) ListCommentJob(ctx context.Context, req *pb.ListCommentJobRequest) (*pb.ListCommentJobReply, error) {
	return &pb.ListCommentJobReply{}, nil
}
