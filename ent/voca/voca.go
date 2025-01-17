// Code generated by entc, DO NOT EDIT.

package voca

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the voca type in the database.
	Label = "voca"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldExample holds the string denoting the example field in the database.
	FieldExample = "example"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// Table holds the table name of the voca in the database.
	Table = "vocas"
)

// Columns holds all SQL columns for voca fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
	FieldExample,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Voca type.
var ForeignKeys = []string{
	"book_vocas",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
