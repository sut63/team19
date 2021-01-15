// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team19/app/ent/term"
)

// Term is the model entity for the Term schema.
type Term struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TERM holds the value of the "TERM" field.
	TERM int `json:"TERM,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TermQuery when eager-loading is set.
	Edges TermEdges `json:"edges"`
}

// TermEdges holds the relations/edges for other nodes in the graph.
type TermEdges struct {
	// SubjectsOffered holds the value of the SubjectsOffered edge.
	SubjectsOffered []*SubjectsOffered
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubjectsOfferedOrErr returns the SubjectsOffered value or an error if the edge
// was not loaded in eager-loading.
func (e TermEdges) SubjectsOfferedOrErr() ([]*SubjectsOffered, error) {
	if e.loadedTypes[0] {
		return e.SubjectsOffered, nil
	}
	return nil, &NotLoadedError{edge: "SubjectsOffered"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Term) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // TERM
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Term fields.
func (t *Term) assignValues(values ...interface{}) error {
	if m, n := len(values), len(term.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field TERM", values[0])
	} else if value.Valid {
		t.TERM = int(value.Int64)
	}
	return nil
}

// QuerySubjectsOffered queries the SubjectsOffered edge of the Term.
func (t *Term) QuerySubjectsOffered() *SubjectsOfferedQuery {
	return (&TermClient{config: t.config}).QuerySubjectsOffered(t)
}

// Update returns a builder for updating this Term.
// Note that, you need to call Term.Unwrap() before calling this method, if this Term
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Term) Update() *TermUpdateOne {
	return (&TermClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Term) Unwrap() *Term {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Term is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Term) String() string {
	var builder strings.Builder
	builder.WriteString("Term(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", TERM=")
	builder.WriteString(fmt.Sprintf("%v", t.TERM))
	builder.WriteByte(')')
	return builder.String()
}

// Terms is a parsable slice of Term.
type Terms []*Term

func (t Terms) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}