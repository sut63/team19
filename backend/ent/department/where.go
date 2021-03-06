// Code generated by entc, DO NOT EDIT.

package department

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team19/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DEPARTMENT applies equality check predicate on the "DEPARTMENT" field. It's identical to DEPARTMENTEQ.
func DEPARTMENT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDEPARTMENT), v))
	})
}

// FACULTY applies equality check predicate on the "FACULTY" field. It's identical to FACULTYEQ.
func FACULTY(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFACULTY), v))
	})
}

// DEPARTMENTEQ applies the EQ predicate on the "DEPARTMENT" field.
func DEPARTMENTEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTNEQ applies the NEQ predicate on the "DEPARTMENT" field.
func DEPARTMENTNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTIn applies the In predicate on the "DEPARTMENT" field.
func DEPARTMENTIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDEPARTMENT), v...))
	})
}

// DEPARTMENTNotIn applies the NotIn predicate on the "DEPARTMENT" field.
func DEPARTMENTNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDEPARTMENT), v...))
	})
}

// DEPARTMENTGT applies the GT predicate on the "DEPARTMENT" field.
func DEPARTMENTGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTGTE applies the GTE predicate on the "DEPARTMENT" field.
func DEPARTMENTGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTLT applies the LT predicate on the "DEPARTMENT" field.
func DEPARTMENTLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTLTE applies the LTE predicate on the "DEPARTMENT" field.
func DEPARTMENTLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTContains applies the Contains predicate on the "DEPARTMENT" field.
func DEPARTMENTContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTHasPrefix applies the HasPrefix predicate on the "DEPARTMENT" field.
func DEPARTMENTHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTHasSuffix applies the HasSuffix predicate on the "DEPARTMENT" field.
func DEPARTMENTHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTEqualFold applies the EqualFold predicate on the "DEPARTMENT" field.
func DEPARTMENTEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDEPARTMENT), v))
	})
}

// DEPARTMENTContainsFold applies the ContainsFold predicate on the "DEPARTMENT" field.
func DEPARTMENTContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDEPARTMENT), v))
	})
}

// FACULTYEQ applies the EQ predicate on the "FACULTY" field.
func FACULTYEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFACULTY), v))
	})
}

// FACULTYNEQ applies the NEQ predicate on the "FACULTY" field.
func FACULTYNEQ(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFACULTY), v))
	})
}

// FACULTYIn applies the In predicate on the "FACULTY" field.
func FACULTYIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFACULTY), v...))
	})
}

// FACULTYNotIn applies the NotIn predicate on the "FACULTY" field.
func FACULTYNotIn(vs ...string) predicate.Department {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Department(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFACULTY), v...))
	})
}

// FACULTYGT applies the GT predicate on the "FACULTY" field.
func FACULTYGT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFACULTY), v))
	})
}

// FACULTYGTE applies the GTE predicate on the "FACULTY" field.
func FACULTYGTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFACULTY), v))
	})
}

// FACULTYLT applies the LT predicate on the "FACULTY" field.
func FACULTYLT(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFACULTY), v))
	})
}

// FACULTYLTE applies the LTE predicate on the "FACULTY" field.
func FACULTYLTE(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFACULTY), v))
	})
}

// FACULTYContains applies the Contains predicate on the "FACULTY" field.
func FACULTYContains(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFACULTY), v))
	})
}

// FACULTYHasPrefix applies the HasPrefix predicate on the "FACULTY" field.
func FACULTYHasPrefix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFACULTY), v))
	})
}

// FACULTYHasSuffix applies the HasSuffix predicate on the "FACULTY" field.
func FACULTYHasSuffix(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFACULTY), v))
	})
}

// FACULTYEqualFold applies the EqualFold predicate on the "FACULTY" field.
func FACULTYEqualFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFACULTY), v))
	})
}

// FACULTYContainsFold applies the ContainsFold predicate on the "FACULTY" field.
func FACULTYContainsFold(v string) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFACULTY), v))
	})
}

// HasInstructorinfos applies the HasEdge predicate on the "instructorinfos" edge.
func HasInstructorinfos() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorinfosTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstructorinfosTable, InstructorinfosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstructorinfosWith applies the HasEdge predicate on the "instructorinfos" edge with a given conditions (other predicates).
func HasInstructorinfosWith(preds ...predicate.InstructorInfo) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorinfosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstructorinfosTable, InstructorinfosColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Course) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(func(s *sql.Selector) {
		p(s.Not())
	})
}
