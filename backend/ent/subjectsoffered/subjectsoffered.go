// Code generated by entc, DO NOT EDIT.

package subjectsoffered

const (
	// Label holds the string label denoting the subjectsoffered type in the database.
	Label = "subjects_offered"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAMOUNT holds the string denoting the amount field in the database.
	FieldAMOUNT = "amount"
	// FieldSTATUS holds the string denoting the status field in the database.
	FieldSTATUS = "status"

	// EdgeSubject holds the string denoting the subject edge name in mutations.
	EdgeSubject = "Subject"
	// EdgeDegree holds the string denoting the degree edge name in mutations.
	EdgeDegree = "Degree"
	// EdgeYear holds the string denoting the year edge name in mutations.
	EdgeYear = "Year"
	// EdgeTerm holds the string denoting the term edge name in mutations.
	EdgeTerm = "Term"

	// Table holds the table name of the subjectsoffered in the database.
	Table = "subjects_offereds"
	// SubjectTable is the table the holds the Subject relation/edge.
	SubjectTable = "subjects_offereds"
	// SubjectInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectInverseTable = "subjects"
	// SubjectColumn is the table column denoting the Subject relation/edge.
	SubjectColumn = "Subject_id"
	// DegreeTable is the table the holds the Degree relation/edge.
	DegreeTable = "subjects_offereds"
	// DegreeInverseTable is the table name for the Degree entity.
	// It exists in this package in order to avoid circular dependency with the "degree" package.
	DegreeInverseTable = "degrees"
	// DegreeColumn is the table column denoting the Degree relation/edge.
	DegreeColumn = "Degree_id"
	// YearTable is the table the holds the Year relation/edge.
	YearTable = "subjects_offereds"
	// YearInverseTable is the table name for the Year entity.
	// It exists in this package in order to avoid circular dependency with the "year" package.
	YearInverseTable = "years"
	// YearColumn is the table column denoting the Year relation/edge.
	YearColumn = "year_id"
	// TermTable is the table the holds the Term relation/edge.
	TermTable = "subjects_offereds"
	// TermInverseTable is the table name for the Term entity.
	// It exists in this package in order to avoid circular dependency with the "term" package.
	TermInverseTable = "terms"
	// TermColumn is the table column denoting the Term relation/edge.
	TermColumn = "term_id"
)

// Columns holds all SQL columns for subjectsoffered fields.
var Columns = []string{
	FieldID,
	FieldAMOUNT,
	FieldSTATUS,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the SubjectsOffered type.
var ForeignKeys = []string{
	"Degree_id",
	"Subject_id",
	"term_id",
	"year_id",
}

var (
	// AMOUNTValidator is a validator for the "AMOUNT" field. It is called by the builders before save.
	AMOUNTValidator func(string) error
	// STATUSValidator is a validator for the "STATUS" field. It is called by the builders before save.
	STATUSValidator func(string) error
)