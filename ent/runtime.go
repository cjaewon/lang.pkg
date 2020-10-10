// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"lang.pkg/ent/book"
	"lang.pkg/ent/schema"
	"lang.pkg/ent/user"
	"lang.pkg/ent/voca"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescCreatedAt is the schema descriptor for created_at field.
	bookDescCreatedAt := bookFields[5].Descriptor()
	// book.DefaultCreatedAt holds the default value on creation for the created_at field.
	book.DefaultCreatedAt = bookDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	vocaFields := schema.Voca{}.Fields()
	_ = vocaFields
	// vocaDescCreatedAt is the schema descriptor for created_at field.
	vocaDescCreatedAt := vocaFields[4].Descriptor()
	// voca.DefaultCreatedAt holds the default value on creation for the created_at field.
	voca.DefaultCreatedAt = vocaDescCreatedAt.Default.(func() time.Time)
	// vocaDescID is the schema descriptor for id field.
	vocaDescID := vocaFields[0].Descriptor()
	// voca.DefaultID holds the default value on creation for the id field.
	voca.DefaultID = vocaDescID.Default.(func() uuid.UUID)
}
