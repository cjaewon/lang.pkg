package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"), // auto increment primary key not uuid
		field.String("book_id"),
		field.String("title"),
		field.String("description"),
		field.Bool("public"),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vocas", Voca.Type),
	}
}
