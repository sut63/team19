// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team19/app/ent/department"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/instructorroom"
	"github.com/team19/app/ent/title"
)

// InstructorInfo is the model entity for the InstructorInfo schema.
type InstructorInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NAME holds the value of the "NAME" field.
	NAME string `json:"NAME,omitempty"`
	// PHONENUMBER holds the value of the "PHONENUMBER" field.
	PHONENUMBER string `json:"PHONENUMBER,omitempty"`
	// EMAIL holds the value of the "EMAIL" field.
	EMAIL string `json:"EMAIL,omitempty"`
	// PASSWORD holds the value of the "PASSWORD" field.
	PASSWORD string `json:"PASSWORD,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstructorInfoQuery when eager-loading is set.
	Edges             InstructorInfoEdges `json:"edges"`
	department_id     *int
	instructorroom_id *int
	title_id          *int
}

// InstructorInfoEdges holds the relations/edges for other nodes in the graph.
type InstructorInfoEdges struct {
	// Title holds the value of the title edge.
	Title *Title
	// Instructorroom holds the value of the instructorroom edge.
	Instructorroom *InstructorRoom
	// Department holds the value of the department edge.
	Department *Department
	// Courseclasses holds the value of the courseclasses edge.
	Courseclasses []*Courseclass
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstructorInfoEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[0] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// InstructorroomOrErr returns the Instructorroom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstructorInfoEdges) InstructorroomOrErr() (*InstructorRoom, error) {
	if e.loadedTypes[1] {
		if e.Instructorroom == nil {
			// The edge instructorroom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instructorroom.Label}
		}
		return e.Instructorroom, nil
	}
	return nil, &NotLoadedError{edge: "instructorroom"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstructorInfoEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[2] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// CourseclassesOrErr returns the Courseclasses value or an error if the edge
// was not loaded in eager-loading.
func (e InstructorInfoEdges) CourseclassesOrErr() ([]*Courseclass, error) {
	if e.loadedTypes[3] {
		return e.Courseclasses, nil
	}
	return nil, &NotLoadedError{edge: "courseclasses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InstructorInfo) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NAME
		&sql.NullString{}, // PHONENUMBER
		&sql.NullString{}, // EMAIL
		&sql.NullString{}, // PASSWORD
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*InstructorInfo) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // instructorroom_id
		&sql.NullInt64{}, // title_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InstructorInfo fields.
func (ii *InstructorInfo) assignValues(values ...interface{}) error {
	if m, n := len(values), len(instructorinfo.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ii.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field NAME", values[0])
	} else if value.Valid {
		ii.NAME = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PHONENUMBER", values[1])
	} else if value.Valid {
		ii.PHONENUMBER = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMAIL", values[2])
	} else if value.Valid {
		ii.EMAIL = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PASSWORD", values[3])
	} else if value.Valid {
		ii.PASSWORD = value.String
	}
	values = values[4:]
	if len(values) == len(instructorinfo.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			ii.department_id = new(int)
			*ii.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field instructorroom_id", value)
		} else if value.Valid {
			ii.instructorroom_id = new(int)
			*ii.instructorroom_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title_id", value)
		} else if value.Valid {
			ii.title_id = new(int)
			*ii.title_id = int(value.Int64)
		}
	}
	return nil
}

// QueryTitle queries the title edge of the InstructorInfo.
func (ii *InstructorInfo) QueryTitle() *TitleQuery {
	return (&InstructorInfoClient{config: ii.config}).QueryTitle(ii)
}

// QueryInstructorroom queries the instructorroom edge of the InstructorInfo.
func (ii *InstructorInfo) QueryInstructorroom() *InstructorRoomQuery {
	return (&InstructorInfoClient{config: ii.config}).QueryInstructorroom(ii)
}

// QueryDepartment queries the department edge of the InstructorInfo.
func (ii *InstructorInfo) QueryDepartment() *DepartmentQuery {
	return (&InstructorInfoClient{config: ii.config}).QueryDepartment(ii)
}

// QueryCourseclasses queries the courseclasses edge of the InstructorInfo.
func (ii *InstructorInfo) QueryCourseclasses() *CourseclassQuery {
	return (&InstructorInfoClient{config: ii.config}).QueryCourseclasses(ii)
}

// Update returns a builder for updating this InstructorInfo.
// Note that, you need to call InstructorInfo.Unwrap() before calling this method, if this InstructorInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ii *InstructorInfo) Update() *InstructorInfoUpdateOne {
	return (&InstructorInfoClient{config: ii.config}).UpdateOne(ii)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ii *InstructorInfo) Unwrap() *InstructorInfo {
	tx, ok := ii.config.driver.(*txDriver)
	if !ok {
		panic("ent: InstructorInfo is not a transactional entity")
	}
	ii.config.driver = tx.drv
	return ii
}

// String implements the fmt.Stringer.
func (ii *InstructorInfo) String() string {
	var builder strings.Builder
	builder.WriteString("InstructorInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", ii.ID))
	builder.WriteString(", NAME=")
	builder.WriteString(ii.NAME)
	builder.WriteString(", PHONENUMBER=")
	builder.WriteString(ii.PHONENUMBER)
	builder.WriteString(", EMAIL=")
	builder.WriteString(ii.EMAIL)
	builder.WriteString(", PASSWORD=")
	builder.WriteString(ii.PASSWORD)
	builder.WriteByte(')')
	return builder.String()
}

// InstructorInfos is a parsable slice of InstructorInfo.
type InstructorInfos []*InstructorInfo

func (ii InstructorInfos) config(cfg config) {
	for _i := range ii {
		ii[_i].config = cfg
	}
}
