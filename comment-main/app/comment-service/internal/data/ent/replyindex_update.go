// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/predicate"
	"comment-main/app/comment-service/internal/data/ent/replyindex"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyIndexUpdate is the builder for updating ReplyIndex entities.
type ReplyIndexUpdate struct {
	config
	hooks    []Hook
	mutation *ReplyIndexMutation
}

// Where appends a list predicates to the ReplyIndexUpdate builder.
func (riu *ReplyIndexUpdate) Where(ps ...predicate.ReplyIndex) *ReplyIndexUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// SetOid sets the "oid" field.
func (riu *ReplyIndexUpdate) SetOid(i int64) *ReplyIndexUpdate {
	riu.mutation.ResetOid()
	riu.mutation.SetOid(i)
	return riu
}

// AddOid adds i to the "oid" field.
func (riu *ReplyIndexUpdate) AddOid(i int64) *ReplyIndexUpdate {
	riu.mutation.AddOid(i)
	return riu
}

// SetType sets the "type" field.
func (riu *ReplyIndexUpdate) SetType(i int8) *ReplyIndexUpdate {
	riu.mutation.ResetType()
	riu.mutation.SetType(i)
	return riu
}

// AddType adds i to the "type" field.
func (riu *ReplyIndexUpdate) AddType(i int8) *ReplyIndexUpdate {
	riu.mutation.AddType(i)
	return riu
}

// SetMid sets the "mid" field.
func (riu *ReplyIndexUpdate) SetMid(i int64) *ReplyIndexUpdate {
	riu.mutation.ResetMid()
	riu.mutation.SetMid(i)
	return riu
}

// AddMid adds i to the "mid" field.
func (riu *ReplyIndexUpdate) AddMid(i int64) *ReplyIndexUpdate {
	riu.mutation.AddMid(i)
	return riu
}

// SetRoot sets the "root" field.
func (riu *ReplyIndexUpdate) SetRoot(i int64) *ReplyIndexUpdate {
	riu.mutation.ResetRoot()
	riu.mutation.SetRoot(i)
	return riu
}

// AddRoot adds i to the "root" field.
func (riu *ReplyIndexUpdate) AddRoot(i int64) *ReplyIndexUpdate {
	riu.mutation.AddRoot(i)
	return riu
}

// SetParent sets the "parent" field.
func (riu *ReplyIndexUpdate) SetParent(i int64) *ReplyIndexUpdate {
	riu.mutation.ResetParent()
	riu.mutation.SetParent(i)
	return riu
}

// AddParent adds i to the "parent" field.
func (riu *ReplyIndexUpdate) AddParent(i int64) *ReplyIndexUpdate {
	riu.mutation.AddParent(i)
	return riu
}

// SetCount sets the "count" field.
func (riu *ReplyIndexUpdate) SetCount(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetCount()
	riu.mutation.SetCount(i)
	return riu
}

// AddCount adds i to the "count" field.
func (riu *ReplyIndexUpdate) AddCount(i int32) *ReplyIndexUpdate {
	riu.mutation.AddCount(i)
	return riu
}

// SetRcount sets the "rcount" field.
func (riu *ReplyIndexUpdate) SetRcount(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetRcount()
	riu.mutation.SetRcount(i)
	return riu
}

// AddRcount adds i to the "rcount" field.
func (riu *ReplyIndexUpdate) AddRcount(i int32) *ReplyIndexUpdate {
	riu.mutation.AddRcount(i)
	return riu
}

// SetAcount sets the "acount" field.
func (riu *ReplyIndexUpdate) SetAcount(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetAcount()
	riu.mutation.SetAcount(i)
	return riu
}

// AddAcount adds i to the "acount" field.
func (riu *ReplyIndexUpdate) AddAcount(i int32) *ReplyIndexUpdate {
	riu.mutation.AddAcount(i)
	return riu
}

// SetLikes sets the "likes" field.
func (riu *ReplyIndexUpdate) SetLikes(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetLikes()
	riu.mutation.SetLikes(i)
	return riu
}

// AddLikes adds i to the "likes" field.
func (riu *ReplyIndexUpdate) AddLikes(i int32) *ReplyIndexUpdate {
	riu.mutation.AddLikes(i)
	return riu
}

// SetHates sets the "hates" field.
func (riu *ReplyIndexUpdate) SetHates(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetHates()
	riu.mutation.SetHates(i)
	return riu
}

// AddHates adds i to the "hates" field.
func (riu *ReplyIndexUpdate) AddHates(i int32) *ReplyIndexUpdate {
	riu.mutation.AddHates(i)
	return riu
}

