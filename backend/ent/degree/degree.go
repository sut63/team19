// Code generated by entc, DO NOT EDIT.

package degree

const (
	// Label holds the string label denoting the degree type in the database.
	Label = "degree"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDegreeName holds the string denoting the degree_name field in the database.
	FieldDegreeName = "degree_name"

	// EdgeDegree holds the string denoting the degree edge name in mutations.
	EdgeDegree = "degree"

	// Table holds the table name of the degree in the database.
	Table = "degrees"
	// DegreeTable is the table the holds the degree relation/edge.
	DegreeTable = "courses"
	// DegreeInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	DegreeInverseTable = "courses"
	// DegreeColumn is the table column denoting the degree relation/edge.
	DegreeColumn = "Degree_id"
)

// Columns holds all SQL columns for degree fields.
var Columns = []string{
	FieldID,
	FieldDegreeName,
}
