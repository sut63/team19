// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team19/app/ent/course"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
	"github.com/team19/app/ent/title"
)

// InstructorInfoCreate is the builder for creating a InstructorInfo entity.
type InstructorInfoCreate struct {
	config
	mutation *InstructorInfoMutation
	hooks    []Hook
}

// SetNAME sets the NAME field.
func (iic *InstructorInfoCreate) SetNAME(s string) *InstructorInfoCreate {
	iic.mutation.SetNAME(s)
	return iic
}

// SetPHONENUMBER sets the PHONENUMBER field.
func (iic *InstructorInfoCreate) SetPHONENUMBER(i int) *InstructorInfoCreate {
	iic.mutation.SetPHONENUMBER(i)
	return iic
}

// SetEMAIL sets the EMAIL field.
func (iic *InstructorInfoCreate) SetEMAIL(s string) *InstructorInfoCreate {
	iic.mutation.SetEMAIL(s)
	return iic
}

// SetPASSWORD sets the PASSWORD field.
func (iic *InstructorInfoCreate) SetPASSWORD(s string) *InstructorInfoCreate {
	iic.mutation.SetPASSWORD(s)
	return iic
}

// SetTitleID sets the title edge to Title by id.
func (iic *InstructorInfoCreate) SetTitleID(id int) *InstructorInfoCreate {
	iic.mutation.SetTitleID(id)
	return iic
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (iic *InstructorInfoCreate) SetNillableTitleID(id *int) *InstructorInfoCreate {
	if id != nil {
		iic = iic.SetTitleID(*id)
	}
	return iic
}

// SetTitle sets the title edge to Title.
func (iic *InstructorInfoCreate) SetTitle(t *Title) *InstructorInfoCreate {
	return iic.SetTitleID(t.ID)
}

// SetInstructorroomID sets the instructorroom edge to InstructorRoom by id.
func (iic *InstructorInfoCreate) SetInstructorroomID(id int) *InstructorInfoCreate {
	iic.mutation.SetInstructorroomID(id)
	return iic
}

// SetNillableInstructorroomID sets the instructorroom edge to InstructorRoom by id if the given value is not nil.
func (iic *InstructorInfoCreate) SetNillableInstructorroomID(id *int) *InstructorInfoCreate {
	if id != nil {
		iic = iic.SetInstructorroomID(*id)
	}
	return iic
}

// SetInstructorroom sets the instructorroom edge to InstructorRoom.
func (iic *InstructorInfoCreate) SetInstructorroom(i *InstructorRoom) *InstructorInfoCreate {
	return iic.SetInstructorroomID(i.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (iic *InstructorInfoCreate) SetDepartmentID(id int) *InstructorInfoCreate {
	iic.mutation.SetDepartmentID(id)
	return iic
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (iic *InstructorInfoCreate) SetNillableDepartmentID(id *int) *InstructorInfoCreate {
	if id != nil {
		iic = iic.SetDepartmentID(*id)
	}
	return iic
}

// SetDepartment sets the department edge to Department.
func (iic *InstructorInfoCreate) SetDepartment(d *Department) *InstructorInfoCreate {
	return iic.SetDepartmentID(d.ID)
}

// AddInstructorIDs adds the instructor edge to Course by ids.
func (iic *InstructorInfoCreate) AddInstructorIDs(ids ...int) *InstructorInfoCreate {
	iic.mutation.AddInstructorIDs(ids...)
	return iic
}

// AddInstructor adds the instructor edges to Course.
func (iic *InstructorInfoCreate) AddInstructor(c ...*Course) *InstructorInfoCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iic.AddInstructorIDs(ids...)
}

// Mutation returns the InstructorInfoMutation object of the builder.
func (iic *InstructorInfoCreate) Mutation() *InstructorInfoMutation {
	return iic.mutation
}

// Save creates the InstructorInfo in the database.
func (iic *InstructorInfoCreate) Save(ctx context.Context) (*InstructorInfo, error) {
	if _, ok := iic.mutation.NAME(); !ok {
		return nil, &ValidationError{Name: "NAME", err: errors.New("ent: missing required field \"NAME\"")}
	}
	if _, ok := iic.mutation.PHONENUMBER(); !ok {
		return nil, &ValidationError{Name: "PHONENUMBER", err: errors.New("ent: missing required field \"PHONENUMBER\"")}
	}
	if _, ok := iic.mutation.EMAIL(); !ok {
		return nil, &ValidationError{Name: "EMAIL", err: errors.New("ent: missing required field \"EMAIL\"")}
	}
	if _, ok := iic.mutation.PASSWORD(); !ok {
		return nil, &ValidationError{Name: "PASSWORD", err: errors.New("ent: missing required field \"PASSWORD\"")}
	}
	var (
		err  error
		node *InstructorInfo
	)
	if len(iic.hooks) == 0 {
		node, err = iic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructorInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iic.mutation = mutation
			node, err = iic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iic.hooks) - 1; i >= 0; i-- {
			mut = iic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iic *InstructorInfoCreate) SaveX(ctx context.Context) *InstructorInfo {
	v, err := iic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iic *InstructorInfoCreate) sqlSave(ctx context.Context) (*InstructorInfo, error) {
	ii, _spec := iic.createSpec()
	if err := sqlgraph.CreateNode(ctx, iic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ii.ID = int(id)
	return ii, nil
}

func (iic *InstructorInfoCreate) createSpec() (*InstructorInfo, *sqlgraph.CreateSpec) {
	var (
		ii    = &InstructorInfo{config: iic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instructorinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instructorinfo.FieldID,
			},
		}
	)
	if value, ok := iic.mutation.NAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructorinfo.FieldNAME,
		})
		ii.NAME = value
	}
	if value, ok := iic.mutation.PHONENUMBER(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instructorinfo.FieldPHONENUMBER,
		})
		ii.PHONENUMBER = value
	}
	if value, ok := iic.mutation.EMAIL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructorinfo.FieldEMAIL,
		})
		ii.EMAIL = value
	}
	if value, ok := iic.mutation.PASSWORD(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructorinfo.FieldPASSWORD,
		})
		ii.PASSWORD = value
	}
	if nodes := iic.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instructorinfo.TitleTable,
			Columns: []string{instructorinfo.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iic.mutation.InstructorroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instructorinfo.InstructorroomTable,
			Columns: []string{instructorinfo.InstructorroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instructorroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iic.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instructorinfo.DepartmentTable,
			Columns: []string{instructorinfo.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iic.mutation.InstructorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructorinfo.InstructorTable,
			Columns: []string{instructorinfo.InstructorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: course.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ii, _spec
}
