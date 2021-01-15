package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Classroom holds the schema definition for the Classroom entity.
type Classroom struct {
	ent.Schema
}

// Fields of the Classroom.
func (Classroom) Fields() []ent.Field {
	return []ent.Field{
		field.String("ROOM").Unique(),
	}
}

// Edges of the Classroom.
func (Classroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("courseclasses", Courseclass.Type).StorageKey(edge.Column("classroom_id")),
	}

}
