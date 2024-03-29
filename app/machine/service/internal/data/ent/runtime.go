// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent/capturelog"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent/cronjob"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent/machine"
	"github.com/ynu-e-young/htp-platform/app/machine/service/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	capturelogFields := schema.CaptureLog{}.Fields()
	_ = capturelogFields
	// capturelogDescCreatedAt is the schema descriptor for created_at field.
	capturelogDescCreatedAt := capturelogFields[8].Descriptor()
	// capturelog.DefaultCreatedAt holds the default value on creation for the created_at field.
	capturelog.DefaultCreatedAt = capturelogDescCreatedAt.Default.(func() time.Time)
	// capturelogDescUpdatedAt is the schema descriptor for updated_at field.
	capturelogDescUpdatedAt := capturelogFields[9].Descriptor()
	// capturelog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	capturelog.DefaultUpdatedAt = capturelogDescUpdatedAt.Default.(func() time.Time)
	cronjobFields := schema.CronJob{}.Fields()
	_ = cronjobFields
	// cronjobDescCreatedAt is the schema descriptor for created_at field.
	cronjobDescCreatedAt := cronjobFields[5].Descriptor()
	// cronjob.DefaultCreatedAt holds the default value on creation for the created_at field.
	cronjob.DefaultCreatedAt = cronjobDescCreatedAt.Default.(func() time.Time)
	// cronjobDescUpdatedAt is the schema descriptor for updated_at field.
	cronjobDescUpdatedAt := cronjobFields[6].Descriptor()
	// cronjob.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cronjob.DefaultUpdatedAt = cronjobDescUpdatedAt.Default.(func() time.Time)
	machineFields := schema.Machine{}.Fields()
	_ = machineFields
	// machineDescCreatedAt is the schema descriptor for created_at field.
	machineDescCreatedAt := machineFields[3].Descriptor()
	// machine.DefaultCreatedAt holds the default value on creation for the created_at field.
	machine.DefaultCreatedAt = machineDescCreatedAt.Default.(func() time.Time)
	// machineDescUpdatedAt is the schema descriptor for updated_at field.
	machineDescUpdatedAt := machineFields[4].Descriptor()
	// machine.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	machine.DefaultUpdatedAt = machineDescUpdatedAt.Default.(func() time.Time)
}
