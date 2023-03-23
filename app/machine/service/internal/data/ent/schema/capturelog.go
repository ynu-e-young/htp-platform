package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CaptureLog holds the schema definition for the CaptureLog entity.
type CaptureLog struct {
	ent.Schema
}

// Fields of the CaptureLog.
func (CaptureLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("machine_id", uuid.UUID{}),
		field.Int64("pixels"),
		field.Float("area"),
		field.String("src_name"),
		field.String("proc_name"),
		field.String("src_oss_url"),
		field.String("proc_oss_url"),
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

// Edges of the CaptureLog.
func (CaptureLog) Edges() []ent.Edge {
	return nil
}
