package schema

import (

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

		field.Int("AMOUNT").Min(1),
		field.Bool("STATUS").Default(false),
		field.Int("Remain").Min(0),
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
