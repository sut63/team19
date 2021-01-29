// Code generated by entc, DO NOT EDIT.

package courseclass

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team19/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Tablecode applies equality check predicate on the "tablecode" field. It's identical to TablecodeEQ.
func Tablecode(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTablecode), v))
	})
}

// TablecodeEQ applies the EQ predicate on the "tablecode" field.
func TablecodeEQ(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTablecode), v))
	})
}

// TablecodeNEQ applies the NEQ predicate on the "tablecode" field.
func TablecodeNEQ(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTablecode), v))
	})
}

// TablecodeIn applies the In predicate on the "tablecode" field.
func TablecodeIn(vs ...string) predicate.Courseclass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Courseclass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTablecode), v...))
	})
}

// TablecodeNotIn applies the NotIn predicate on the "tablecode" field.
func TablecodeNotIn(vs ...string) predicate.Courseclass {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Courseclass(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTablecode), v...))
	})
}

// TablecodeGT applies the GT predicate on the "tablecode" field.
func TablecodeGT(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTablecode), v))
	})
}

// TablecodeGTE applies the GTE predicate on the "tablecode" field.
func TablecodeGTE(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTablecode), v))
	})
}

// TablecodeLT applies the LT predicate on the "tablecode" field.
func TablecodeLT(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTablecode), v))
	})
}

// TablecodeLTE applies the LTE predicate on the "tablecode" field.
func TablecodeLTE(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTablecode), v))
	})
}

// TablecodeContains applies the Contains predicate on the "tablecode" field.
func TablecodeContains(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTablecode), v))
	})
}

// TablecodeHasPrefix applies the HasPrefix predicate on the "tablecode" field.
func TablecodeHasPrefix(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTablecode), v))
	})
}

// TablecodeHasSuffix applies the HasSuffix predicate on the "tablecode" field.
func TablecodeHasSuffix(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTablecode), v))
	})
}

// TablecodeEqualFold applies the EqualFold predicate on the "tablecode" field.
func TablecodeEqualFold(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTablecode), v))
	})
}

// TablecodeContainsFold applies the ContainsFold predicate on the "tablecode" field.
func TablecodeContainsFold(v string) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTablecode), v))
	})
}

// HasClasstime applies the HasEdge predicate on the "classtime" edge.
func HasClasstime() predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClasstimeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClasstimeTable, ClasstimeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClasstimeWith applies the HasEdge predicate on the "classtime" edge with a given conditions (other predicates).
func HasClasstimeWith(preds ...predicate.Classtime) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClasstimeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClasstimeTable, ClasstimeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassdate applies the HasEdge predicate on the "classdate" edge.
func HasClassdate() predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassdateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassdateTable, ClassdateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassdateWith applies the HasEdge predicate on the "classdate" edge with a given conditions (other predicates).
func HasClassdateWith(preds ...predicate.Classdate) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassdateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassdateTable, ClassdateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassroom applies the HasEdge predicate on the "classroom" edge.
func HasClassroom() predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassroomTable, ClassroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassroomWith applies the HasEdge predicate on the "classroom" edge with a given conditions (other predicates).
func HasClassroomWith(preds ...predicate.Classroom) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassroomTable, ClassroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstructorInfo applies the HasEdge predicate on the "instructorInfo" edge.
func HasInstructorInfo() predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorInfoTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorInfoTable, InstructorInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstructorInfoWith applies the HasEdge predicate on the "instructorInfo" edge with a given conditions (other predicates).
func HasInstructorInfoWith(preds ...predicate.InstructorInfo) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorInfoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorInfoTable, InstructorInfoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Courseclass) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Courseclass) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
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
func Not(p predicate.Courseclass) predicate.Courseclass {
	return predicate.Courseclass(func(s *sql.Selector) {
		p(s.Not())
	})
}
