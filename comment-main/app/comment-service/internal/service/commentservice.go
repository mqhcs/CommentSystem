package service

import (
	"context"

	pb "comment-main/api/commentService/v1"
)

func (s *CommentServiceService) CreateReplyService(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyReply, error) {
	return &pb.CreateReplyReply{}, nil
}
func (s *CommentServiceService) UpdateReplyService(ctx context.Context, req *pb.UpdateReplyRequest) (*pb.UpdateReplyReply, error) {
	return &pb.UpdateReplyReply{}, nil
}
func (s *CommentServiceService) DeleteReplyService(ctx context.Context, req *pb.DeleteReplyRequest) (*pb.DeleteReplyReply, error) {
	return &pb.DeleteReplyReply{}, nil
}
func (s *CommentServiceService) GetReplyService(ctx context.Context, req *pb.GetReplyRequest) (*pb.GetReplyReply, error) {
	return &pb.GetReplyReply{}, nil
}
func (s *CommentServiceService) ListReplyService(ctx context.Context, req *pb.ListReplyRequest) (*pb.ListReplyReply, error) {
	return &pb.ListReplyReply{}, nil
}
