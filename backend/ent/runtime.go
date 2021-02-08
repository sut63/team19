// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team19/app/ent/course"
	"github.com/team19/app/ent/courseclass"
	"github.com/team19/app/ent/instructorinfo"
	"github.com/team19/app/ent/schema"
	"github.com/team19/app/ent/subjectsoffered"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescCourseYear is the schema descriptor for Course_year field.
	courseDescCourseYear := courseFields[0].Descriptor()
	// course.CourseYearValidator is a validator for the "Course_year" field. It is called by the builders before save.
	course.CourseYearValidator = courseDescCourseYear.Validators[0].(func(string) error)
	// courseDescCourseName is the schema descriptor for Course_name field.
	courseDescCourseName := courseFields[1].Descriptor()
	// course.CourseNameValidator is a validator for the "Course_name" field. It is called by the builders before save.
	course.CourseNameValidator = courseDescCourseName.Validators[0].(func(string) error)
	// courseDescTeacherID is the schema descriptor for Teacher_id field.
	courseDescTeacherID := courseFields[2].Descriptor()
	// course.TeacherIDValidator is a validator for the "Teacher_id" field. It is called by the builders before save.
	course.TeacherIDValidator = courseDescTeacherID.Validators[0].(func(string) error)
	courseclassFields := schema.Courseclass{}.Fields()
	_ = courseclassFields
	// courseclassDescTablecode is the schema descriptor for tablecode field.
	courseclassDescTablecode := courseclassFields[0].Descriptor()
	// courseclass.TablecodeValidator is a validator for the "tablecode" field. It is called by the builders before save.
	courseclass.TablecodeValidator = func() func(string) error {
		validators := courseclassDescTablecode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(tablecode string) error {
			for _, fn := range fns {
				if err := fn(tablecode); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	instructorinfoFields := schema.InstructorInfo{}.Fields()
	_ = instructorinfoFields
	// instructorinfoDescNAME is the schema descriptor for NAME field.
	instructorinfoDescNAME := instructorinfoFields[0].Descriptor()
	// instructorinfo.NAMEValidator is a validator for the "NAME" field. It is called by the builders before save.
	instructorinfo.NAMEValidator = func() func(string) error {
		validators := instructorinfoDescNAME.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_NAME string) error {
			for _, fn := range fns {
				if err := fn(_NAME); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// instructorinfoDescPHONENUMBER is the schema descriptor for PHONENUMBER field.
	instructorinfoDescPHONENUMBER := instructorinfoFields[1].Descriptor()
	// instructorinfo.PHONENUMBERValidator is a validator for the "PHONENUMBER" field. It is called by the builders before save.
	instructorinfo.PHONENUMBERValidator = func() func(string) error {
		validators := instructorinfoDescPHONENUMBER.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_PHONENUMBER string) error {
			for _, fn := range fns {
				if err := fn(_PHONENUMBER); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// instructorinfoDescEMAIL is the schema descriptor for EMAIL field.
	instructorinfoDescEMAIL := instructorinfoFields[2].Descriptor()
	// instructorinfo.EMAILValidator is a validator for the "EMAIL" field. It is called by the builders before save.
	instructorinfo.EMAILValidator = instructorinfoDescEMAIL.Validators[0].(func(string) error)
	// instructorinfoDescPASSWORD is the schema descriptor for PASSWORD field.
	instructorinfoDescPASSWORD := instructorinfoFields[3].Descriptor()
	// instructorinfo.PASSWORDValidator is a validator for the "PASSWORD" field. It is called by the builders before save.
	instructorinfo.PASSWORDValidator = instructorinfoDescPASSWORD.Validators[0].(func(string) error)
	subjectsofferedFields := schema.SubjectsOffered{}.Fields()
	_ = subjectsofferedFields
	// subjectsofferedDescAMOUNT is the schema descriptor for AMOUNT field.
	subjectsofferedDescAMOUNT := subjectsofferedFields[0].Descriptor()
	// subjectsoffered.AMOUNTValidator is a validator for the "AMOUNT" field. It is called by the builders before save.
	subjectsoffered.AMOUNTValidator = subjectsofferedDescAMOUNT.Validators[0].(func(string) error)
	// subjectsofferedDescSTATUS is the schema descriptor for STATUS field.
	subjectsofferedDescSTATUS := subjectsofferedFields[1].Descriptor()
	// subjectsoffered.STATUSValidator is a validator for the "STATUS" field. It is called by the builders before save.
	subjectsoffered.STATUSValidator = subjectsofferedDescSTATUS.Validators[0].(func(string) error)
	// subjectsofferedDescRemain is the schema descriptor for Remain field.
	subjectsofferedDescRemain := subjectsofferedFields[2].Descriptor()
	// subjectsoffered.RemainValidator is a validator for the "Remain" field. It is called by the builders before save.
	subjectsoffered.RemainValidator = subjectsofferedDescRemain.Validators[0].(func(string) error)
}
