// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/replyarea"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReplyAreaCreate is the builder for creating a ReplyArea entity.
type ReplyAreaCreate struct {
	config
	mutation *ReplyAreaMutation
	hooks    []Hook
}

// SetOid sets the "oid" field.
func (rac *ReplyAreaCreate) SetOid(i int64) *ReplyAreaCreate {
	rac.mutation.SetOid(i)
	return rac
}

// SetType sets the "type" field.
func (rac *ReplyAreaCreate) SetType(i int8) *ReplyAreaCreate {
	rac.mutation.SetType(i)
	return rac
}

// SetMid sets the "mid" field.
func (rac *ReplyAreaCreate) SetMid(i int64) *ReplyAreaCreate {
	rac.mutation.SetMid(i)
	return rac
}

// SetCount sets the "count" field.
func (rac *ReplyAreaCreate) SetCount(i int32) *ReplyAreaCreate {
	rac.mutation.SetCount(i)
	return rac
}

// SetRootCount sets the "root_count" field.
func (rac *ReplyAreaCreate) SetRootCount(i int32) *ReplyAreaCreate {
	rac.mutation.SetRootCount(i)
	return rac
}

// SetAcount sets the "acount" field.
func (rac *ReplyAreaCreate) SetAcount(i int32) *ReplyAreaCreate {
	rac.mutation.SetAcount(i)
	return rac
}

// SetState sets the "state" field.
func (rac *ReplyAreaCreate) SetState(i int8) *ReplyAreaCreate {
	rac.mutation.SetState(i)
	return rac
}

