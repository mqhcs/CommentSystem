package service

import (
	"context"

	pb "comment-main/api/commentService/v1"
)

func (s *CommentServiceService) CreateReplyAreaService(ctx context.Context, req *pb.CreateReplyAreaRequest) (*pb.CreateReplyAreaReply, error) {
	return &pb.CreateReplyAreaReply{}, nil
}
func (s *CommentServiceService) UpdateReplyAreaService(ctx context.Context, req *pb.UpdateReplyAreaRequest) (*pb.UpdateReplyAreaReply, error) {
	return &pb.UpdateReplyAreaReply{}, nil
}
func (s *CommentServiceService) DeleteReplyAreaService(ctx context.Context, req *pb.DeleteReplyAreaRequest) (*pb.DeleteReplyAreaReply, error) {
	return &pb.DeleteReplyAreaReply{}, nil
}
func (s *CommentServiceService) GetReplyAreaService(ctx context.Context, req *pb.GetReplyAreaRequest) (*pb.GetReplyAreaReply, error) {
	return &pb.GetReplyAreaReply{}, nil
}
func (s *CommentServiceService) ListReplyAreaService(ctx context.Context, req *pb.ListReplyRequest) (*pb.ListReplyReply, error) {
	return &pb.ListReplyReply{}, nil
}
