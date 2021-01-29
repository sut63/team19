package schema

import (
	"errors"
	"regexp"
	"strings"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// InstructorInfo holds the schema definition for the InstructorInfo entity.
type InstructorInfo struct {
	ent.Schema
}

// Fields of the InstructorInfo.
func (InstructorInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("Your name must begin with Uppercase")
				}
				return nil
			}).Unique().NotEmpty(),
		field.String("PHONENUMBER").Unique().
			//Match(regexp.MustCompile("[0]+[9]||[0]+[8]||[0]+[6]")).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[0]+[9]\\d{8}$|[0]+[8]\\d{8}$|[0]+[6]\\d{8}$", s)
				if !match {
					return errors.New("Your Phone-number must begin with 09 08 or 06 and limit 10 digits")
				}
				return nil
			}).NotEmpty(),
		field.String("EMAIL").Unique().NotEmpty(),
		field.String("PASSWORD").NotEmpty(),
	}
}

// Edges of the InstructorInfo.
func (InstructorInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("title", Title.Type).Ref("instructorinfos").Unique(),
		edge.From("instructorroom", InstructorRoom.Type).Ref("instructorinfos").Unique(),
		edge.From("department", Department.Type).Ref("instructorinfos").Unique(),
		edge.To("courseclasses", Courseclass.Type).StorageKey(edge.Column("InstructorInfo_id")),
	}
}
