// Code generated by entc, DO NOT EDIT.

package cronjob

import (
	"htp-platform/app/machine/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MachineID applies equality check predicate on the "machine_id" field. It's identical to MachineIDEQ.
func MachineID(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// CheckName applies equality check predicate on the "check_name" field. It's identical to CheckNameEQ.
func CheckName(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckName), v))
	})
}

// CronString applies equality check predicate on the "cron_string" field. It's identical to CronStringEQ.
func CronString(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCronString), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MachineIDEQ applies the EQ predicate on the "machine_id" field.
func MachineIDEQ(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// MachineIDNEQ applies the NEQ predicate on the "machine_id" field.
func MachineIDNEQ(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMachineID), v))
	})
}

// MachineIDIn applies the In predicate on the "machine_id" field.
func MachineIDIn(vs ...uuid.UUID) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMachineID), v...))
	})
}

// MachineIDNotIn applies the NotIn predicate on the "machine_id" field.
func MachineIDNotIn(vs ...uuid.UUID) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMachineID), v...))
	})
}

// MachineIDGT applies the GT predicate on the "machine_id" field.
func MachineIDGT(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMachineID), v))
	})
}

// MachineIDGTE applies the GTE predicate on the "machine_id" field.
func MachineIDGTE(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMachineID), v))
	})
}

// MachineIDLT applies the LT predicate on the "machine_id" field.
func MachineIDLT(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMachineID), v))
	})
}

// MachineIDLTE applies the LTE predicate on the "machine_id" field.
func MachineIDLTE(v uuid.UUID) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMachineID), v))
	})
}

// CheckNameEQ applies the EQ predicate on the "check_name" field.
func CheckNameEQ(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckName), v))
	})
}

// CheckNameNEQ applies the NEQ predicate on the "check_name" field.
func CheckNameNEQ(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCheckName), v))
	})
}

// CheckNameIn applies the In predicate on the "check_name" field.
func CheckNameIn(vs ...string) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCheckName), v...))
	})
}

// CheckNameNotIn applies the NotIn predicate on the "check_name" field.
func CheckNameNotIn(vs ...string) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCheckName), v...))
	})
}

// CheckNameGT applies the GT predicate on the "check_name" field.
func CheckNameGT(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCheckName), v))
	})
}

// CheckNameGTE applies the GTE predicate on the "check_name" field.
func CheckNameGTE(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCheckName), v))
	})
}

// CheckNameLT applies the LT predicate on the "check_name" field.
func CheckNameLT(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCheckName), v))
	})
}

// CheckNameLTE applies the LTE predicate on the "check_name" field.
func CheckNameLTE(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCheckName), v))
	})
}

// CheckNameContains applies the Contains predicate on the "check_name" field.
func CheckNameContains(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCheckName), v))
	})
}

// CheckNameHasPrefix applies the HasPrefix predicate on the "check_name" field.
func CheckNameHasPrefix(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCheckName), v))
	})
}

// CheckNameHasSuffix applies the HasSuffix predicate on the "check_name" field.
func CheckNameHasSuffix(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCheckName), v))
	})
}

// CheckNameEqualFold applies the EqualFold predicate on the "check_name" field.
func CheckNameEqualFold(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCheckName), v))
	})
}

// CheckNameContainsFold applies the ContainsFold predicate on the "check_name" field.
func CheckNameContainsFold(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCheckName), v))
	})
}

// CronStringEQ applies the EQ predicate on the "cron_string" field.
func CronStringEQ(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCronString), v))
	})
}

// CronStringNEQ applies the NEQ predicate on the "cron_string" field.
func CronStringNEQ(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCronString), v))
	})
}

// CronStringIn applies the In predicate on the "cron_string" field.
func CronStringIn(vs ...string) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCronString), v...))
	})
}

// CronStringNotIn applies the NotIn predicate on the "cron_string" field.
func CronStringNotIn(vs ...string) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCronString), v...))
	})
}

// CronStringGT applies the GT predicate on the "cron_string" field.
func CronStringGT(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCronString), v))
	})
}

// CronStringGTE applies the GTE predicate on the "cron_string" field.
func CronStringGTE(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCronString), v))
	})
}

// CronStringLT applies the LT predicate on the "cron_string" field.
func CronStringLT(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCronString), v))
	})
}

// CronStringLTE applies the LTE predicate on the "cron_string" field.
func CronStringLTE(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCronString), v))
	})
}

// CronStringContains applies the Contains predicate on the "cron_string" field.
func CronStringContains(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCronString), v))
	})
}

// CronStringHasPrefix applies the HasPrefix predicate on the "cron_string" field.
func CronStringHasPrefix(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCronString), v))
	})
}

// CronStringHasSuffix applies the HasSuffix predicate on the "cron_string" field.
func CronStringHasSuffix(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCronString), v))
	})
}

// CronStringEqualFold applies the EqualFold predicate on the "cron_string" field.
func CronStringEqualFold(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCronString), v))
	})
}

// CronStringContainsFold applies the ContainsFold predicate on the "cron_string" field.
func CronStringContainsFold(v string) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCronString), v))
	})
}

// CoordinatesIsNil applies the IsNil predicate on the "coordinates" field.
func CoordinatesIsNil() predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoordinates)))
	})
}

// CoordinatesNotNil applies the NotNil predicate on the "coordinates" field.
func CoordinatesNotNil() predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoordinates)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CronJob {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronJob(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CronJob) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CronJob) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
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
func Not(p predicate.CronJob) predicate.CronJob {
	return predicate.CronJob(func(s *sql.Selector) {
		p(s.Not())
	})
}
