package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Year holds the schema definition for the Year entity.
type Year struct {
	ent.Schema
}

// Fields of the Year.
func (Year) Fields() []ent.Field {
	return []ent.Field{
		field.Int("YEAR").Unique(),
	}
}

// Edges of the Year.
func (Year) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("SubjectsOffered", SubjectsOffered.Type).StorageKey(edge.Column("year_id")),
	}
}
