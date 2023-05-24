package service

import (
	"context"

	pb "comment-main/api/commentService/v1"
)

func (s *CommentServiceService) CreateReplyIndexService(ctx context.Context, req *pb.CreateReplyIndexRequest) (*pb.CreateReplyIndexReply, error) {
	return &pb.CreateReplyIndexReply{}, nil
}
func (s *CommentServiceService) UpdateReplyIndexService(ctx context.Context, req *pb.UpdateReplyIndexRequest) (*pb.UpdateReplyIndexReply, error) {
	return &pb.UpdateReplyIndexReply{}, nil
}
func (s *CommentServiceService) DeleteReplyIndexService(ctx context.Context, req *pb.DeleteReplyIndexRequest) (*pb.DeleteReplyIndexReply, error) {
	return &pb.DeleteReplyIndexReply{}, nil
}
func (s *CommentServiceService) GetReplyIndexService(ctx context.Context, req *pb.GetReplyIndexRequest) (*pb.GetReplyIndexReply, error) {
	return &pb.GetReplyIndexReply{}, nil
}
func (s *CommentServiceService) ListReplyIndexService(ctx context.Context, req *pb.ListReplyRequest) (*pb.ListReplyIndexReply, error) {
	return &pb.ListReplyIndexReply{}, nil
}
