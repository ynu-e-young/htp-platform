package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/biz"
)

// CronJob holds the schema definition for the CronJob entity.
type CronJob struct {
	ent.Schema
}

// Fields of the CronJob.
func (CronJob) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("machine_id", uuid.UUID{}),
		field.String("check_name"),
		field.String("cron_string"),
		field.JSON("coordinates", []*biz.CheckCoordinate{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the CronJob.
func (CronJob) Edges() []ent.Edge {
	return nil
}
