package data

import (
	"context"
	"time"

	"comment-main/app/comment-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type replyRepo struct {
	data *Data
	log  *log.Helper
}

func NewReplyRepo(data *Data, logger log.Logger) biz.ReplyRepo {
	return &replyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (rr *replyRepo) ListReplyContent(ctx context.Context) ([]*biz.Reply, error) {
	ps, err := rr.data.db.Reply.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Reply, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Reply{
			ID:      int64(p.ID),
			Message: p.Message,
			Ats:     p.Ats,
			IP:      p.IP,
			Plat:    p.Plat,
			Device:  p.Device,
			Version: p.Version,
			Ctime:   p.Ctime,
			Mtime:   p.Mtime,
			Topics:  p.Topics,
			Addr:    p.Addr,
		})
	}
	return rv, nil
}

func (rr *replyRepo) GetReplyContent(ctx context.Context, id int64) (*biz.Reply, error) {
	p, err := rr.data.db.Reply.Get(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return &biz.Reply{
		ID:      int64(p.ID),
		Message: p.Message,
		Ats:     p.Ats,
		IP:      p.IP,
		Plat:    p.Plat,
		Device:  p.Device,
		Version: p.Version,
		Ctime:   p.Ctime,
		Mtime:   p.Mtime,
		Topics:  p.Topics,
		Addr:    p.Addr,
	}, nil
}

func (rr *replyRepo) CreateReplyContent(ctx context.Context, reply *biz.Reply) error {
	_, err := rr.data.db.Reply.
		Create().
		SetIP(reply.IP).
		SetMessage(reply.Message).
		Save(ctx)
	return err
}

func (rr *replyRepo) UpdateReplyContent(ctx context.Context, id int64, reply *biz.Reply) error {
	p, err := rr.data.db.Reply.Get(ctx, int(id))
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetIP(reply.IP).
		SetMessage(reply.Message).
		SetMtime(time.Now()).
		Save(ctx)
	return err
}

func (rr *replyRepo) DeleteReplyContent(ctx context.Context, id int64) error {
	return rr.data.db.Reply.DeleteOneID(int(id)).Exec(ctx)
}
