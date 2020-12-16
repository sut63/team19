package schema

import "github.com/facebookincubator/ent"

// Cause holds the schema definition for the Cause entity.
type Cause struct {
	ent.Schema
}

// Fields of the Cause.
func (Cause) Fields() []ent.Field {
	return nil
}

// Edges of the Cause.
func (Cause) Edges() []ent.Edge {
	return nil
}
