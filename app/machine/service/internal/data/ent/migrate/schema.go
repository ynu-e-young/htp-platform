// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CaptureLogsColumns holds the columns for the "capture_logs" table.
	CaptureLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "machine_id", Type: field.TypeUUID},
		{Name: "pixels", Type: field.TypeInt64},
		{Name: "area", Type: field.TypeFloat64},
		{Name: "src_name", Type: field.TypeString},
		{Name: "proc_name", Type: field.TypeString},
		{Name: "src_oss_url", Type: field.TypeString},
		{Name: "proc_oss_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// CaptureLogsTable holds the schema information for the "capture_logs" table.
	CaptureLogsTable = &schema.Table{
		Name:       "capture_logs",
		Columns:    CaptureLogsColumns,
		PrimaryKey: []*schema.Column{CaptureLogsColumns[0]},
	}
	// CronJobsColumns holds the columns for the "cron_jobs" table.
	CronJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "machine_id", Type: field.TypeUUID},
		{Name: "check_name", Type: field.TypeString},
		{Name: "cron_string", Type: field.TypeString},
		{Name: "coordinates", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// CronJobsTable holds the schema information for the "cron_jobs" table.
	CronJobsTable = &schema.Table{
		Name:       "cron_jobs",
		Columns:    CronJobsColumns,
		PrimaryKey: []*schema.Column{CronJobsColumns[0]},
	}
	// MachinesColumns holds the columns for the "machines" table.
	MachinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// MachinesTable holds the schema information for the "machines" table.
	MachinesTable = &schema.Table{
		Name:       "machines",
		Columns:    MachinesColumns,
		PrimaryKey: []*schema.Column{MachinesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CaptureLogsTable,
		CronJobsTable,
		MachinesTable,
	}
)

func init() {
}
