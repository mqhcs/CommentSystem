package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ReplyIndex struct {
	ID      int64
	Message string
	Ats     string
	IP      int64
	Plat    int8
	Device  string
	Version string
	Ctime   time.Time
	Mtime   time.Time
	Topics  string
	Addr    string
}

type ReplyIndexRepo interface {
	// db
	ListReplyIndexContent(ctx context.Context) ([]*ReplyIndex, error)
	GetReplyIndexContent(ctx context.Context, id int64) (*ReplyIndex, error)
	CreateReplyIndexContent(ctx context.Context, reply *ReplyIndex) error
	UpdateReplyIndexContent(ctx context.Context, id int64, reply *ReplyIndex) error
	DeleteReplyIndexContent(ctx context.Context, id int64) error
}

type ReplyIndexUsecase struct {
	repo ReplyIndexRepo
}

func NewReplyIndexUsecase(repo ReplyIndexRepo, logger log.Logger) *ReplyIndexUsecase {
	return &ReplyIndexUsecase{repo: repo}
}

func (uc *ReplyIndexUsecase) List(ctx context.Context) (ps []*ReplyIndex, err error) {
	ps, err = uc.repo.ListReplyIndexContent(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyIndexUsecase) Get(ctx context.Context, id int64) (p *ReplyIndex, err error) {
	p, err = uc.repo.GetReplyIndexContent(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyIndexUsecase) Create(ctx context.Context, reply *ReplyIndex) error {
	return uc.repo.CreateReplyIndexContent(ctx, reply)
}

func (uc *ReplyIndexUsecase) Update(ctx context.Context, id int64, reply *ReplyIndex) error {
	return uc.repo.UpdateReplyIndexContent(ctx, id, reply)
}

func (uc *ReplyIndexUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteReplyIndexContent(ctx, id)
}
