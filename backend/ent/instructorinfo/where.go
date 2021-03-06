// Code generated by entc, DO NOT EDIT.

package instructorinfo

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team19/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
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
func IDGT(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// NAME applies equality check predicate on the "NAME" field. It's identical to NAMEEQ.
func NAME(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNAME), v))
	})
}

// PHONENUMBER applies equality check predicate on the "PHONENUMBER" field. It's identical to PHONENUMBEREQ.
func PHONENUMBER(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// EMAIL applies equality check predicate on the "EMAIL" field. It's identical to EMAILEQ.
func EMAIL(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMAIL), v))
	})
}

// PASSWORD applies equality check predicate on the "PASSWORD" field. It's identical to PASSWORDEQ.
func PASSWORD(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPASSWORD), v))
	})
}

// NAMEEQ applies the EQ predicate on the "NAME" field.
func NAMEEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNAME), v))
	})
}

// NAMENEQ applies the NEQ predicate on the "NAME" field.
func NAMENEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNAME), v))
	})
}

// NAMEIn applies the In predicate on the "NAME" field.
func NAMEIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNAME), v...))
	})
}

// NAMENotIn applies the NotIn predicate on the "NAME" field.
func NAMENotIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNAME), v...))
	})
}

// NAMEGT applies the GT predicate on the "NAME" field.
func NAMEGT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNAME), v))
	})
}

// NAMEGTE applies the GTE predicate on the "NAME" field.
func NAMEGTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNAME), v))
	})
}

// NAMELT applies the LT predicate on the "NAME" field.
func NAMELT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNAME), v))
	})
}

// NAMELTE applies the LTE predicate on the "NAME" field.
func NAMELTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNAME), v))
	})
}

// NAMEContains applies the Contains predicate on the "NAME" field.
func NAMEContains(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNAME), v))
	})
}

// NAMEHasPrefix applies the HasPrefix predicate on the "NAME" field.
func NAMEHasPrefix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNAME), v))
	})
}

// NAMEHasSuffix applies the HasSuffix predicate on the "NAME" field.
func NAMEHasSuffix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNAME), v))
	})
}

// NAMEEqualFold applies the EqualFold predicate on the "NAME" field.
func NAMEEqualFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNAME), v))
	})
}

// NAMEContainsFold applies the ContainsFold predicate on the "NAME" field.
func NAMEContainsFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNAME), v))
	})
}

// PHONENUMBEREQ applies the EQ predicate on the "PHONENUMBER" field.
func PHONENUMBEREQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERNEQ applies the NEQ predicate on the "PHONENUMBER" field.
func PHONENUMBERNEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERIn applies the In predicate on the "PHONENUMBER" field.
func PHONENUMBERIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERNotIn applies the NotIn predicate on the "PHONENUMBER" field.
func PHONENUMBERNotIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHONENUMBER), v...))
	})
}

// PHONENUMBERGT applies the GT predicate on the "PHONENUMBER" field.
func PHONENUMBERGT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERGTE applies the GTE predicate on the "PHONENUMBER" field.
func PHONENUMBERGTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLT applies the LT predicate on the "PHONENUMBER" field.
func PHONENUMBERLT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERLTE applies the LTE predicate on the "PHONENUMBER" field.
func PHONENUMBERLTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContains applies the Contains predicate on the "PHONENUMBER" field.
func PHONENUMBERContains(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasPrefix applies the HasPrefix predicate on the "PHONENUMBER" field.
func PHONENUMBERHasPrefix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERHasSuffix applies the HasSuffix predicate on the "PHONENUMBER" field.
func PHONENUMBERHasSuffix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBEREqualFold applies the EqualFold predicate on the "PHONENUMBER" field.
func PHONENUMBEREqualFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHONENUMBER), v))
	})
}

// PHONENUMBERContainsFold applies the ContainsFold predicate on the "PHONENUMBER" field.
func PHONENUMBERContainsFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHONENUMBER), v))
	})
}

// EMAILEQ applies the EQ predicate on the "EMAIL" field.
func EMAILEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEMAIL), v))
	})
}

// EMAILNEQ applies the NEQ predicate on the "EMAIL" field.
func EMAILNEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEMAIL), v))
	})
}