// SetCtime sets the "ctime" field.
func (rac *ReplyAreaCreate) SetCtime(t time.Time) *ReplyAreaCreate {
	rac.mutation.SetCtime(t)
	return rac
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (rac *ReplyAreaCreate) SetNillableCtime(t *time.Time) *ReplyAreaCreate {
	if t != nil {
		rac.SetCtime(*t)
	}
	return rac
}

// SetMtime sets the "mtime" field.
func (rac *ReplyAreaCreate) SetMtime(t time.Time) *ReplyAreaCreate {
	rac.mutation.SetMtime(t)
	return rac
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (rac *ReplyAreaCreate) SetNillableMtime(t *time.Time) *ReplyAreaCreate {
	if t != nil {
		rac.SetMtime(*t)
	}
	return rac
}

// SetAttrs sets the "attrs" field.
func (rac *ReplyAreaCreate) SetAttrs(i int32) *ReplyAreaCreate {
	rac.mutation.SetAttrs(i)
	return rac
}

// SetMeta sets the "meta" field.
func (rac *ReplyAreaCreate) SetMeta(s string) *ReplyAreaCreate {
	rac.mutation.SetMeta(s)
	return rac
}

// SetID sets the "id" field.
func (rac *ReplyAreaCreate) SetID(i int64) *ReplyAreaCreate {
	rac.mutation.SetID(i)
	return rac
}

// Mutation returns the ReplyAreaMutation object of the builder.
func (rac *ReplyAreaCreate) Mutation() *ReplyAreaMutation {
	return rac.mutation
}

// Save creates the ReplyArea in the database.
func (rac *ReplyAreaCreate) Save(ctx context.Context) (*ReplyArea, error) {
	rac.defaults()
	return withHooks(ctx, rac.sqlSave, rac.mutation, rac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rac *ReplyAreaCreate) SaveX(ctx context.Context) *ReplyArea {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *ReplyAreaCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *ReplyAreaCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rac *ReplyAreaCreate) defaults() {
	if _, ok := rac.mutation.Ctime(); !ok {
		v := replyarea.DefaultCtime()
		rac.mutation.SetCtime(v)
	}
	if _, ok := rac.mutation.Mtime(); !ok {
		v := replyarea.DefaultMtime()
		rac.mutation.SetMtime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *ReplyAreaCreate) check() error {
	if _, ok := rac.mutation.Oid(); !ok {
		return &ValidationError{Name: "oid", err: errors.New(`ent: missing required field "ReplyArea.oid"`)}
	}
	if _, ok := rac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ReplyArea.type"`)}
	}
	if _, ok := rac.mutation.Mid(); !ok {
		return &ValidationError{Name: "mid", err: errors.New(`ent: missing required field "ReplyArea.mid"`)}
	}
	if _, ok := rac.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "ReplyArea.count"`)}
	}
	if _, ok := rac.mutation.RootCount(); !ok {
		return &ValidationError{Name: "root_count", err: errors.New(`ent: missing required field "ReplyArea.root_count"`)}
	}
	if _, ok := rac.mutation.Acount(); !ok {
		return &ValidationError{Name: "acount", err: errors.New(`ent: missing required field "ReplyArea.acount"`)}
	}
	if _, ok := rac.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "ReplyArea.state"`)}
	}
	if _, ok := rac.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "ReplyArea.ctime"`)}
	}
	if _, ok := rac.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New(`ent: missing required field "ReplyArea.mtime"`)}
	}
	if _, ok := rac.mutation.Attrs(); !ok {
		return &ValidationError{Name: "attrs", err: errors.New(`ent: missing required field "ReplyArea.attrs"`)}
	}
	if _, ok := rac.mutation.Meta(); !ok {
		return &ValidationError{Name: "meta", err: errors.New(`ent: missing required field "ReplyArea.meta"`)}
	}
	return nil
}

func (rac *ReplyAreaCreate) sqlSave(ctx context.Context) (*ReplyArea, error) {
	if err := rac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rac.mutation.id = &_node.ID
	rac.mutation.done = true
	return _node, nil
}

func (rac *ReplyAreaCreate) createSpec() (*ReplyArea, *sqlgraph.CreateSpec) {
	var (
		_node = &ReplyArea{config: rac.config}
		_spec = sqlgraph.NewCreateSpec(replyarea.Table, sqlgraph.NewFieldSpec(replyarea.FieldID, field.TypeInt64))
	)
	if id, ok := rac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rac.mutation.Oid(); ok {
		_spec.SetField(replyarea.FieldOid, field.TypeInt64, value)
		_node.Oid = value
	}
	if value, ok := rac.mutation.GetType(); ok {
		_spec.SetField(replyarea.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := rac.mutation.Mid(); ok {
		_spec.SetField(replyarea.FieldMid, field.TypeInt64, value)
		_node.Mid = value
	}
	if value, ok := rac.mutation.Count(); ok {
		_spec.SetField(replyarea.FieldCount, field.TypeInt32, value)
		_node.Count = value
	}
	if value, ok := rac.mutation.RootCount(); ok {
		_spec.SetField(replyarea.FieldRootCount, field.TypeInt32, value)
		_node.RootCount = value
	}
	if value, ok := rac.mutation.Acount(); ok {
		_spec.SetField(replyarea.FieldAcount, field.TypeInt32, value)
		_node.Acount = value
	}
	if value, ok := rac.mutation.State(); ok {
		_spec.SetField(replyarea.FieldState, field.TypeInt8, value)
		_node.State = value
	}
	if value, ok := rac.mutation.Ctime(); ok {
		_spec.SetField(replyarea.FieldCtime, field.TypeTime, value)
		_node.Ctime = value
	}
	if value, ok := rac.mutation.Mtime(); ok {
		_spec.SetField(replyarea.FieldMtime, field.TypeTime, value)
		_node.Mtime = value
	}
	if value, ok := rac.mutation.Attrs(); ok {
		_spec.SetField(replyarea.FieldAttrs, field.TypeInt32, value)
		_node.Attrs = value
	}
	if value, ok := rac.mutation.Meta(); ok {
		_spec.SetField(replyarea.FieldMeta, field.TypeString, value)
		_node.Meta = value
	}
	return _node, _spec
}

// ReplyAreaCreateBulk is the builder for creating many ReplyArea entities in bulk.
type ReplyAreaCreateBulk struct {
	config
	builders []*ReplyAreaCreate
}

// Save creates the ReplyArea entities in the database.
func (racb *ReplyAreaCreateBulk) Save(ctx context.Context) ([]*ReplyArea, error) {
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*ReplyArea, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReplyAreaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *ReplyAreaCreateBulk) SaveX(ctx context.Context) []*ReplyArea {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *ReplyAreaCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *ReplyAreaCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}
