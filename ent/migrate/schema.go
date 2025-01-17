// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "book_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "public", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_books", Type: field.TypeUUID, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "books_users_books",
				Columns: []*schema.Column{BooksColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// VocasColumns holds the columns for the "vocas" table.
	VocasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "example", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "book_vocas", Type: field.TypeInt, Nullable: true},
	}
	// VocasTable holds the schema information for the "vocas" table.
	VocasTable = &schema.Table{
		Name:       "vocas",
		Columns:    VocasColumns,
		PrimaryKey: []*schema.Column{VocasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "vocas_books_vocas",
				Columns: []*schema.Column{VocasColumns[5]},

				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "voca_created_at",
				Unique:  false,
				Columns: []*schema.Column{VocasColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BooksTable,
		UsersTable,
		VocasTable,
	}
)

func init() {
	BooksTable.ForeignKeys[0].RefTable = UsersTable
	VocasTable.ForeignKeys[0].RefTable = BooksTable
}
