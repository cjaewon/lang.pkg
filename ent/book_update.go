// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"lang.pkg/ent/book"
	"lang.pkg/ent/predicate"
	"lang.pkg/ent/user"
	"lang.pkg/ent/voca"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks      []Hook
	mutation   *BookMutation
	predicates []predicate.Book
}

// Where adds a new predicate for the builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBookID sets the book_id field.
func (bu *BookUpdate) SetBookID(s string) *BookUpdate {
	bu.mutation.SetBookID(s)
	return bu
}

// SetNillableBookID sets the book_id field if the given value is not nil.
func (bu *BookUpdate) SetNillableBookID(s *string) *BookUpdate {
	if s != nil {
		bu.SetBookID(*s)
	}
	return bu
}

// ClearBookID clears the value of book_id.
func (bu *BookUpdate) ClearBookID() *BookUpdate {
	bu.mutation.ClearBookID()
	return bu
}

// SetName sets the name field.
func (bu *BookUpdate) SetName(s string) *BookUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetDescription sets the description field.
func (bu *BookUpdate) SetDescription(s string) *BookUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetPublic sets the public field.
func (bu *BookUpdate) SetPublic(b bool) *BookUpdate {
	bu.mutation.SetPublic(b)
	return bu
}

// SetCreatedAt sets the created_at field.
func (bu *BookUpdate) SetCreatedAt(t time.Time) *BookUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (bu *BookUpdate) SetNillableCreatedAt(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// AddVocaIDs adds the vocas edge to Voca by ids.
func (bu *BookUpdate) AddVocaIDs(ids ...uuid.UUID) *BookUpdate {
	bu.mutation.AddVocaIDs(ids...)
	return bu
}

// AddVocas adds the vocas edges to Voca.
func (bu *BookUpdate) AddVocas(v ...*Voca) *BookUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return bu.AddVocaIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (bu *BookUpdate) SetOwnerID(id uuid.UUID) *BookUpdate {
	bu.mutation.SetOwnerID(id)
	return bu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (bu *BookUpdate) SetNillableOwnerID(id *uuid.UUID) *BookUpdate {
	if id != nil {
		bu = bu.SetOwnerID(*id)
	}
	return bu
}

// SetOwner sets the owner edge to User.
func (bu *BookUpdate) SetOwner(u *User) *BookUpdate {
	return bu.SetOwnerID(u.ID)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearVocas clears all "vocas" edges to type Voca.
func (bu *BookUpdate) ClearVocas() *BookUpdate {
	bu.mutation.ClearVocas()
	return bu
}

// RemoveVocaIDs removes the vocas edge to Voca by ids.
func (bu *BookUpdate) RemoveVocaIDs(ids ...uuid.UUID) *BookUpdate {
	bu.mutation.RemoveVocaIDs(ids...)
	return bu
}

// RemoveVocas removes vocas edges to Voca.
func (bu *BookUpdate) RemoveVocas(v ...*Voca) *BookUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return bu.RemoveVocaIDs(ids...)
}

// ClearOwner clears the "owner" edge to type User.
func (bu *BookUpdate) ClearOwner() *BookUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BookID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldBookID,
		})
	}
	if bu.mutation.BookIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: book.FieldBookID,
		})
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldName,
		})
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldDescription,
		})
	}
	if value, ok := bu.mutation.Public(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: book.FieldPublic,
		})
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldCreatedAt,
		})
	}
	if bu.mutation.VocasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedVocasIDs(); len(nodes) > 0 && !bu.mutation.VocasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.VocasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.OwnerTable,
			Columns: []string{book.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.OwnerTable,
			Columns: []string{book.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// SetBookID sets the book_id field.
func (buo *BookUpdateOne) SetBookID(s string) *BookUpdateOne {
	buo.mutation.SetBookID(s)
	return buo
}

// SetNillableBookID sets the book_id field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableBookID(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetBookID(*s)
	}
	return buo
}

// ClearBookID clears the value of book_id.
func (buo *BookUpdateOne) ClearBookID() *BookUpdateOne {
	buo.mutation.ClearBookID()
	return buo
}

// SetName sets the name field.
func (buo *BookUpdateOne) SetName(s string) *BookUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetDescription sets the description field.
func (buo *BookUpdateOne) SetDescription(s string) *BookUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetPublic sets the public field.
func (buo *BookUpdateOne) SetPublic(b bool) *BookUpdateOne {
	buo.mutation.SetPublic(b)
	return buo
}

// SetCreatedAt sets the created_at field.
func (buo *BookUpdateOne) SetCreatedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCreatedAt(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// AddVocaIDs adds the vocas edge to Voca by ids.
func (buo *BookUpdateOne) AddVocaIDs(ids ...uuid.UUID) *BookUpdateOne {
	buo.mutation.AddVocaIDs(ids...)
	return buo
}

// AddVocas adds the vocas edges to Voca.
func (buo *BookUpdateOne) AddVocas(v ...*Voca) *BookUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return buo.AddVocaIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (buo *BookUpdateOne) SetOwnerID(id uuid.UUID) *BookUpdateOne {
	buo.mutation.SetOwnerID(id)
	return buo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (buo *BookUpdateOne) SetNillableOwnerID(id *uuid.UUID) *BookUpdateOne {
	if id != nil {
		buo = buo.SetOwnerID(*id)
	}
	return buo
}

// SetOwner sets the owner edge to User.
func (buo *BookUpdateOne) SetOwner(u *User) *BookUpdateOne {
	return buo.SetOwnerID(u.ID)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearVocas clears all "vocas" edges to type Voca.
func (buo *BookUpdateOne) ClearVocas() *BookUpdateOne {
	buo.mutation.ClearVocas()
	return buo
}

// RemoveVocaIDs removes the vocas edge to Voca by ids.
func (buo *BookUpdateOne) RemoveVocaIDs(ids ...uuid.UUID) *BookUpdateOne {
	buo.mutation.RemoveVocaIDs(ids...)
	return buo
}

// RemoveVocas removes vocas edges to Voca.
func (buo *BookUpdateOne) RemoveVocas(v ...*Voca) *BookUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return buo.RemoveVocaIDs(ids...)
}

// ClearOwner clears the "owner" edge to type User.
func (buo *BookUpdateOne) ClearOwner() *BookUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	var (
		err  error
		node *Book
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Book.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.BookID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldBookID,
		})
	}
	if buo.mutation.BookIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: book.FieldBookID,
		})
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldName,
		})
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldDescription,
		})
	}
	if value, ok := buo.mutation.Public(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: book.FieldPublic,
		})
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldCreatedAt,
		})
	}
	if buo.mutation.VocasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedVocasIDs(); len(nodes) > 0 && !buo.mutation.VocasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.VocasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   book.VocasTable,
			Columns: []string{book.VocasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voca.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.OwnerTable,
			Columns: []string{book.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.OwnerTable,
			Columns: []string{book.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
