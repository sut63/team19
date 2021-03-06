// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
)

// InstructorRoomCreate is the builder for creating a InstructorRoom entity.
type InstructorRoomCreate struct {
	config
	mutation *InstructorRoomMutation
	hooks    []Hook
}

// SetROOM sets the ROOM field.
func (irc *InstructorRoomCreate) SetROOM(s string) *InstructorRoomCreate {
	irc.mutation.SetROOM(s)
	return irc
}

// SetBUILDING sets the BUILDING field.
func (irc *InstructorRoomCreate) SetBUILDING(s string) *InstructorRoomCreate {
	irc.mutation.SetBUILDING(s)
	return irc
}

// AddInstructorinfoIDs adds the instructorinfos edge to InstructorInfo by ids.
func (irc *InstructorRoomCreate) AddInstructorinfoIDs(ids ...int) *InstructorRoomCreate {
	irc.mutation.AddInstructorinfoIDs(ids...)
	return irc
}

// AddInstructorinfos adds the instructorinfos edges to InstructorInfo.
func (irc *InstructorRoomCreate) AddInstructorinfos(i ...*InstructorInfo) *InstructorRoomCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return irc.AddInstructorinfoIDs(ids...)
}

// Mutation returns the InstructorRoomMutation object of the builder.
func (irc *InstructorRoomCreate) Mutation() *InstructorRoomMutation {
	return irc.mutation
}

// Save creates the InstructorRoom in the database.
func (irc *InstructorRoomCreate) Save(ctx context.Context) (*InstructorRoom, error) {
	if _, ok := irc.mutation.ROOM(); !ok {
		return nil, &ValidationError{Name: "ROOM", err: errors.New("ent: missing required field \"ROOM\"")}
	}
	if _, ok := irc.mutation.BUILDING(); !ok {
		return nil, &ValidationError{Name: "BUILDING", err: errors.New("ent: missing required field \"BUILDING\"")}
	}
	var (
		err  error
		node *InstructorRoom
	)
	if len(irc.hooks) == 0 {
		node, err = irc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructorRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			irc.mutation = mutation
			node, err = irc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(irc.hooks) - 1; i >= 0; i-- {
			mut = irc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, irc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (irc *InstructorRoomCreate) SaveX(ctx context.Context) *InstructorRoom {
	v, err := irc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (irc *InstructorRoomCreate) sqlSave(ctx context.Context) (*InstructorRoom, error) {
	ir, _spec := irc.createSpec()
	if err := sqlgraph.CreateNode(ctx, irc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ir.ID = int(id)
	return ir, nil
}

func (irc *InstructorRoomCreate) createSpec() (*InstructorRoom, *sqlgraph.CreateSpec) {
	var (
		ir    = &InstructorRoom{config: irc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instructorroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instructorroom.FieldID,
			},
		}
	)
	if value, ok := irc.mutation.ROOM(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructorroom.FieldROOM,
		})
		ir.ROOM = value
	}
	if value, ok := irc.mutation.BUILDING(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructorroom.FieldBUILDING,
		})
		ir.BUILDING = value
	}
	if nodes := irc.mutation.InstructorinfosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructorroom.InstructorinfosTable,
			Columns: []string{instructorroom.InstructorinfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instructorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ir, _spec
}
