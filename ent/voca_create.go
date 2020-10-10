// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"lang.pkg/ent/voca"
)

// VocaCreate is the builder for creating a Voca entity.
type VocaCreate struct {
	config
	mutation *VocaMutation
	hooks    []Hook
}

// SetKey sets the key field.
func (vc *VocaCreate) SetKey(s string) *VocaCreate {
	vc.mutation.SetKey(s)
	return vc
}

// SetValue sets the value field.
func (vc *VocaCreate) SetValue(s string) *VocaCreate {
	vc.mutation.SetValue(s)
	return vc
}

// SetExample sets the example field.
func (vc *VocaCreate) SetExample(s string) *VocaCreate {
	vc.mutation.SetExample(s)
	return vc
}

// SetNillableExample sets the example field if the given value is not nil.
func (vc *VocaCreate) SetNillableExample(s *string) *VocaCreate {
	if s != nil {
		vc.SetExample(*s)
	}
	return vc
}

// SetCreatedAt sets the created_at field.
func (vc *VocaCreate) SetCreatedAt(t time.Time) *VocaCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (vc *VocaCreate) SetNillableCreatedAt(t *time.Time) *VocaCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetID sets the id field.
func (vc *VocaCreate) SetID(u uuid.UUID) *VocaCreate {
	vc.mutation.SetID(u)
	return vc
}

// Mutation returns the VocaMutation object of the builder.
func (vc *VocaCreate) Mutation() *VocaMutation {
	return vc.mutation
}

// Save creates the Voca in the database.
func (vc *VocaCreate) Save(ctx context.Context) (*Voca, error) {
	var (
		err  error
		node *Voca
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VocaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VocaCreate) SaveX(ctx context.Context) *Voca {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (vc *VocaCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := voca.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.ID(); !ok {
		v := voca.DefaultID()
		vc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VocaCreate) check() error {
	if _, ok := vc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New("ent: missing required field \"key\"")}
	}
	if _, ok := vc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (vc *VocaCreate) sqlSave(ctx context.Context) (*Voca, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (vc *VocaCreate) createSpec() (*Voca, *sqlgraph.CreateSpec) {
	var (
		_node = &Voca{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: voca.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voca.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: voca.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := vc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: voca.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := vc.mutation.Example(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: voca.FieldExample,
		})
		_node.Example = &value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: voca.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// VocaCreateBulk is the builder for creating a bulk of Voca entities.
type VocaCreateBulk struct {
	config
	builders []*VocaCreate
}

// Save creates the Voca entities in the database.
func (vcb *VocaCreateBulk) Save(ctx context.Context) ([]*Voca, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Voca, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VocaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (vcb *VocaCreateBulk) SaveX(ctx context.Context) []*Voca {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
