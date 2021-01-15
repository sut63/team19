package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Classtime holds the schema definition for the Classtime entity.
type Classtime struct {
	ent.Schema
}

// Fields of the Classtime.
func (Classtime) Fields() []ent.Field {
	return []ent.Field{
		field.String("TIME").Unique(),
	}
}

// Edges of the Classtime.
func (Classtime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courseclasses",Courseclass.Type).StorageKey(edge.Column("classtime_id")),
	}
}
