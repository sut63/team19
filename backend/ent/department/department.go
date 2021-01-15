// Code generated by entc, DO NOT EDIT.

package department

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDEPARTMENT holds the string denoting the department field in the database.
	FieldDEPARTMENT = "department"
	// FieldFACULTY holds the string denoting the faculty field in the database.
	FieldFACULTY = "faculty"

	// EdgeInstructorinfos holds the string denoting the instructorinfos edge name in mutations.
	EdgeInstructorinfos = "instructorinfos"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"

	// Table holds the table name of the department in the database.
	Table = "departments"
	// InstructorinfosTable is the table the holds the instructorinfos relation/edge.
	InstructorinfosTable = "instructor_infos"
	// InstructorinfosInverseTable is the table name for the InstructorInfo entity.
	// It exists in this package in order to avoid circular dependency with the "instructorinfo" package.
	InstructorinfosInverseTable = "instructor_infos"
	// InstructorinfosColumn is the table column denoting the instructorinfos relation/edge.
	InstructorinfosColumn = "department_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "courses"
	// DepartmentInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	DepartmentInverseTable = "courses"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldDEPARTMENT,
	FieldFACULTY,
}
