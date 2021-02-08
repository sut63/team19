package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Courseclass holds the schema definition for the Courseclass entity.
type Courseclass struct {
	ent.Schema
}

// Fields of the Courseclass.
func (Courseclass) Fields() []ent.Field {
	return []ent.Field{
		field.String("tablecode").
		Validate(func(s string) error {
			match, _ := regexp.MatchString("[T]\\d{2}", s)
			if !match {
				return errors.New("Tablecode must begin with T and limit 2 digits")
			}
			return nil
		}).NotEmpty(),
	}
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
