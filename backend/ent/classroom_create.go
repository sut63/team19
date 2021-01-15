// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team19/app/ent/classroom"
	"github.com/team19/app/ent/courseclass"
)

// ClassroomCreate is the builder for creating a Classroom entity.
type ClassroomCreate struct {
	config
	mutation *ClassroomMutation
	hooks    []Hook
}

// SetROOM sets the ROOM field.
func (cc *ClassroomCreate) SetROOM(s string) *ClassroomCreate {
	cc.mutation.SetROOM(s)
	return cc
}

// AddCourseclassIDs adds the courseclasses edge to Courseclass by ids.
func (cc *ClassroomCreate) AddCourseclassIDs(ids ...int) *ClassroomCreate {
	cc.mutation.AddCourseclassIDs(ids...)
	return cc
}

// AddCourseclasses adds the courseclasses edges to Courseclass.
func (cc *ClassroomCreate) AddCourseclasses(c ...*Courseclass) *ClassroomCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCourseclassIDs(ids...)
}

// Mutation returns the ClassroomMutation object of the builder.
func (cc *ClassroomCreate) Mutation() *ClassroomMutation {
	return cc.mutation
}

// Save creates the Classroom in the database.
func (cc *ClassroomCreate) Save(ctx context.Context) (*Classroom, error) {
	if _, ok := cc.mutation.ROOM(); !ok {
		return nil, &ValidationError{Name: "ROOM", err: errors.New("ent: missing required field \"ROOM\"")}
	}
	var (
		err  error
		node *Classroom
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassroomCreate) SaveX(ctx context.Context) *Classroom {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *ClassroomCreate) sqlSave(ctx context.Context) (*Classroom, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *ClassroomCreate) createSpec() (*Classroom, *sqlgraph.CreateSpec) {
	var (
		c     = &Classroom{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: classroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: classroom.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.ROOM(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: classroom.FieldROOM,
		})
		c.ROOM = value
	}
	if nodes := cc.mutation.CourseclassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classroom.CourseclassesTable,
			Columns: []string{classroom.CourseclassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: courseclass.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}