// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/predicate"
	"comment-main/app/comment-service/internal/data/ent/replyarea"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyAreaUpdate is the builder for updating ReplyArea entities.
type ReplyAreaUpdate struct {
	config
	hooks    []Hook
	mutation *ReplyAreaMutation
}

// Where appends a list predicates to the ReplyAreaUpdate builder.
func (rau *ReplyAreaUpdate) Where(ps ...predicate.ReplyArea) *ReplyAreaUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetOid sets the "oid" field.
func (rau *ReplyAreaUpdate) SetOid(i int64) *ReplyAreaUpdate {
	rau.mutation.ResetOid()
	rau.mutation.SetOid(i)
	return rau
}

// AddOid adds i to the "oid" field.
func (rau *ReplyAreaUpdate) AddOid(i int64) *ReplyAreaUpdate {
	rau.mutation.AddOid(i)
	return rau
}

// SetType sets the "type" field.
func (rau *ReplyAreaUpdate) SetType(i int8) *ReplyAreaUpdate {
	rau.mutation.ResetType()
	rau.mutation.SetType(i)
	return rau
}

// AddType adds i to the "type" field.
func (rau *ReplyAreaUpdate) AddType(i int8) *ReplyAreaUpdate {
	rau.mutation.AddType(i)
	return rau
}

// SetMid sets the "mid" field.
func (rau *ReplyAreaUpdate) SetMid(i int64) *ReplyAreaUpdate {
	rau.mutation.ResetMid()
	rau.mutation.SetMid(i)
	return rau
}

// AddMid adds i to the "mid" field.
func (rau *ReplyAreaUpdate) AddMid(i int64) *ReplyAreaUpdate {
	rau.mutation.AddMid(i)
	return rau
}

// SetCount sets the "count" field.
func (rau *ReplyAreaUpdate) SetCount(i int32) *ReplyAreaUpdate {
	rau.mutation.ResetCount()
	rau.mutation.SetCount(i)
	return rau
}

// AddCount adds i to the "count" field.
func (rau *ReplyAreaUpdate) AddCount(i int32) *ReplyAreaUpdate {
	rau.mutation.AddCount(i)
	return rau
}

// SetRootCount sets the "root_count" field.
func (rau *ReplyAreaUpdate) SetRootCount(i int32) *ReplyAreaUpdate {
	rau.mutation.ResetRootCount()
	rau.mutation.SetRootCount(i)
	return rau
}

// AddRootCount adds i to the "root_count" field.
func (rau *ReplyAreaUpdate) AddRootCount(i int32) *ReplyAreaUpdate {
	rau.mutation.AddRootCount(i)
	return rau
}

// SetAcount sets the "acount" field.
func (rau *ReplyAreaUpdate) SetAcount(i int32) *ReplyAreaUpdate {
	rau.mutation.ResetAcount()
	rau.mutation.SetAcount(i)
	return rau
}

// AddAcount adds i to the "acount" field.
func (rau *ReplyAreaUpdate) AddAcount(i int32) *ReplyAreaUpdate {
	rau.mutation.AddAcount(i)
	return rau
}

// SetState sets the "state" field.
func (rau *ReplyAreaUpdate) SetState(i int8) *ReplyAreaUpdate {
	rau.mutation.ResetState()
	rau.mutation.SetState(i)
	return rau
}

// AddState adds i to the "state" field.
func (rau *ReplyAreaUpdate) AddState(i int8) *ReplyAreaUpdate {
	rau.mutation.AddState(i)
	return rau
}

