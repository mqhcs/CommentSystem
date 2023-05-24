// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/predicate"
	"comment-main/app/comment-service/internal/data/ent/replyindex"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyIndexDelete is the builder for deleting a ReplyIndex entity.
type ReplyIndexDelete struct {
	config
	hooks    []Hook
	mutation *ReplyIndexMutation
}

// Where appends a list predicates to the ReplyIndexDelete builder.
func (rid *ReplyIndexDelete) Where(ps ...predicate.ReplyIndex) *ReplyIndexDelete {
	rid.mutation.Where(ps...)
	return rid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rid *ReplyIndexDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rid.sqlExec, rid.mutation, rid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rid *ReplyIndexDelete) ExecX(ctx context.Context) int {
	n, err := rid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rid *ReplyIndexDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(replyindex.Table, sqlgraph.NewFieldSpec(replyindex.FieldID, field.TypeInt64))
	if ps := rid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rid.mutation.done = true
	return affected, err
}

// ReplyIndexDeleteOne is the builder for deleting a single ReplyIndex entity.
type ReplyIndexDeleteOne struct {
	rid *ReplyIndexDelete
}

// Where appends a list predicates to the ReplyIndexDelete builder.
func (rido *ReplyIndexDeleteOne) Where(ps ...predicate.ReplyIndex) *ReplyIndexDeleteOne {
	rido.rid.mutation.Where(ps...)
	return rido
}

// Exec executes the deletion query.
func (rido *ReplyIndexDeleteOne) Exec(ctx context.Context) error {
	n, err := rido.rid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{replyindex.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rido *ReplyIndexDeleteOne) ExecX(ctx context.Context) {
	if err := rido.Exec(ctx); err != nil {
		panic(err)
	}
}
