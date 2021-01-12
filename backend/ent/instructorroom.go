// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team19/app/ent/instructorroom"
)

// InstructorRoom is the model entity for the InstructorRoom schema.
type InstructorRoom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ROOM holds the value of the "ROOM" field.
	ROOM string `json:"ROOM,omitempty"`
	// BUILDING holds the value of the "BUILDING" field.
	BUILDING string `json:"BUILDING,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstructorRoomQuery when eager-loading is set.
	Edges InstructorRoomEdges `json:"edges"`
}

// InstructorRoomEdges holds the relations/edges for other nodes in the graph.
type InstructorRoomEdges struct {
	// Instructorinfos holds the value of the instructorinfos edge.
	Instructorinfos []*InstructorInfo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstructorinfosOrErr returns the Instructorinfos value or an error if the edge
// was not loaded in eager-loading.
func (e InstructorRoomEdges) InstructorinfosOrErr() ([]*InstructorInfo, error) {
	if e.loadedTypes[0] {
		return e.Instructorinfos, nil
	}
	return nil, &NotLoadedError{edge: "instructorinfos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InstructorRoom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // ROOM
		&sql.NullString{}, // BUILDING
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InstructorRoom fields.
func (ir *InstructorRoom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(instructorroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ir.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ROOM", values[0])
	} else if value.Valid {
		ir.ROOM = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field BUILDING", values[1])
	} else if value.Valid {
		ir.BUILDING = value.String
	}
	return nil
}

// QueryInstructorinfos queries the instructorinfos edge of the InstructorRoom.
func (ir *InstructorRoom) QueryInstructorinfos() *InstructorInfoQuery {
	return (&InstructorRoomClient{config: ir.config}).QueryInstructorinfos(ir)
}

// Update returns a builder for updating this InstructorRoom.
// Note that, you need to call InstructorRoom.Unwrap() before calling this method, if this InstructorRoom
// was returned from a transaction, and the transaction was committed or rolled back.
func (ir *InstructorRoom) Update() *InstructorRoomUpdateOne {
	return (&InstructorRoomClient{config: ir.config}).UpdateOne(ir)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ir *InstructorRoom) Unwrap() *InstructorRoom {
	tx, ok := ir.config.driver.(*txDriver)
	if !ok {
		panic("ent: InstructorRoom is not a transactional entity")
	}
	ir.config.driver = tx.drv
	return ir
}

// String implements the fmt.Stringer.
func (ir *InstructorRoom) String() string {
	var builder strings.Builder
	builder.WriteString("InstructorRoom(")
	builder.WriteString(fmt.Sprintf("id=%v", ir.ID))
	builder.WriteString(", ROOM=")
	builder.WriteString(ir.ROOM)
	builder.WriteString(", BUILDING=")
	builder.WriteString(ir.BUILDING)
	builder.WriteByte(')')
	return builder.String()
}

// InstructorRooms is a parsable slice of InstructorRoom.
type InstructorRooms []*InstructorRoom

func (ir InstructorRooms) config(cfg config) {
	for _i := range ir {
		ir[_i].config = cfg
	}
}
