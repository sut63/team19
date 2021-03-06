// Code generated by entc, DO NOT EDIT.

package course

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCourseYear holds the string denoting the course_year field in the database.
	FieldCourseYear = "course_year"
	// FieldCourseName holds the string denoting the course_name field in the database.
	FieldCourseName = "course_name"
	// FieldTeacherID holds the string denoting the teacher_id field in the database.
	FieldTeacherID = "teacher_id"

	// EdgeDepartmentID holds the string denoting the department_id edge name in mutations.
	EdgeDepartmentID = "Department_id"
	// EdgeDegreeID holds the string denoting the degree_id edge name in mutations.
	EdgeDegreeID = "Degree_id"
	// EdgeSubjectID holds the string denoting the subject_id edge name in mutations.
	EdgeSubjectID = "Subject_id"

	// Table holds the table name of the course in the database.
	Table = "courses"
	// DepartmentIDTable is the table the holds the Department_id relation/edge.
	DepartmentIDTable = "courses"
	// DepartmentIDInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentIDInverseTable = "departments"
	// DepartmentIDColumn is the table column denoting the Department_id relation/edge.
	DepartmentIDColumn = "department_id"
	// DegreeIDTable is the table the holds the Degree_id relation/edge.
	DegreeIDTable = "courses"
	// DegreeIDInverseTable is the table name for the Degree entity.
	// It exists in this package in order to avoid circular dependency with the "degree" package.
	DegreeIDInverseTable = "degrees"
	// DegreeIDColumn is the table column denoting the Degree_id relation/edge.
	DegreeIDColumn = "Degree_id"
	// SubjectIDTable is the table the holds the Subject_id relation/edge.
	SubjectIDTable = "courses"
	// SubjectIDInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectIDInverseTable = "subjects"
	// SubjectIDColumn is the table column denoting the Subject_id relation/edge.
	SubjectIDColumn = "Subject_id"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldCourseYear,
	FieldCourseName,
	FieldTeacherID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Course type.
var ForeignKeys = []string{
	"Degree_id",
	"department_id",
	"Subject_id",
}

var (
	// CourseYearValidator is a validator for the "Course_year" field. It is called by the builders before save.
	CourseYearValidator func(int) error
	// CourseNameValidator is a validator for the "Course_name" field. It is called by the builders before save.
	CourseNameValidator func(string) error
	// TeacherIDValidator is a validator for the "Teacher_id" field. It is called by the builders before save.
	TeacherIDValidator func(string) error
)
