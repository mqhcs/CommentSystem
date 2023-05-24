package biz

import (
	serviceV1 "comment-main/api/commentService/v1"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)


type CommentRepo interface {
	DisplayCommentList(ctx context.Context, aid int64, userId int64) (*domain.CommentListDisplay, error)
	SendComment(ctx context.Context, aid int64, userId int64) (*domain.SendCommentEntity, error)
}

type CommentUsecase struct {
	repo       CommentRepo
	serviceRPC serviceV1.CommentServiceClient
	log        *log.Helper
}

func NewCommentUsecase(repo CommentRepo, serviceRPC serviceV1.CommentServiceClient, logger log.Logger) *CommentUsecase {

	return &CommentUsecase{
		repo:       repo,
		serviceRPC: serviceRPC,
		log:        log.NewHelper(logger),
	}
}

func (cc *CommentUsecase) SendComment(ctx context.Context, comment *domain.SendCommentEntity) {
	// 跨服务查询信息
	{
		// 已选中，获取评论区信息
		areaList, err := cc.serviceRPC.ListReplyArea(ctx, &serviceV1.ListReplyAreaRequest{})
		if err != nil {
			fmt.Println(err)
		}

		// 判断评论区是否关闭
		//评论区是否只读
		//判断账号是否禁封
		//答题通过
		//节操值>=60

		//当前up是否拉黑发送者
		//...
		//允许发送
		//跨服务发评论comment-service
		cc.serviceRPC.CreateReply()
		cc.serviceRPC.CreateReplyIndex()
		cc.serviceRPC.CreateReplyArea()
	}

}
func (cc *CommentUsecase) DisplayCommentList(ctx context.Context, comment *domain.CommentListDisplay) {
	// 跨服务查询信息
	{

		// 已选中，根据OID，查询对应的评论区
		areaList, err := cc.serviceRPC.ListReplyArea(ctx, &serviceV1.ListReplyAreaRequest{})
		if err != nil {
			fmt.Println(err)
		}




}
