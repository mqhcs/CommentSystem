package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Reply struct {
	ID        int64
	Message   string
	Ats       string
	IP        int64
	Plat      int8
	Device    string
	Version   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Topics    string
	Addr      string
}

type ReplyRepo interface {
	// db
	ListReplyContent(ctx context.Context) ([]*Reply, error)
	GetReplyContent(ctx context.Context, id int64) (*Reply, error)
	CreateReplyContent(ctx context.Context, reply *Reply) error
	UpdateReplyContent(ctx context.Context, id int64, reply *Reply) error
	DeleteReplyContent(ctx context.Context, id int64) error
}

type ReplyUsecase struct {
	repo ReplyRepo
}

func NewReplyUsecase(repo ReplyRepo, logger log.Logger) *ReplyUsecase {
	return &ReplyUsecase{repo: repo}
}

func (uc *ReplyUsecase) List(ctx context.Context) (ps []*Reply, err error) {
	ps, err = uc.repo.ListReplyContent(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyUsecase) Get(ctx context.Context, id int64) (p *Reply, err error) {
	p, err = uc.repo.GetReplyContent(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyUsecase) Create(ctx context.Context, reply *Reply) error {
	return uc.repo.CreateReplyContent(ctx, reply)
}

func (uc *ReplyUsecase) Update(ctx context.Context, id int64, reply *Reply) error {
	return uc.repo.UpdateReplyContent(ctx, id, reply)
}

func (uc *ReplyUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteReplyContent(ctx, id)
}
