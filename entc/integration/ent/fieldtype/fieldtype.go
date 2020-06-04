// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package fieldtype

import (
	"fmt"
)

const (
	// Label holds the string label denoting the fieldtype type in the database.
	Label = "field_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInt holds the string denoting the int field in the database.
	FieldInt = "int"
	// FieldInt8 holds the string denoting the int8 field in the database.
	FieldInt8 = "int8"
	// FieldInt16 holds the string denoting the int16 field in the database.
	FieldInt16 = "int16"
	// FieldInt32 holds the string denoting the int32 field in the database.
	FieldInt32 = "int32"
	// FieldInt64 holds the string denoting the int64 field in the database.
	FieldInt64 = "int64"
	// FieldOptionalInt holds the string denoting the optional_int field in the database.
	FieldOptionalInt = "optional_int"
	// FieldOptionalInt8 holds the string denoting the optional_int8 field in the database.
	FieldOptionalInt8 = "optional_int8"
	// FieldOptionalInt16 holds the string denoting the optional_int16 field in the database.
	FieldOptionalInt16 = "optional_int16"
	// FieldOptionalInt32 holds the string denoting the optional_int32 field in the database.
	FieldOptionalInt32 = "optional_int32"
	// FieldOptionalInt64 holds the string denoting the optional_int64 field in the database.
	FieldOptionalInt64 = "optional_int64"
	// FieldNillableInt holds the string denoting the nillable_int field in the database.
	FieldNillableInt = "nillable_int"
	// FieldNillableInt8 holds the string denoting the nillable_int8 field in the database.
	FieldNillableInt8 = "nillable_int8"
	// FieldNillableInt16 holds the string denoting the nillable_int16 field in the database.
	FieldNillableInt16 = "nillable_int16"
	// FieldNillableInt32 holds the string denoting the nillable_int32 field in the database.
	FieldNillableInt32 = "nillable_int32"
	// FieldNillableInt64 holds the string denoting the nillable_int64 field in the database.
	FieldNillableInt64 = "nillable_int64"
	// FieldValidateOptionalInt32 holds the string denoting the validate_optional_int32 field in the database.
	FieldValidateOptionalInt32 = "validate_optional_int32"
	// FieldOptionalUint holds the string denoting the optional_uint field in the database.
	FieldOptionalUint = "optional_uint"
	// FieldOptionalUint8 holds the string denoting the optional_uint8 field in the database.
	FieldOptionalUint8 = "optional_uint8"
	// FieldOptionalUint16 holds the string denoting the optional_uint16 field in the database.
	FieldOptionalUint16 = "optional_uint16"
	// FieldOptionalUint32 holds the string denoting the optional_uint32 field in the database.
	FieldOptionalUint32 = "optional_uint32"
	// FieldOptionalUint64 holds the string denoting the optional_uint64 field in the database.
	FieldOptionalUint64 = "optional_uint64"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldOptionalFloat holds the string denoting the optional_float field in the database.
	FieldOptionalFloat = "optional_float"
	// FieldOptionalFloat32 holds the string denoting the optional_float32 field in the database.
	FieldOptionalFloat32 = "optional_float32"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"
	// FieldDecimal holds the string denoting the decimal field in the database.
	FieldDecimal = "decimal"
	// FieldDir holds the string denoting the dir field in the database.
	FieldDir = "dir"
	// FieldNdir holds the string denoting the ndir field in the database.
	FieldNdir = "ndir"
	// FieldStr holds the string denoting the str field in the database.
	FieldStr = "str"
	// FieldNullStr holds the string denoting the null_str field in the database.
	FieldNullStr = "null_str"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldNullLink holds the string denoting the null_link field in the database.
	FieldNullLink = "null_link"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldNullActive holds the string denoting the null_active field in the database.
	FieldNullActive = "null_active"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"

	// Table holds the table name of the fieldtype in the database.
	Table = "field_types"
)

// Columns holds all SQL columns for fieldtype fields.
var Columns = []string{
	FieldID,
	FieldInt,
	FieldInt8,
	FieldInt16,
	FieldInt32,
	FieldInt64,
	FieldOptionalInt,
	FieldOptionalInt8,
	FieldOptionalInt16,
	FieldOptionalInt32,
	FieldOptionalInt64,
	FieldNillableInt,
	FieldNillableInt8,
	FieldNillableInt16,
	FieldNillableInt32,
	FieldNillableInt64,
	FieldValidateOptionalInt32,
	FieldOptionalUint,
	FieldOptionalUint8,
	FieldOptionalUint16,
	FieldOptionalUint32,
	FieldOptionalUint64,
	FieldState,
	FieldOptionalFloat,
	FieldOptionalFloat32,
	FieldDatetime,
	FieldDecimal,
	FieldDir,
	FieldNdir,
	FieldStr,
	FieldNullStr,
	FieldLink,
	FieldNullLink,
	FieldActive,
	FieldNullActive,
	FieldDeleted,
	FieldDeletedAt,
	FieldIP,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FieldType type.
var ForeignKeys = []string{
	"file_field",
}

var (
	// ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	ValidateOptionalInt32Validator func(int32) error
)

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateOn  State = "on"
	StateOff State = "off"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateOff:
		return nil
	default:
		return fmt.Errorf("fieldtype: invalid enum value for state field: %q", s)
	}
}
