package service

import (
	pb "comment-main/api/commentService/v1"
	"comment-main/app/comment-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCommentServiceService)

type CommentServiceService struct {
	pb.UnimplementedCommentServiceServer
	ru  *biz.ReplyUsecase
	riu *biz.ReplyIndexUsecase
	rau *biz.ReplyAreaUsecase
	log *log.Helper
}

func NewCommentServiceService(ru *biz.ReplyUsecase, riu *biz.ReplyIndexUsecase, rau *biz.ReplyAreaUsecase,
	logger log.Logger) *CommentServiceService {
	return &CommentServiceService{
		ru:  ru,
		riu: riu,
		rau: rau,
		log: log.NewHelper(logger),
	}
}
