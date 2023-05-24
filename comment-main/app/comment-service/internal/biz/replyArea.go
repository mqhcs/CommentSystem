package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ReplyArea struct {
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

type ReplyAreaRepo interface {
	// db
	ListReplyAreaContent(ctx context.Context) ([]*ReplyArea, error)
	GetReplyAreaContent(ctx context.Context, id int64) (*ReplyArea, error)
	CreateReplyAreaContent(ctx context.Context, reply *ReplyArea) error
	UpdateReplyAreaContent(ctx context.Context, id int64, reply *ReplyArea) error
	DeleteReplyAreaContent(ctx context.Context, id int64) error
}

type ReplyAreaUsecase struct {
	repo ReplyAreaRepo
}

func NewReplyAreaUsecase(repo ReplyAreaRepo, logger log.Logger) *ReplyAreaUsecase {
	return &ReplyAreaUsecase{repo: repo}
}

func (uc *ReplyAreaUsecase) List(ctx context.Context) (ps []*ReplyArea, err error) {
	ps, err = uc.repo.ListReplyAreaContent(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyAreaUsecase) Get(ctx context.Context, id int64) (p *ReplyArea, err error) {
	p, err = uc.repo.GetReplyAreaContent(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ReplyAreaUsecase) Create(ctx context.Context, reply *ReplyArea) error {
	return uc.repo.CreateReplyAreaContent(ctx, reply)
}

func (uc *ReplyAreaUsecase) Update(ctx context.Context, id int64, reply *ReplyArea) error {
	return uc.repo.UpdateReplyAreaContent(ctx, id, reply)
}

func (uc *ReplyAreaUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteReplyAreaContent(ctx, id)
}
