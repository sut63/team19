// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team19/app/ent/course"
	"github.com/team19/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescCourseName is the schema descriptor for Course_name field.
	courseDescCourseName := courseFields[0].Descriptor()
	// course.CourseNameValidator is a validator for the "Course_name" field. It is called by the builders before save.
	course.CourseNameValidator = courseDescCourseName.Validators[0].(func(string) error)
}
