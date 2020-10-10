// Code generated by entc, DO NOT EDIT.

package voca

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"lang.pkg/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Example applies equality check predicate on the "example" field. It's identical to ExampleEQ.
func Example(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExample), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// ExampleEQ applies the EQ predicate on the "example" field.
func ExampleEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExample), v))
	})
}

// ExampleNEQ applies the NEQ predicate on the "example" field.
func ExampleNEQ(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExample), v))
	})
}

// ExampleIn applies the In predicate on the "example" field.
func ExampleIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExample), v...))
	})
}

// ExampleNotIn applies the NotIn predicate on the "example" field.
func ExampleNotIn(vs ...string) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExample), v...))
	})
}

// ExampleGT applies the GT predicate on the "example" field.
func ExampleGT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExample), v))
	})
}

// ExampleGTE applies the GTE predicate on the "example" field.
func ExampleGTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExample), v))
	})
}

// ExampleLT applies the LT predicate on the "example" field.
func ExampleLT(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExample), v))
	})
}

// ExampleLTE applies the LTE predicate on the "example" field.
func ExampleLTE(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExample), v))
	})
}

// ExampleContains applies the Contains predicate on the "example" field.
func ExampleContains(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExample), v))
	})
}

// ExampleHasPrefix applies the HasPrefix predicate on the "example" field.
func ExampleHasPrefix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExample), v))
	})
}

// ExampleHasSuffix applies the HasSuffix predicate on the "example" field.
func ExampleHasSuffix(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExample), v))
	})
}

// ExampleEqualFold applies the EqualFold predicate on the "example" field.
func ExampleEqualFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExample), v))
	})
}

// ExampleContainsFold applies the ContainsFold predicate on the "example" field.
func ExampleContainsFold(v string) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExample), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Voca {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Voca(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Voca) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Voca) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
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
func Not(p predicate.Voca) predicate.Voca {
	return predicate.Voca(func(s *sql.Selector) {
		p(s.Not())
	})
}
