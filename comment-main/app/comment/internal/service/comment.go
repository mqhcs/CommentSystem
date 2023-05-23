package service

import (
	pb "comment-main/api/comment/v1"
	"comment-main/app/comment/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCommentService())

type CommentService struct {
	pb.UnimplementedCommentServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewCommentService() *CommentService {
	return &CommentService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/comment")),
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	return &pb.CreateCommentReply{}, nil
}
func (s *CommentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentReply, error) {
	return &pb.UpdateCommentReply{}, nil
}
func (s *CommentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, nil
}
func (s *CommentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	return &pb.GetCommentReply{}, nil
}
func (s *CommentService) ListComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	return &pb.ListCommentReply{}, nil
}
