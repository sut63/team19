// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team19/app/ent/subjectsoffered"
	"github.com/team19/app/ent/term"
)

// TermCreate is the builder for creating a Term entity.
type TermCreate struct {
	config
	mutation *TermMutation
	hooks    []Hook
}

// SetTERM sets the TERM field.
func (tc *TermCreate) SetTERM(i int) *TermCreate {
	tc.mutation.SetTERM(i)
	return tc
}

// AddSubjectsOfferedIDs adds the SubjectsOffered edge to SubjectsOffered by ids.
func (tc *TermCreate) AddSubjectsOfferedIDs(ids ...int) *TermCreate {
	tc.mutation.AddSubjectsOfferedIDs(ids...)
	return tc
}

// AddSubjectsOffered adds the SubjectsOffered edges to SubjectsOffered.
func (tc *TermCreate) AddSubjectsOffered(s ...*SubjectsOffered) *TermCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddSubjectsOfferedIDs(ids...)
}

// Mutation returns the TermMutation object of the builder.
func (tc *TermCreate) Mutation() *TermMutation {
	return tc.mutation
}

// Save creates the Term in the database.
func (tc *TermCreate) Save(ctx context.Context) (*Term, error) {
	if _, ok := tc.mutation.TERM(); !ok {
		return nil, &ValidationError{Name: "TERM", err: errors.New("ent: missing required field \"TERM\"")}
	}
	var (
		err  error
		node *Term
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TermMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TermCreate) SaveX(ctx context.Context) *Term {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TermCreate) sqlSave(ctx context.Context) (*Term, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TermCreate) createSpec() (*Term, *sqlgraph.CreateSpec) {
	var (
		t     = &Term{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: term.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: term.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.TERM(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: term.FieldTERM,
		})
		t.TERM = value
	}
	if nodes := tc.mutation.SubjectsOfferedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   term.SubjectsOfferedTable,
			Columns: []string{term.SubjectsOfferedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subjectsoffered.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
