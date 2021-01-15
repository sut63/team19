// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team19/app/ent/subject"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubjectName holds the value of the "Subject_name" field.
	SubjectName string `json:"Subject_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges SubjectEdges `json:"edges"`
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// Subject holds the value of the subject edge.
	Subject []*Course
	// SubjectsOffered holds the value of the SubjectsOffered edge.
	SubjectsOffered []*SubjectsOffered
	// Courseclasses holds the value of the courseclasses edge.
	Courseclasses []*Courseclass
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SubjectOrErr() ([]*Course, error) {
	if e.loadedTypes[0] {
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// SubjectsOfferedOrErr returns the SubjectsOffered value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SubjectsOfferedOrErr() ([]*SubjectsOffered, error) {
	if e.loadedTypes[1] {
		return e.SubjectsOffered, nil
	}
	return nil, &NotLoadedError{edge: "SubjectsOffered"}
}

// CourseclassesOrErr returns the Courseclasses value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) CourseclassesOrErr() ([]*Courseclass, error) {
	if e.loadedTypes[2] {
		return e.Courseclasses, nil
	}
	return nil, &NotLoadedError{edge: "courseclasses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Subject_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(values ...interface{}) error {
	if m, n := len(values), len(subject.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Subject_name", values[0])
	} else if value.Valid {
		s.SubjectName = value.String
	}
	return nil
}

// QuerySubject queries the subject edge of the Subject.
func (s *Subject) QuerySubject() *CourseQuery {
	return (&SubjectClient{config: s.config}).QuerySubject(s)
}

// QuerySubjectsOffered queries the SubjectsOffered edge of the Subject.
func (s *Subject) QuerySubjectsOffered() *SubjectsOfferedQuery {
	return (&SubjectClient{config: s.config}).QuerySubjectsOffered(s)
}

// QueryCourseclasses queries the courseclasses edge of the Subject.
func (s *Subject) QueryCourseclasses() *CourseclassQuery {
	return (&SubjectClient{config: s.config}).QueryCourseclasses(s)
}

// Update returns a builder for updating this Subject.
// Note that, you need to call Subject.Unwrap() before calling this method, if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return (&SubjectClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Subject_name=")
	builder.WriteString(s.SubjectName)
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

func (s Subjects) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}