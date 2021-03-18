// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ClassdatesColumns holds the columns for the "classdates" table.
	ClassdatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "day", Type: field.TypeString, Unique: true},
	}
	// ClassdatesTable holds the schema information for the "classdates" table.
	ClassdatesTable = &schema.Table{
		Name:        "classdates",
		Columns:     ClassdatesColumns,
		PrimaryKey:  []*schema.Column{ClassdatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClassroomsColumns holds the columns for the "classrooms" table.
	ClassroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "room", Type: field.TypeString, Unique: true},
	}
	// ClassroomsTable holds the schema information for the "classrooms" table.
	ClassroomsTable = &schema.Table{
		Name:        "classrooms",
		Columns:     ClassroomsColumns,
		PrimaryKey:  []*schema.Column{ClassroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClasstimesColumns holds the columns for the "classtimes" table.
	ClasstimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeString, Unique: true},
	}
	// ClasstimesTable holds the schema information for the "classtimes" table.
	ClasstimesTable = &schema.Table{
		Name:        "classtimes",
		Columns:     ClasstimesColumns,
		PrimaryKey:  []*schema.Column{ClasstimesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "course_year", Type: field.TypeInt},
		{Name: "course_name", Type: field.TypeString},
		{Name: "teacher_id", Type: field.TypeString},
		{Name: "Degree_id", Type: field.TypeInt, Nullable: true},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "Subject_id", Type: field.TypeInt, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "courses_degrees_degree",
				Columns: []*schema.Column{CoursesColumns[4]},

				RefColumns: []*schema.Column{DegreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courses_departments_department",
				Columns: []*schema.Column{CoursesColumns[5]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courses_subjects_subject",
				Columns: []*schema.Column{CoursesColumns[6]},

				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CourseclassesColumns holds the columns for the "courseclasses" table.
	CourseclassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tablecode", Type: field.TypeString},
		{Name: "group_class", Type: field.TypeString},
		{Name: "annotation", Type: field.TypeString},
		{Name: "classdate_id", Type: field.TypeInt, Nullable: true},
		{Name: "classroom_id", Type: field.TypeInt, Nullable: true},
		{Name: "classtime_id", Type: field.TypeInt, Nullable: true},
		{Name: "InstructorInfo_id", Type: field.TypeInt, Nullable: true},
		{Name: "subject_id", Type: field.TypeInt, Nullable: true},
	}
	// CourseclassesTable holds the schema information for the "courseclasses" table.
	CourseclassesTable = &schema.Table{
		Name:       "courseclasses",
		Columns:    CourseclassesColumns,
		PrimaryKey: []*schema.Column{CourseclassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "courseclasses_classdates_courseclasses",
				Columns: []*schema.Column{CourseclassesColumns[4]},

				RefColumns: []*schema.Column{ClassdatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courseclasses_classrooms_courseclasses",
				Columns: []*schema.Column{CourseclassesColumns[5]},

				RefColumns: []*schema.Column{ClassroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courseclasses_classtimes_courseclasses",
				Columns: []*schema.Column{CourseclassesColumns[6]},

				RefColumns: []*schema.Column{ClasstimesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courseclasses_instructor_infos_courseclasses",
				Columns: []*schema.Column{CourseclassesColumns[7]},

				RefColumns: []*schema.Column{InstructorInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "courseclasses_subjects_courseclasses",
				Columns: []*schema.Column{CourseclassesColumns[8]},

				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DegreesColumns holds the columns for the "degrees" table.
	DegreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "degree_name", Type: field.TypeString, Unique: true},
	}
	// DegreesTable holds the schema information for the "degrees" table.
	DegreesTable = &schema.Table{
		Name:        "degrees",
		Columns:     DegreesColumns,
		PrimaryKey:  []*schema.Column{DegreesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department", Type: field.TypeString, Unique: true},
		{Name: "faculty", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// InstructorInfosColumns holds the columns for the "instructor_infos" table.
	InstructorInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "phonenumber", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "instructorroom_id", Type: field.TypeInt, Nullable: true},
		{Name: "title_id", Type: field.TypeInt, Nullable: true},
	}
	// InstructorInfosTable holds the schema information for the "instructor_infos" table.
	InstructorInfosTable = &schema.Table{
		Name:       "instructor_infos",
		Columns:    InstructorInfosColumns,
		PrimaryKey: []*schema.Column{InstructorInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "instructor_infos_departments_instructorinfos",
				Columns: []*schema.Column{InstructorInfosColumns[5]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "instructor_infos_instructor_rooms_instructorinfos",
				Columns: []*schema.Column{InstructorInfosColumns[6]},

				RefColumns: []*schema.Column{InstructorRoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "instructor_infos_titles_instructorinfos",
				Columns: []*schema.Column{InstructorInfosColumns[7]},

				RefColumns: []*schema.Column{TitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InstructorRoomsColumns holds the columns for the "instructor_rooms" table.
	InstructorRoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "room", Type: field.TypeString, Unique: true},
		{Name: "building", Type: field.TypeString},
	}
	// InstructorRoomsTable holds the schema information for the "instructor_rooms" table.
	InstructorRoomsTable = &schema.Table{
		Name:        "instructor_rooms",
		Columns:     InstructorRoomsColumns,
		PrimaryKey:  []*schema.Column{InstructorRoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "subject_name", Type: field.TypeString, Unique: true},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:        "subjects",
		Columns:     SubjectsColumns,
		PrimaryKey:  []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SubjectsOfferedsColumns holds the columns for the "subjects_offereds" table.
	SubjectsOfferedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
		{Name: "status", Type: field.TypeBool},
		{Name: "remain", Type: field.TypeInt},
		{Name: "Degree_id", Type: field.TypeInt, Nullable: true},
		{Name: "Subject_id", Type: field.TypeInt, Nullable: true},
		{Name: "term_id", Type: field.TypeInt, Nullable: true},
		{Name: "year_id", Type: field.TypeInt, Nullable: true},
	}
	// SubjectsOfferedsTable holds the schema information for the "subjects_offereds" table.
	SubjectsOfferedsTable = &schema.Table{
		Name:       "subjects_offereds",
		Columns:    SubjectsOfferedsColumns,
		PrimaryKey: []*schema.Column{SubjectsOfferedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "subjects_offereds_degrees_SubjectsOffered",
				Columns: []*schema.Column{SubjectsOfferedsColumns[4]},

				RefColumns: []*schema.Column{DegreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "subjects_offereds_subjects_SubjectsOffered",
				Columns: []*schema.Column{SubjectsOfferedsColumns[5]},

				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "subjects_offereds_terms_SubjectsOffered",
				Columns: []*schema.Column{SubjectsOfferedsColumns[6]},

				RefColumns: []*schema.Column{TermsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "subjects_offereds_years_SubjectsOffered",
				Columns: []*schema.Column{SubjectsOfferedsColumns[7]},

				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TermsColumns holds the columns for the "terms" table.
	TermsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "term", Type: field.TypeInt, Unique: true},
	}
	// TermsTable holds the schema information for the "terms" table.
	TermsTable = &schema.Table{
		Name:        "terms",
		Columns:     TermsColumns,
		PrimaryKey:  []*schema.Column{TermsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TitlesColumns holds the columns for the "titles" table.
	TitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
	}
	// TitlesTable holds the schema information for the "titles" table.
	TitlesTable = &schema.Table{
		Name:        "titles",
		Columns:     TitlesColumns,
		PrimaryKey:  []*schema.Column{TitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// YearsColumns holds the columns for the "years" table.
	YearsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "year", Type: field.TypeInt, Unique: true},
	}
	// YearsTable holds the schema information for the "years" table.
	YearsTable = &schema.Table{
		Name:        "years",
		Columns:     YearsColumns,
		PrimaryKey:  []*schema.Column{YearsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassdatesTable,
		ClassroomsTable,
		ClasstimesTable,
		CoursesTable,
		CourseclassesTable,
		DegreesTable,
		DepartmentsTable,
		InstructorInfosTable,
		InstructorRoomsTable,
		SubjectsTable,
		SubjectsOfferedsTable,
		TermsTable,
		TitlesTable,
		YearsTable,
	}
)

func init() {
	CoursesTable.ForeignKeys[0].RefTable = DegreesTable
	CoursesTable.ForeignKeys[1].RefTable = DepartmentsTable
	CoursesTable.ForeignKeys[2].RefTable = SubjectsTable
	CourseclassesTable.ForeignKeys[0].RefTable = ClassdatesTable
	CourseclassesTable.ForeignKeys[1].RefTable = ClassroomsTable
	CourseclassesTable.ForeignKeys[2].RefTable = ClasstimesTable
	CourseclassesTable.ForeignKeys[3].RefTable = InstructorInfosTable
	CourseclassesTable.ForeignKeys[4].RefTable = SubjectsTable
	InstructorInfosTable.ForeignKeys[0].RefTable = DepartmentsTable
	InstructorInfosTable.ForeignKeys[1].RefTable = InstructorRoomsTable
	InstructorInfosTable.ForeignKeys[2].RefTable = TitlesTable
	SubjectsOfferedsTable.ForeignKeys[0].RefTable = DegreesTable
	SubjectsOfferedsTable.ForeignKeys[1].RefTable = SubjectsTable
	SubjectsOfferedsTable.ForeignKeys[2].RefTable = TermsTable
	SubjectsOfferedsTable.ForeignKeys[3].RefTable = YearsTable
}