// EMAILIn applies the In predicate on the "EMAIL" field.
func EMAILIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEMAIL), v...))
	})
}

// EMAILNotIn applies the NotIn predicate on the "EMAIL" field.
func EMAILNotIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEMAIL), v...))
	})
}

// EMAILGT applies the GT predicate on the "EMAIL" field.
func EMAILGT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEMAIL), v))
	})
}

// EMAILGTE applies the GTE predicate on the "EMAIL" field.
func EMAILGTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEMAIL), v))
	})
}

// EMAILLT applies the LT predicate on the "EMAIL" field.
func EMAILLT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEMAIL), v))
	})
}

// EMAILLTE applies the LTE predicate on the "EMAIL" field.
func EMAILLTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEMAIL), v))
	})
}

// EMAILContains applies the Contains predicate on the "EMAIL" field.
func EMAILContains(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEMAIL), v))
	})
}

// EMAILHasPrefix applies the HasPrefix predicate on the "EMAIL" field.
func EMAILHasPrefix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEMAIL), v))
	})
}

// EMAILHasSuffix applies the HasSuffix predicate on the "EMAIL" field.
func EMAILHasSuffix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEMAIL), v))
	})
}

// EMAILEqualFold applies the EqualFold predicate on the "EMAIL" field.
func EMAILEqualFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEMAIL), v))
	})
}

// EMAILContainsFold applies the ContainsFold predicate on the "EMAIL" field.
func EMAILContainsFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEMAIL), v))
	})
}

// PASSWORDEQ applies the EQ predicate on the "PASSWORD" field.
func PASSWORDEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDNEQ applies the NEQ predicate on the "PASSWORD" field.
func PASSWORDNEQ(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDIn applies the In predicate on the "PASSWORD" field.
func PASSWORDIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPASSWORD), v...))
	})
}

// PASSWORDNotIn applies the NotIn predicate on the "PASSWORD" field.
func PASSWORDNotIn(vs ...string) predicate.InstructorInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstructorInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPASSWORD), v...))
	})
}

// PASSWORDGT applies the GT predicate on the "PASSWORD" field.
func PASSWORDGT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDGTE applies the GTE predicate on the "PASSWORD" field.
func PASSWORDGTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDLT applies the LT predicate on the "PASSWORD" field.
func PASSWORDLT(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDLTE applies the LTE predicate on the "PASSWORD" field.
func PASSWORDLTE(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDContains applies the Contains predicate on the "PASSWORD" field.
func PASSWORDContains(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDHasPrefix applies the HasPrefix predicate on the "PASSWORD" field.
func PASSWORDHasPrefix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDHasSuffix applies the HasSuffix predicate on the "PASSWORD" field.
func PASSWORDHasSuffix(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDEqualFold applies the EqualFold predicate on the "PASSWORD" field.
func PASSWORDEqualFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPASSWORD), v))
	})
}

// PASSWORDContainsFold applies the ContainsFold predicate on the "PASSWORD" field.
func PASSWORDContainsFold(v string) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPASSWORD), v))
	})
}

// HasTitle applies the HasEdge predicate on the "title" edge.
func HasTitle() predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleWith applies the HasEdge predicate on the "title" edge with a given conditions (other predicates).
func HasTitleWith(preds ...predicate.Title) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstructorroom applies the HasEdge predicate on the "instructorroom" edge.
func HasInstructorroom() predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorroomTable, InstructorroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstructorroomWith applies the HasEdge predicate on the "instructorroom" edge with a given conditions (other predicates).
func HasInstructorroomWith(preds ...predicate.InstructorRoom) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorroomTable, InstructorroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourseclasses applies the HasEdge predicate on the "courseclasses" edge.
func HasCourseclasses() predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseclassesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseclassesTable, CourseclassesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseclassesWith applies the HasEdge predicate on the "courseclasses" edge with a given conditions (other predicates).
func HasCourseclassesWith(preds ...predicate.Courseclass) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseclassesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseclassesTable, CourseclassesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.InstructorInfo) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.InstructorInfo) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
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
func Not(p predicate.InstructorInfo) predicate.InstructorInfo {
	return predicate.InstructorInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
