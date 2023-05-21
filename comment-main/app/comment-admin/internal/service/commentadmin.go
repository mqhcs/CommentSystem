package service

import (
	"context"

	pb "comment-main/api/commentAdmin/v1"
)

type CommentAdminService struct {
	pb.UnimplementedCommentAdminServer
}

func NewCommentAdminService() *CommentAdminService {
	return &CommentAdminService{}
}

func (s *CommentAdminService) CreateCommentAdmin(ctx context.Context, req *pb.CreateCommentAdminRequest) (*pb.CreateCommentAdminReply, error) {
	return &pb.CreateCommentAdminReply{}, nil
}
func (s *CommentAdminService) UpdateCommentAdmin(ctx context.Context, req *pb.UpdateCommentAdminRequest) (*pb.UpdateCommentAdminReply, error) {
	return &pb.UpdateCommentAdminReply{}, nil
}
func (s *CommentAdminService) DeleteCommentAdmin(ctx context.Context, req *pb.DeleteCommentAdminRequest) (*pb.DeleteCommentAdminReply, error) {
	return &pb.DeleteCommentAdminReply{}, nil
}
func (s *CommentAdminService) GetCommentAdmin(ctx context.Context, req *pb.GetCommentAdminRequest) (*pb.GetCommentAdminReply, error) {
	return &pb.GetCommentAdminReply{}, nil
}
func (s *CommentAdminService) ListCommentAdmin(ctx context.Context, req *pb.ListCommentAdminRequest) (*pb.ListCommentAdminReply, error) {
	return &pb.ListCommentAdminReply{}, nil
}
