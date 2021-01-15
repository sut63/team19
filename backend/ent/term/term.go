// Code generated by entc, DO NOT EDIT.

package term

const (
	// Label holds the string label denoting the term type in the database.
	Label = "term"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTERM holds the string denoting the term field in the database.
	FieldTERM = "term"

	// EdgeSubjectsOffered holds the string denoting the subjectsoffered edge name in mutations.
	EdgeSubjectsOffered = "SubjectsOffered"

	// Table holds the table name of the term in the database.
	Table = "terms"
	// SubjectsOfferedTable is the table the holds the SubjectsOffered relation/edge.
	SubjectsOfferedTable = "subjects_offereds"
	// SubjectsOfferedInverseTable is the table name for the SubjectsOffered entity.
	// It exists in this package in order to avoid circular dependency with the "subjectsoffered" package.
	SubjectsOfferedInverseTable = "subjects_offereds"
	// SubjectsOfferedColumn is the table column denoting the SubjectsOffered relation/edge.
	SubjectsOfferedColumn = "term_id"
)

// Columns holds all SQL columns for term fields.
var Columns = []string{
	FieldID,
	FieldTERM,
}
