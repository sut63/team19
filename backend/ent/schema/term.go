package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Year holds the schema definition for the Year entity.
type Term struct {
	ent.Schema
}

// Fields of the Year.
func (Term) Fields() []ent.Field {
	return []ent.Field{

		field.Int("TERM").Positive(),
	}
}

// Edges of the Year.
func (Term) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("SubjectsOffered", SubjectsOffered.Type).StorageKey(edge.Column("term_id")),
	}
}
