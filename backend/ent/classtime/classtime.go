// Code generated by entc, DO NOT EDIT.

package classtime

const (
	// Label holds the string label denoting the classtime type in the database.
	Label = "classtime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTIME holds the string denoting the time field in the database.
	FieldTIME = "time"

	// EdgeCourseclasses holds the string denoting the courseclasses edge name in mutations.
	EdgeCourseclasses = "courseclasses"

	// Table holds the table name of the classtime in the database.
	Table = "classtimes"
	// CourseclassesTable is the table the holds the courseclasses relation/edge.
	CourseclassesTable = "courseclasses"
	// CourseclassesInverseTable is the table name for the Courseclass entity.
	// It exists in this package in order to avoid circular dependency with the "courseclass" package.
	CourseclassesInverseTable = "courseclasses"
	// CourseclassesColumn is the table column denoting the courseclasses relation/edge.
	CourseclassesColumn = "classtime_id"
)

// Columns holds all SQL columns for classtime fields.
var Columns = []string{
	FieldID,
	FieldTIME,
}
