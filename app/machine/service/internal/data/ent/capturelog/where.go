// Code generated by entc, DO NOT EDIT.

package capturelog

import (
	"htp-platform/app/machine/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MachineID applies equality check predicate on the "machine_id" field. It's identical to MachineIDEQ.
func MachineID(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// Pixels applies equality check predicate on the "pixels" field. It's identical to PixelsEQ.
func Pixels(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPixels), v))
	})
}

// Area applies equality check predicate on the "area" field. It's identical to AreaEQ.
func Area(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArea), v))
	})
}

// ImageName applies equality check predicate on the "image_name" field. It's identical to ImageNameEQ.
func ImageName(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageName), v))
	})
}

// OssURL applies equality check predicate on the "oss_url" field. It's identical to OssURLEQ.
func OssURL(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOssURL), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MachineIDEQ applies the EQ predicate on the "machine_id" field.
func MachineIDEQ(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineID), v))
	})
}

// MachineIDNEQ applies the NEQ predicate on the "machine_id" field.
func MachineIDNEQ(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMachineID), v))
	})
}

// MachineIDIn applies the In predicate on the "machine_id" field.
func MachineIDIn(vs ...int64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func MachineIDNotIn(vs ...int64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func MachineIDGT(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMachineID), v))
	})
}

// MachineIDGTE applies the GTE predicate on the "machine_id" field.
func MachineIDGTE(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMachineID), v))
	})
}

// MachineIDLT applies the LT predicate on the "machine_id" field.
func MachineIDLT(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMachineID), v))
	})
}

// MachineIDLTE applies the LTE predicate on the "machine_id" field.
func MachineIDLTE(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMachineID), v))
	})
}

// PixelsEQ applies the EQ predicate on the "pixels" field.
func PixelsEQ(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPixels), v))
	})
}

// PixelsNEQ applies the NEQ predicate on the "pixels" field.
func PixelsNEQ(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPixels), v))
	})
}

// PixelsIn applies the In predicate on the "pixels" field.
func PixelsIn(vs ...int64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPixels), v...))
	})
}

// PixelsNotIn applies the NotIn predicate on the "pixels" field.
func PixelsNotIn(vs ...int64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPixels), v...))
	})
}

// PixelsGT applies the GT predicate on the "pixels" field.
func PixelsGT(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPixels), v))
	})
}

// PixelsGTE applies the GTE predicate on the "pixels" field.
func PixelsGTE(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPixels), v))
	})
}

// PixelsLT applies the LT predicate on the "pixels" field.
func PixelsLT(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPixels), v))
	})
}

// PixelsLTE applies the LTE predicate on the "pixels" field.
func PixelsLTE(v int64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPixels), v))
	})
}

// AreaEQ applies the EQ predicate on the "area" field.
func AreaEQ(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArea), v))
	})
}

// AreaNEQ applies the NEQ predicate on the "area" field.
func AreaNEQ(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArea), v))
	})
}

// AreaIn applies the In predicate on the "area" field.
func AreaIn(vs ...float64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArea), v...))
	})
}

// AreaNotIn applies the NotIn predicate on the "area" field.
func AreaNotIn(vs ...float64) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArea), v...))
	})
}

// AreaGT applies the GT predicate on the "area" field.
func AreaGT(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArea), v))
	})
}

// AreaGTE applies the GTE predicate on the "area" field.
func AreaGTE(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArea), v))
	})
}

// AreaLT applies the LT predicate on the "area" field.
func AreaLT(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArea), v))
	})
}

// AreaLTE applies the LTE predicate on the "area" field.
func AreaLTE(v float64) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArea), v))
	})
}

// ImageNameEQ applies the EQ predicate on the "image_name" field.
func ImageNameEQ(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageName), v))
	})
}

// ImageNameNEQ applies the NEQ predicate on the "image_name" field.
func ImageNameNEQ(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageName), v))
	})
}

// ImageNameIn applies the In predicate on the "image_name" field.
func ImageNameIn(vs ...string) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImageName), v...))
	})
}

// ImageNameNotIn applies the NotIn predicate on the "image_name" field.
func ImageNameNotIn(vs ...string) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImageName), v...))
	})
}

// ImageNameGT applies the GT predicate on the "image_name" field.
func ImageNameGT(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageName), v))
	})
}

// ImageNameGTE applies the GTE predicate on the "image_name" field.
func ImageNameGTE(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageName), v))
	})
}

// ImageNameLT applies the LT predicate on the "image_name" field.
func ImageNameLT(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageName), v))
	})
}

// ImageNameLTE applies the LTE predicate on the "image_name" field.
func ImageNameLTE(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageName), v))
	})
}

// ImageNameContains applies the Contains predicate on the "image_name" field.
func ImageNameContains(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageName), v))
	})
}

// ImageNameHasPrefix applies the HasPrefix predicate on the "image_name" field.
func ImageNameHasPrefix(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageName), v))
	})
}

// ImageNameHasSuffix applies the HasSuffix predicate on the "image_name" field.
func ImageNameHasSuffix(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageName), v))
	})
}

// ImageNameEqualFold applies the EqualFold predicate on the "image_name" field.
func ImageNameEqualFold(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageName), v))
	})
}

// ImageNameContainsFold applies the ContainsFold predicate on the "image_name" field.
func ImageNameContainsFold(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageName), v))
	})
}

// OssURLEQ applies the EQ predicate on the "oss_url" field.
func OssURLEQ(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOssURL), v))
	})
}

// OssURLNEQ applies the NEQ predicate on the "oss_url" field.
func OssURLNEQ(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOssURL), v))
	})
}

// OssURLIn applies the In predicate on the "oss_url" field.
func OssURLIn(vs ...string) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOssURL), v...))
	})
}

// OssURLNotIn applies the NotIn predicate on the "oss_url" field.
func OssURLNotIn(vs ...string) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOssURL), v...))
	})
}

// OssURLGT applies the GT predicate on the "oss_url" field.
func OssURLGT(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOssURL), v))
	})
}

// OssURLGTE applies the GTE predicate on the "oss_url" field.
func OssURLGTE(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOssURL), v))
	})
}

// OssURLLT applies the LT predicate on the "oss_url" field.
func OssURLLT(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOssURL), v))
	})
}

// OssURLLTE applies the LTE predicate on the "oss_url" field.
func OssURLLTE(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOssURL), v))
	})
}

// OssURLContains applies the Contains predicate on the "oss_url" field.
func OssURLContains(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOssURL), v))
	})
}

// OssURLHasPrefix applies the HasPrefix predicate on the "oss_url" field.
func OssURLHasPrefix(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOssURL), v))
	})
}

// OssURLHasSuffix applies the HasSuffix predicate on the "oss_url" field.
func OssURLHasSuffix(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOssURL), v))
	})
}

// OssURLEqualFold applies the EqualFold predicate on the "oss_url" field.
func OssURLEqualFold(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOssURL), v))
	})
}

// OssURLContainsFold applies the ContainsFold predicate on the "oss_url" field.
func OssURLContainsFold(v string) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOssURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.CaptureLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CaptureLog) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CaptureLog) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
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
func Not(p predicate.CaptureLog) predicate.CaptureLog {
	return predicate.CaptureLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
