// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team19/app/ent/degree"
	"github.com/team19/app/ent/subject"
	"github.com/team19/app/ent/subjectsoffered"
	"github.com/team19/app/ent/term"
	"github.com/team19/app/ent/year"
)

// SubjectsOffered is the model entity for the SubjectsOffered schema.
type SubjectsOffered struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AMOUNT holds the value of the "AMOUNT" field.
	AMOUNT string `json:"AMOUNT,omitempty"`
	// STATUS holds the value of the "STATUS" field.
	STATUS string `json:"STATUS,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectsOfferedQuery when eager-loading is set.
	Edges      SubjectsOfferedEdges `json:"edges"`
	Degree_id  *int
	Subject_id *int
	term_id    *int
	year_id    *int
}

// SubjectsOfferedEdges holds the relations/edges for other nodes in the graph.
type SubjectsOfferedEdges struct {
	// Subject holds the value of the Subject edge.
	Subject *Subject
	// Degree holds the value of the Degree edge.
	Degree *Degree
	// Year holds the value of the Year edge.
	Year *Year
	// Term holds the value of the Term edge.
	Term *Term
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectsOfferedEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[0] {
		if e.Subject == nil {
			// The edge Subject was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "Subject"}
}

// DegreeOrErr returns the Degree value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectsOfferedEdges) DegreeOrErr() (*Degree, error) {
	if e.loadedTypes[1] {
		if e.Degree == nil {
			// The edge Degree was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: degree.Label}
		}
		return e.Degree, nil
	}
	return nil, &NotLoadedError{edge: "Degree"}
}

// YearOrErr returns the Year value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectsOfferedEdges) YearOrErr() (*Year, error) {
	if e.loadedTypes[2] {
		if e.Year == nil {
			// The edge Year was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: year.Label}
		}
		return e.Year, nil
	}
	return nil, &NotLoadedError{edge: "Year"}
}

// TermOrErr returns the Term value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectsOfferedEdges) TermOrErr() (*Term, error) {
	if e.loadedTypes[3] {
		if e.Term == nil {
			// The edge Term was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: term.Label}
		}
		return e.Term, nil
	}
	return nil, &NotLoadedError{edge: "Term"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubjectsOffered) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // AMOUNT
		&sql.NullString{}, // STATUS
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*SubjectsOffered) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // Degree_id
		&sql.NullInt64{}, // Subject_id
		&sql.NullInt64{}, // term_id
		&sql.NullInt64{}, // year_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubjectsOffered fields.
func (so *SubjectsOffered) assignValues(values ...interface{}) error {
	if m, n := len(values), len(subjectsoffered.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	so.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field AMOUNT", values[0])
	} else if value.Valid {
		so.AMOUNT = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field STATUS", values[1])
	} else if value.Valid {
		so.STATUS = value.String
	}
	values = values[2:]
	if len(values) == len(subjectsoffered.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Degree_id", value)
		} else if value.Valid {
			so.Degree_id = new(int)
			*so.Degree_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Subject_id", value)
		} else if value.Valid {
			so.Subject_id = new(int)
			*so.Subject_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field term_id", value)
		} else if value.Valid {
			so.term_id = new(int)
			*so.term_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field year_id", value)
		} else if value.Valid {
			so.year_id = new(int)
			*so.year_id = int(value.Int64)
		}
	}
	return nil
}

// QuerySubject queries the Subject edge of the SubjectsOffered.
func (so *SubjectsOffered) QuerySubject() *SubjectQuery {
	return (&SubjectsOfferedClient{config: so.config}).QuerySubject(so)
}

// QueryDegree queries the Degree edge of the SubjectsOffered.
func (so *SubjectsOffered) QueryDegree() *DegreeQuery {
	return (&SubjectsOfferedClient{config: so.config}).QueryDegree(so)
}

// QueryYear queries the Year edge of the SubjectsOffered.
func (so *SubjectsOffered) QueryYear() *YearQuery {
	return (&SubjectsOfferedClient{config: so.config}).QueryYear(so)
}

// QueryTerm queries the Term edge of the SubjectsOffered.
func (so *SubjectsOffered) QueryTerm() *TermQuery {
	return (&SubjectsOfferedClient{config: so.config}).QueryTerm(so)
}

// Update returns a builder for updating this SubjectsOffered.
// Note that, you need to call SubjectsOffered.Unwrap() before calling this method, if this SubjectsOffered
// was returned from a transaction, and the transaction was committed or rolled back.
func (so *SubjectsOffered) Update() *SubjectsOfferedUpdateOne {
	return (&SubjectsOfferedClient{config: so.config}).UpdateOne(so)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (so *SubjectsOffered) Unwrap() *SubjectsOffered {
	tx, ok := so.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubjectsOffered is not a transactional entity")
	}
	so.config.driver = tx.drv
	return so
}

// String implements the fmt.Stringer.
func (so *SubjectsOffered) String() string {
	var builder strings.Builder
	builder.WriteString("SubjectsOffered(")
	builder.WriteString(fmt.Sprintf("id=%v", so.ID))
	builder.WriteString(", AMOUNT=")
	builder.WriteString(so.AMOUNT)
	builder.WriteString(", STATUS=")
	builder.WriteString(so.STATUS)
	builder.WriteByte(')')
	return builder.String()
}

// SubjectsOffereds is a parsable slice of SubjectsOffered.
type SubjectsOffereds []*SubjectsOffered

func (so SubjectsOffereds) config(cfg config) {
	for _i := range so {
		so[_i].config = cfg
	}
}