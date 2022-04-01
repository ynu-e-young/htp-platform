// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/data/ent/cronjob"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CronJob is the model entity for the CronJob schema.
type CronJob struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// MachineID holds the value of the "machine_id" field.
	MachineID uuid.UUID `json:"machine_id,omitempty"`
	// CheckName holds the value of the "check_name" field.
	CheckName string `json:"check_name,omitempty"`
	// CronString holds the value of the "cron_string" field.
	CronString string `json:"cron_string,omitempty"`
	// Coordinates holds the value of the "coordinates" field.
	Coordinates []*biz.CheckCoordinate `json:"coordinates,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CronJob) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cronjob.FieldCoordinates:
			values[i] = new([]byte)
		case cronjob.FieldID:
			values[i] = new(sql.NullInt64)
		case cronjob.FieldCheckName, cronjob.FieldCronString:
			values[i] = new(sql.NullString)
		case cronjob.FieldCreatedAt, cronjob.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case cronjob.FieldMachineID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CronJob", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CronJob fields.
func (cj *CronJob) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cronjob.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cj.ID = int64(value.Int64)
		case cronjob.FieldMachineID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field machine_id", values[i])
			} else if value != nil {
				cj.MachineID = *value
			}
		case cronjob.FieldCheckName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field check_name", values[i])
			} else if value.Valid {
				cj.CheckName = value.String
			}
		case cronjob.FieldCronString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_string", values[i])
			} else if value.Valid {
				cj.CronString = value.String
			}
		case cronjob.FieldCoordinates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field coordinates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &cj.Coordinates); err != nil {
					return fmt.Errorf("unmarshal field coordinates: %w", err)
				}
			}
		case cronjob.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cj.CreatedAt = value.Time
			}
		case cronjob.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cj.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CronJob.
// Note that you need to call CronJob.Unwrap() before calling this method if this CronJob
// was returned from a transaction, and the transaction was committed or rolled back.
func (cj *CronJob) Update() *CronJobUpdateOne {
	return (&CronJobClient{config: cj.config}).UpdateOne(cj)
}

// Unwrap unwraps the CronJob entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cj *CronJob) Unwrap() *CronJob {
	tx, ok := cj.config.driver.(*txDriver)
	if !ok {
		panic("ent: CronJob is not a transactional entity")
	}
	cj.config.driver = tx.drv
	return cj
}

// String implements the fmt.Stringer.
func (cj *CronJob) String() string {
	var builder strings.Builder
	builder.WriteString("CronJob(")
	builder.WriteString(fmt.Sprintf("id=%v", cj.ID))
	builder.WriteString(", machine_id=")
	builder.WriteString(fmt.Sprintf("%v", cj.MachineID))
	builder.WriteString(", check_name=")
	builder.WriteString(cj.CheckName)
	builder.WriteString(", cron_string=")
	builder.WriteString(cj.CronString)
	builder.WriteString(", coordinates=")
	builder.WriteString(fmt.Sprintf("%v", cj.Coordinates))
	builder.WriteString(", created_at=")
	builder.WriteString(cj.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cj.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CronJobs is a parsable slice of CronJob.
type CronJobs []*CronJob

func (cj CronJobs) config(cfg config) {
	for _i := range cj {
		cj[_i].config = cfg
	}
}
