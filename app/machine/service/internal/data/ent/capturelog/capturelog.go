// Code generated by ent, DO NOT EDIT.

package capturelog

import (
	"time"
)

const (
	// Label holds the string label denoting the capturelog type in the database.
	Label = "capture_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMachineID holds the string denoting the machine_id field in the database.
	FieldMachineID = "machine_id"
	// FieldPixels holds the string denoting the pixels field in the database.
	FieldPixels = "pixels"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// FieldSrcName holds the string denoting the src_name field in the database.
	FieldSrcName = "src_name"
	// FieldProcName holds the string denoting the proc_name field in the database.
	FieldProcName = "proc_name"
	// FieldSrcOssURL holds the string denoting the src_oss_url field in the database.
	FieldSrcOssURL = "src_oss_url"
	// FieldProcOssURL holds the string denoting the proc_oss_url field in the database.
	FieldProcOssURL = "proc_oss_url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the capturelog in the database.
	Table = "capture_logs"
)

// Columns holds all SQL columns for capturelog fields.
var Columns = []string{
	FieldID,
	FieldMachineID,
	FieldPixels,
	FieldArea,
	FieldSrcName,
	FieldProcName,
	FieldSrcOssURL,
	FieldProcOssURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
