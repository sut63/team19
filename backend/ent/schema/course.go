package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{

		field.Int("Course_year").Positive().Range(1000,9000),
		field.String("Course_name").Match(regexp.MustCompile("^([A-Za-z0-9])+$")),
		field.String("Teacher_id").Match(regexp.MustCompile("[T]\\d{7}$")),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("Department_id", Department.Type).
			Ref("department").Unique(),

		edge.From("Degree_id", Degree.Type).
			Ref("degree").Unique(),

		edge.From("Subject_id", Subject.Type).
			Ref("subject").Unique(),
	}
}
