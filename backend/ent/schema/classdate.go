package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Classdate holds the schema definition for the Classdate entity.
type Classdate struct {
	ent.Schema
}

// Fields of the Classdate.
func (Classdate) Fields() []ent.Field {
	return []ent.Field{
		field.String("DAY").Unique(),
	}
}

// Edges of the Classdate.
func (Classdate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courseclasses", Courseclass.Type).StorageKey(edge.Column("classdate_id")),
	}

}
