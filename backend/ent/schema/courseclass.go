package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Courseclass holds the schema definition for the Courseclass entity.
type Courseclass struct {
	ent.Schema
}

// Fields of the Courseclass.
func (Courseclass) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Courseclass.
func (Courseclass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classtime", Classtime.Type).Ref("courseclasses").Unique(),
		edge.From("classdate", Classdate.Type).Ref("courseclasses").Unique(),
		edge.From("classroom", Classroom.Type).Ref("courseclasses").Unique(),
		edge.From("instructorInfo", InstructorInfo.Type).Ref("courseclasses").Unique(),
		edge.From("subject", Subject.Type).Ref("courseclasses").Unique(),
	}
}
