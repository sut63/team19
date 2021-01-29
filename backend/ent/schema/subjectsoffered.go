package schema

import (
	"regexp"
	"errors"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// SubjectsOffered holds the schema definition for the SubjectsOffered entity.
type SubjectsOffered struct {
	ent.Schema
}

// Fields of the SubjectsOffered.
func (SubjectsOffered) Fields() []ent.Field {
	return []ent.Field{

		field.String("AMOUNT").Validate(func(s string) error {
			match, _ := regexp.MatchString("[0123456789]",s)
			if !match {
				return errors.New("รูปแบบจำนวนที่รับไม่ถูกต้อง")
			}
			return nil
	}),
		field.String("STATUS").Validate(func(s string) error {
			match, _ := regexp.MatchString("Open",s)
			if !match {
				return errors.New("รูปแบบสถานะไม่ถูกต้อง")
			}
			return nil
	}),
	field.String("Remain").Validate(func(s string) error {
		match, _ := regexp.MatchString("[0123456789]",s)
		if !match {
			return errors.New("รูปแบบคงเหลือไม่ถูกต้อง")
		}
		return nil
}),
	}
}

// Edges of the SubjectsOffered.
func (SubjectsOffered) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Subject", Subject.Type).
			Ref("SubjectsOffered").Unique(),
		edge.From("Degree", Degree.Type).
			Ref("SubjectsOffered").Unique(),
		edge.From("Year", Year.Type).
			Ref("SubjectsOffered").Unique(),
		edge.From("Term", Term.Type).
			Ref("SubjectsOffered").Unique(),
	}
}
