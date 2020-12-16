package schema

import "github.com/facebookincubator/ent"

// Krit holds the schema definition for the Krit entity.
type Krit struct {
	ent.Schema
}

// Fields of the Krit.
func (Krit) Fields() []ent.Field {
	return nil
}

// Edges of the Krit.
func (Krit) Edges() []ent.Edge {
	return nil
}
