// Code generated by entc, DO NOT EDIT.

package instructorinfo

const (
	// Label holds the string label denoting the instructorinfo type in the database.
	Label = "instructor_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNAME holds the string denoting the name field in the database.
	FieldNAME = "name"
	// FieldPHONENUMBER holds the string denoting the phonenumber field in the database.
	FieldPHONENUMBER = "phonenumber"
	// FieldEMAIL holds the string denoting the email field in the database.
	FieldEMAIL = "email"
	// FieldPASSWORD holds the string denoting the password field in the database.
	FieldPASSWORD = "password"

	// EdgeTitle holds the string denoting the title edge name in mutations.
	EdgeTitle = "title"
	// EdgeInstructorroom holds the string denoting the instructorroom edge name in mutations.
	EdgeInstructorroom = "instructorroom"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgeInstructor holds the string denoting the instructor edge name in mutations.
	EdgeInstructor = "instructor"
	// EdgeCourseclasses holds the string denoting the courseclasses edge name in mutations.
	EdgeCourseclasses = "courseclasses"

	// Table holds the table name of the instructorinfo in the database.
	Table = "instructor_infos"
	// TitleTable is the table the holds the title relation/edge.
	TitleTable = "instructor_infos"
	// TitleInverseTable is the table name for the Title entity.
	// It exists in this package in order to avoid circular dependency with the "title" package.
	TitleInverseTable = "titles"
	// TitleColumn is the table column denoting the title relation/edge.
	TitleColumn = "title_id"
	// InstructorroomTable is the table the holds the instructorroom relation/edge.
	InstructorroomTable = "instructor_infos"
	// InstructorroomInverseTable is the table name for the InstructorRoom entity.
	// It exists in this package in order to avoid circular dependency with the "instructorroom" package.
	InstructorroomInverseTable = "instructor_rooms"
	// InstructorroomColumn is the table column denoting the instructorroom relation/edge.
	InstructorroomColumn = "instructorroom_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "instructor_infos"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// InstructorTable is the table the holds the instructor relation/edge.
	InstructorTable = "courses"
	// InstructorInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	InstructorInverseTable = "courses"
	// InstructorColumn is the table column denoting the instructor relation/edge.
	InstructorColumn = "InstructorInfo_id"
	// CourseclassesTable is the table the holds the courseclasses relation/edge.
	CourseclassesTable = "courseclasses"
	// CourseclassesInverseTable is the table name for the Courseclass entity.
	// It exists in this package in order to avoid circular dependency with the "courseclass" package.
	CourseclassesInverseTable = "courseclasses"
	// CourseclassesColumn is the table column denoting the courseclasses relation/edge.
	CourseclassesColumn = "InstructorInfo_id"
)

// Columns holds all SQL columns for instructorinfo fields.
var Columns = []string{
	FieldID,
	FieldNAME,
	FieldPHONENUMBER,
	FieldEMAIL,
	FieldPASSWORD,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the InstructorInfo type.
var ForeignKeys = []string{
	"department_id",
	"instructorroom_id",
	"title_id",
}