// SetReportCount sets the "report_count" field.
func (riu *ReplyIndexUpdate) SetReportCount(i int32) *ReplyIndexUpdate {
	riu.mutation.ResetReportCount()
	riu.mutation.SetReportCount(i)
	return riu
}

// AddReportCount adds i to the "report_count" field.
func (riu *ReplyIndexUpdate) AddReportCount(i int32) *ReplyIndexUpdate {
	riu.mutation.AddReportCount(i)
	return riu
}

// SetState sets the "state" field.
func (riu *ReplyIndexUpdate) SetState(i int8) *ReplyIndexUpdate {
	riu.mutation.ResetState()
	riu.mutation.SetState(i)
	return riu
}

// AddState adds i to the "state" field.
func (riu *ReplyIndexUpdate) AddState(i int8) *ReplyIndexUpdate {
	riu.mutation.AddState(i)
	return riu
}

// SetCtime sets the "ctime" field.
func (riu *ReplyIndexUpdate) SetCtime(t time.Time) *ReplyIndexUpdate {
	riu.mutation.SetCtime(t)
	return riu
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (riu *ReplyIndexUpdate) SetNillableCtime(t *time.Time) *ReplyIndexUpdate {
	if t != nil {
		riu.SetCtime(*t)
	}
	return riu
}

// SetMtime sets the "mtime" field.
func (riu *ReplyIndexUpdate) SetMtime(t time.Time) *ReplyIndexUpdate {
	riu.mutation.SetMtime(t)
	return riu
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (riu *ReplyIndexUpdate) SetNillableMtime(t *time.Time) *ReplyIndexUpdate {
	if t != nil {
		riu.SetMtime(*t)
	}
	return riu
}

// Mutation returns the ReplyIndexMutation object of the builder.
func (riu *ReplyIndexUpdate) Mutation() *ReplyIndexMutation {
	return riu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *ReplyIndexUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, riu.sqlSave, riu.mutation, riu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riu *ReplyIndexUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *ReplyIndexUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *ReplyIndexUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riu *ReplyIndexUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(replyindex.Table, replyindex.Columns, sqlgraph.NewFieldSpec(replyindex.FieldID, field.TypeInt64))
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riu.mutation.Oid(); ok {
		_spec.SetField(replyindex.FieldOid, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.AddedOid(); ok {
		_spec.AddField(replyindex.FieldOid, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.GetType(); ok {
		_spec.SetField(replyindex.FieldType, field.TypeInt8, value)
	}
	if value, ok := riu.mutation.AddedType(); ok {
		_spec.AddField(replyindex.FieldType, field.TypeInt8, value)
	}
	if value, ok := riu.mutation.Mid(); ok {
		_spec.SetField(replyindex.FieldMid, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.AddedMid(); ok {
		_spec.AddField(replyindex.FieldMid, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.Root(); ok {
		_spec.SetField(replyindex.FieldRoot, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.AddedRoot(); ok {
		_spec.AddField(replyindex.FieldRoot, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.Parent(); ok {
		_spec.SetField(replyindex.FieldParent, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.AddedParent(); ok {
		_spec.AddField(replyindex.FieldParent, field.TypeInt64, value)
	}
	if value, ok := riu.mutation.Count(); ok {
		_spec.SetField(replyindex.FieldCount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedCount(); ok {
		_spec.AddField(replyindex.FieldCount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.Rcount(); ok {
		_spec.SetField(replyindex.FieldRcount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedRcount(); ok {
		_spec.AddField(replyindex.FieldRcount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.Acount(); ok {
		_spec.SetField(replyindex.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedAcount(); ok {
		_spec.AddField(replyindex.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.Likes(); ok {
		_spec.SetField(replyindex.FieldLikes, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedLikes(); ok {
		_spec.AddField(replyindex.FieldLikes, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.Hates(); ok {
		_spec.SetField(replyindex.FieldHates, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedHates(); ok {
		_spec.AddField(replyindex.FieldHates, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.ReportCount(); ok {
		_spec.SetField(replyindex.FieldReportCount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.AddedReportCount(); ok {
		_spec.AddField(replyindex.FieldReportCount, field.TypeInt32, value)
	}
	if value, ok := riu.mutation.State(); ok {
		_spec.SetField(replyindex.FieldState, field.TypeInt8, value)
	}
	if value, ok := riu.mutation.AddedState(); ok {
		_spec.AddField(replyindex.FieldState, field.TypeInt8, value)
	}
	if value, ok := riu.mutation.Ctime(); ok {
		_spec.SetField(replyindex.FieldCtime, field.TypeTime, value)
	}
	if value, ok := riu.mutation.Mtime(); ok {
		_spec.SetField(replyindex.FieldMtime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{replyindex.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	riu.mutation.done = true
	return n, nil
}

// ReplyIndexUpdateOne is the builder for updating a single ReplyIndex entity.
type ReplyIndexUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReplyIndexMutation
}

// SetOid sets the "oid" field.
func (riuo *ReplyIndexUpdateOne) SetOid(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.ResetOid()
	riuo.mutation.SetOid(i)
	return riuo
}

// AddOid adds i to the "oid" field.
func (riuo *ReplyIndexUpdateOne) AddOid(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.AddOid(i)
	return riuo
}

// SetType sets the "type" field.
func (riuo *ReplyIndexUpdateOne) SetType(i int8) *ReplyIndexUpdateOne {
	riuo.mutation.ResetType()
	riuo.mutation.SetType(i)
	return riuo
}

// AddType adds i to the "type" field.
func (riuo *ReplyIndexUpdateOne) AddType(i int8) *ReplyIndexUpdateOne {
	riuo.mutation.AddType(i)
	return riuo
}

// SetMid sets the "mid" field.
func (riuo *ReplyIndexUpdateOne) SetMid(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.ResetMid()
	riuo.mutation.SetMid(i)
	return riuo
}

// AddMid adds i to the "mid" field.
func (riuo *ReplyIndexUpdateOne) AddMid(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.AddMid(i)
	return riuo
}

// SetRoot sets the "root" field.
func (riuo *ReplyIndexUpdateOne) SetRoot(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.ResetRoot()
	riuo.mutation.SetRoot(i)
	return riuo
}

// AddRoot adds i to the "root" field.
func (riuo *ReplyIndexUpdateOne) AddRoot(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.AddRoot(i)
	return riuo
}

// SetParent sets the "parent" field.
func (riuo *ReplyIndexUpdateOne) SetParent(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.ResetParent()
	riuo.mutation.SetParent(i)
	return riuo
}

// AddParent adds i to the "parent" field.
func (riuo *ReplyIndexUpdateOne) AddParent(i int64) *ReplyIndexUpdateOne {
	riuo.mutation.AddParent(i)
	return riuo
}

// SetCount sets the "count" field.
func (riuo *ReplyIndexUpdateOne) SetCount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetCount()
	riuo.mutation.SetCount(i)
	return riuo
}

// AddCount adds i to the "count" field.
func (riuo *ReplyIndexUpdateOne) AddCount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddCount(i)
	return riuo
}

// SetRcount sets the "rcount" field.
func (riuo *ReplyIndexUpdateOne) SetRcount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetRcount()
	riuo.mutation.SetRcount(i)
	return riuo
}

// AddRcount adds i to the "rcount" field.
func (riuo *ReplyIndexUpdateOne) AddRcount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddRcount(i)
	return riuo
}

// SetAcount sets the "acount" field.
func (riuo *ReplyIndexUpdateOne) SetAcount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetAcount()
	riuo.mutation.SetAcount(i)
	return riuo
}

// AddAcount adds i to the "acount" field.
func (riuo *ReplyIndexUpdateOne) AddAcount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddAcount(i)
	return riuo
}

// SetLikes sets the "likes" field.
func (riuo *ReplyIndexUpdateOne) SetLikes(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetLikes()
	riuo.mutation.SetLikes(i)
	return riuo
}

// AddLikes adds i to the "likes" field.
func (riuo *ReplyIndexUpdateOne) AddLikes(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddLikes(i)
	return riuo
}

// SetHates sets the "hates" field.
func (riuo *ReplyIndexUpdateOne) SetHates(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetHates()
	riuo.mutation.SetHates(i)
	return riuo
}

// AddHates adds i to the "hates" field.
func (riuo *ReplyIndexUpdateOne) AddHates(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddHates(i)
	return riuo
}

// SetReportCount sets the "report_count" field.
func (riuo *ReplyIndexUpdateOne) SetReportCount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.ResetReportCount()
	riuo.mutation.SetReportCount(i)
	return riuo
}

// AddReportCount adds i to the "report_count" field.
func (riuo *ReplyIndexUpdateOne) AddReportCount(i int32) *ReplyIndexUpdateOne {
	riuo.mutation.AddReportCount(i)
	return riuo
}

// SetState sets the "state" field.
func (riuo *ReplyIndexUpdateOne) SetState(i int8) *ReplyIndexUpdateOne {
	riuo.mutation.ResetState()
	riuo.mutation.SetState(i)
	return riuo
}

// AddState adds i to the "state" field.
func (riuo *ReplyIndexUpdateOne) AddState(i int8) *ReplyIndexUpdateOne {
	riuo.mutation.AddState(i)
	return riuo
}

// SetCtime sets the "ctime" field.
func (riuo *ReplyIndexUpdateOne) SetCtime(t time.Time) *ReplyIndexUpdateOne {
	riuo.mutation.SetCtime(t)
	return riuo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (riuo *ReplyIndexUpdateOne) SetNillableCtime(t *time.Time) *ReplyIndexUpdateOne {
	if t != nil {
		riuo.SetCtime(*t)
	}
	return riuo
}

// SetMtime sets the "mtime" field.
func (riuo *ReplyIndexUpdateOne) SetMtime(t time.Time) *ReplyIndexUpdateOne {
	riuo.mutation.SetMtime(t)
	return riuo
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (riuo *ReplyIndexUpdateOne) SetNillableMtime(t *time.Time) *ReplyIndexUpdateOne {
	if t != nil {
		riuo.SetMtime(*t)
	}
	return riuo
}

// Mutation returns the ReplyIndexMutation object of the builder.
func (riuo *ReplyIndexUpdateOne) Mutation() *ReplyIndexMutation {
	return riuo.mutation
}

// Where appends a list predicates to the ReplyIndexUpdate builder.
func (riuo *ReplyIndexUpdateOne) Where(ps ...predicate.ReplyIndex) *ReplyIndexUpdateOne {
	riuo.mutation.Where(ps...)
	return riuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *ReplyIndexUpdateOne) Select(field string, fields ...string) *ReplyIndexUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated ReplyIndex entity.
func (riuo *ReplyIndexUpdateOne) Save(ctx context.Context) (*ReplyIndex, error) {
	return withHooks(ctx, riuo.sqlSave, riuo.mutation, riuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *ReplyIndexUpdateOne) SaveX(ctx context.Context) *ReplyIndex {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *ReplyIndexUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *ReplyIndexUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (riuo *ReplyIndexUpdateOne) sqlSave(ctx context.Context) (_node *ReplyIndex, err error) {
	_spec := sqlgraph.NewUpdateSpec(replyindex.Table, replyindex.Columns, sqlgraph.NewFieldSpec(replyindex.FieldID, field.TypeInt64))
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReplyIndex.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, replyindex.FieldID)
		for _, f := range fields {
			if !replyindex.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != replyindex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riuo.mutation.Oid(); ok {
		_spec.SetField(replyindex.FieldOid, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.AddedOid(); ok {
		_spec.AddField(replyindex.FieldOid, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.GetType(); ok {
		_spec.SetField(replyindex.FieldType, field.TypeInt8, value)
	}
	if value, ok := riuo.mutation.AddedType(); ok {
		_spec.AddField(replyindex.FieldType, field.TypeInt8, value)
	}
	if value, ok := riuo.mutation.Mid(); ok {
		_spec.SetField(replyindex.FieldMid, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.AddedMid(); ok {
		_spec.AddField(replyindex.FieldMid, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.Root(); ok {
		_spec.SetField(replyindex.FieldRoot, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.AddedRoot(); ok {
		_spec.AddField(replyindex.FieldRoot, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.Parent(); ok {
		_spec.SetField(replyindex.FieldParent, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.AddedParent(); ok {
		_spec.AddField(replyindex.FieldParent, field.TypeInt64, value)
	}
	if value, ok := riuo.mutation.Count(); ok {
		_spec.SetField(replyindex.FieldCount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedCount(); ok {
		_spec.AddField(replyindex.FieldCount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.Rcount(); ok {
		_spec.SetField(replyindex.FieldRcount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedRcount(); ok {
		_spec.AddField(replyindex.FieldRcount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.Acount(); ok {
		_spec.SetField(replyindex.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedAcount(); ok {
		_spec.AddField(replyindex.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.Likes(); ok {
		_spec.SetField(replyindex.FieldLikes, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedLikes(); ok {
		_spec.AddField(replyindex.FieldLikes, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.Hates(); ok {
		_spec.SetField(replyindex.FieldHates, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedHates(); ok {
		_spec.AddField(replyindex.FieldHates, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.ReportCount(); ok {
		_spec.SetField(replyindex.FieldReportCount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.AddedReportCount(); ok {
		_spec.AddField(replyindex.FieldReportCount, field.TypeInt32, value)
	}
	if value, ok := riuo.mutation.State(); ok {
		_spec.SetField(replyindex.FieldState, field.TypeInt8, value)
	}
	if value, ok := riuo.mutation.AddedState(); ok {
		_spec.AddField(replyindex.FieldState, field.TypeInt8, value)
	}
	if value, ok := riuo.mutation.Ctime(); ok {
		_spec.SetField(replyindex.FieldCtime, field.TypeTime, value)
	}
	if value, ok := riuo.mutation.Mtime(); ok {
		_spec.SetField(replyindex.FieldMtime, field.TypeTime, value)
	}
	_node = &ReplyIndex{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{replyindex.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	riuo.mutation.done = true
	return _node, nil
}