// SetCtime sets the "ctime" field.
func (rau *ReplyAreaUpdate) SetCtime(t time.Time) *ReplyAreaUpdate {
	rau.mutation.SetCtime(t)
	return rau
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (rau *ReplyAreaUpdate) SetNillableCtime(t *time.Time) *ReplyAreaUpdate {
	if t != nil {
		rau.SetCtime(*t)
	}
	return rau
}

// SetMtime sets the "mtime" field.
func (rau *ReplyAreaUpdate) SetMtime(t time.Time) *ReplyAreaUpdate {
	rau.mutation.SetMtime(t)
	return rau
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (rau *ReplyAreaUpdate) SetNillableMtime(t *time.Time) *ReplyAreaUpdate {
	if t != nil {
		rau.SetMtime(*t)
	}
	return rau
}

// SetAttrs sets the "attrs" field.
func (rau *ReplyAreaUpdate) SetAttrs(i int32) *ReplyAreaUpdate {
	rau.mutation.ResetAttrs()
	rau.mutation.SetAttrs(i)
	return rau
}

// AddAttrs adds i to the "attrs" field.
func (rau *ReplyAreaUpdate) AddAttrs(i int32) *ReplyAreaUpdate {
	rau.mutation.AddAttrs(i)
	return rau
}

// SetMeta sets the "meta" field.
func (rau *ReplyAreaUpdate) SetMeta(s string) *ReplyAreaUpdate {
	rau.mutation.SetMeta(s)
	return rau
}

// Mutation returns the ReplyAreaMutation object of the builder.
func (rau *ReplyAreaUpdate) Mutation() *ReplyAreaMutation {
	return rau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *ReplyAreaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *ReplyAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *ReplyAreaUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *ReplyAreaUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rau *ReplyAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(replyarea.Table, replyarea.Columns, sqlgraph.NewFieldSpec(replyarea.FieldID, field.TypeInt64))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.Oid(); ok {
		_spec.SetField(replyarea.FieldOid, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedOid(); ok {
		_spec.AddField(replyarea.FieldOid, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.GetType(); ok {
		_spec.SetField(replyarea.FieldType, field.TypeInt8, value)
	}
	if value, ok := rau.mutation.AddedType(); ok {
		_spec.AddField(replyarea.FieldType, field.TypeInt8, value)
	}
	if value, ok := rau.mutation.Mid(); ok {
		_spec.SetField(replyarea.FieldMid, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedMid(); ok {
		_spec.AddField(replyarea.FieldMid, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.Count(); ok {
		_spec.SetField(replyarea.FieldCount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.AddedCount(); ok {
		_spec.AddField(replyarea.FieldCount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.RootCount(); ok {
		_spec.SetField(replyarea.FieldRootCount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.AddedRootCount(); ok {
		_spec.AddField(replyarea.FieldRootCount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.Acount(); ok {
		_spec.SetField(replyarea.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.AddedAcount(); ok {
		_spec.AddField(replyarea.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.State(); ok {
		_spec.SetField(replyarea.FieldState, field.TypeInt8, value)
	}
	if value, ok := rau.mutation.AddedState(); ok {
		_spec.AddField(replyarea.FieldState, field.TypeInt8, value)
	}
	if value, ok := rau.mutation.Ctime(); ok {
		_spec.SetField(replyarea.FieldCtime, field.TypeTime, value)
	}
	if value, ok := rau.mutation.Mtime(); ok {
		_spec.SetField(replyarea.FieldMtime, field.TypeTime, value)
	}
	if value, ok := rau.mutation.Attrs(); ok {
		_spec.SetField(replyarea.FieldAttrs, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.AddedAttrs(); ok {
		_spec.AddField(replyarea.FieldAttrs, field.TypeInt32, value)
	}
	if value, ok := rau.mutation.Meta(); ok {
		_spec.SetField(replyarea.FieldMeta, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{replyarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// ReplyAreaUpdateOne is the builder for updating a single ReplyArea entity.
type ReplyAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReplyAreaMutation
}

// SetOid sets the "oid" field.
func (rauo *ReplyAreaUpdateOne) SetOid(i int64) *ReplyAreaUpdateOne {
	rauo.mutation.ResetOid()
	rauo.mutation.SetOid(i)
	return rauo
}

// AddOid adds i to the "oid" field.
func (rauo *ReplyAreaUpdateOne) AddOid(i int64) *ReplyAreaUpdateOne {
	rauo.mutation.AddOid(i)
	return rauo
}

// SetType sets the "type" field.
func (rauo *ReplyAreaUpdateOne) SetType(i int8) *ReplyAreaUpdateOne {
	rauo.mutation.ResetType()
	rauo.mutation.SetType(i)
	return rauo
}

// AddType adds i to the "type" field.
func (rauo *ReplyAreaUpdateOne) AddType(i int8) *ReplyAreaUpdateOne {
	rauo.mutation.AddType(i)
	return rauo
}

// SetMid sets the "mid" field.
func (rauo *ReplyAreaUpdateOne) SetMid(i int64) *ReplyAreaUpdateOne {
	rauo.mutation.ResetMid()
	rauo.mutation.SetMid(i)
	return rauo
}

// AddMid adds i to the "mid" field.
func (rauo *ReplyAreaUpdateOne) AddMid(i int64) *ReplyAreaUpdateOne {
	rauo.mutation.AddMid(i)
	return rauo
}

// SetCount sets the "count" field.
func (rauo *ReplyAreaUpdateOne) SetCount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.ResetCount()
	rauo.mutation.SetCount(i)
	return rauo
}

// AddCount adds i to the "count" field.
func (rauo *ReplyAreaUpdateOne) AddCount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.AddCount(i)
	return rauo
}

// SetRootCount sets the "root_count" field.
func (rauo *ReplyAreaUpdateOne) SetRootCount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.ResetRootCount()
	rauo.mutation.SetRootCount(i)
	return rauo
}

// AddRootCount adds i to the "root_count" field.
func (rauo *ReplyAreaUpdateOne) AddRootCount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.AddRootCount(i)
	return rauo
}

// SetAcount sets the "acount" field.
func (rauo *ReplyAreaUpdateOne) SetAcount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.ResetAcount()
	rauo.mutation.SetAcount(i)
	return rauo
}

// AddAcount adds i to the "acount" field.
func (rauo *ReplyAreaUpdateOne) AddAcount(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.AddAcount(i)
	return rauo
}

// SetState sets the "state" field.
func (rauo *ReplyAreaUpdateOne) SetState(i int8) *ReplyAreaUpdateOne {
	rauo.mutation.ResetState()
	rauo.mutation.SetState(i)
	return rauo
}

// AddState adds i to the "state" field.
func (rauo *ReplyAreaUpdateOne) AddState(i int8) *ReplyAreaUpdateOne {
	rauo.mutation.AddState(i)
	return rauo
}

// SetCtime sets the "ctime" field.
func (rauo *ReplyAreaUpdateOne) SetCtime(t time.Time) *ReplyAreaUpdateOne {
	rauo.mutation.SetCtime(t)
	return rauo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (rauo *ReplyAreaUpdateOne) SetNillableCtime(t *time.Time) *ReplyAreaUpdateOne {
	if t != nil {
		rauo.SetCtime(*t)
	}
	return rauo
}

// SetMtime sets the "mtime" field.
func (rauo *ReplyAreaUpdateOne) SetMtime(t time.Time) *ReplyAreaUpdateOne {
	rauo.mutation.SetMtime(t)
	return rauo
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (rauo *ReplyAreaUpdateOne) SetNillableMtime(t *time.Time) *ReplyAreaUpdateOne {
	if t != nil {
		rauo.SetMtime(*t)
	}
	return rauo
}

// SetAttrs sets the "attrs" field.
func (rauo *ReplyAreaUpdateOne) SetAttrs(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.ResetAttrs()
	rauo.mutation.SetAttrs(i)
	return rauo
}

// AddAttrs adds i to the "attrs" field.
func (rauo *ReplyAreaUpdateOne) AddAttrs(i int32) *ReplyAreaUpdateOne {
	rauo.mutation.AddAttrs(i)
	return rauo
}

// SetMeta sets the "meta" field.
func (rauo *ReplyAreaUpdateOne) SetMeta(s string) *ReplyAreaUpdateOne {
	rauo.mutation.SetMeta(s)
	return rauo
}

// Mutation returns the ReplyAreaMutation object of the builder.
func (rauo *ReplyAreaUpdateOne) Mutation() *ReplyAreaMutation {
	return rauo.mutation
}

// Where appends a list predicates to the ReplyAreaUpdate builder.
func (rauo *ReplyAreaUpdateOne) Where(ps ...predicate.ReplyArea) *ReplyAreaUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *ReplyAreaUpdateOne) Select(field string, fields ...string) *ReplyAreaUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated ReplyArea entity.
func (rauo *ReplyAreaUpdateOne) Save(ctx context.Context) (*ReplyArea, error) {
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *ReplyAreaUpdateOne) SaveX(ctx context.Context) *ReplyArea {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *ReplyAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *ReplyAreaUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rauo *ReplyAreaUpdateOne) sqlSave(ctx context.Context) (_node *ReplyArea, err error) {
	_spec := sqlgraph.NewUpdateSpec(replyarea.Table, replyarea.Columns, sqlgraph.NewFieldSpec(replyarea.FieldID, field.TypeInt64))
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReplyArea.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, replyarea.FieldID)
		for _, f := range fields {
			if !replyarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != replyarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.Oid(); ok {
		_spec.SetField(replyarea.FieldOid, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedOid(); ok {
		_spec.AddField(replyarea.FieldOid, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.GetType(); ok {
		_spec.SetField(replyarea.FieldType, field.TypeInt8, value)
	}
	if value, ok := rauo.mutation.AddedType(); ok {
		_spec.AddField(replyarea.FieldType, field.TypeInt8, value)
	}
	if value, ok := rauo.mutation.Mid(); ok {
		_spec.SetField(replyarea.FieldMid, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedMid(); ok {
		_spec.AddField(replyarea.FieldMid, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.Count(); ok {
		_spec.SetField(replyarea.FieldCount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.AddedCount(); ok {
		_spec.AddField(replyarea.FieldCount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.RootCount(); ok {
		_spec.SetField(replyarea.FieldRootCount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.AddedRootCount(); ok {
		_spec.AddField(replyarea.FieldRootCount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.Acount(); ok {
		_spec.SetField(replyarea.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.AddedAcount(); ok {
		_spec.AddField(replyarea.FieldAcount, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.State(); ok {
		_spec.SetField(replyarea.FieldState, field.TypeInt8, value)
	}
	if value, ok := rauo.mutation.AddedState(); ok {
		_spec.AddField(replyarea.FieldState, field.TypeInt8, value)
	}
	if value, ok := rauo.mutation.Ctime(); ok {
		_spec.SetField(replyarea.FieldCtime, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.Mtime(); ok {
		_spec.SetField(replyarea.FieldMtime, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.Attrs(); ok {
		_spec.SetField(replyarea.FieldAttrs, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.AddedAttrs(); ok {
		_spec.AddField(replyarea.FieldAttrs, field.TypeInt32, value)
	}
	if value, ok := rauo.mutation.Meta(); ok {
		_spec.SetField(replyarea.FieldMeta, field.TypeString, value)
	}
	_node = &ReplyArea{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{replyarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}