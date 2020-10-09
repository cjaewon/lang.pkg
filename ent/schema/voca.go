package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"lang.pkg/lib"
)

// Voca holds the schema definition for the Voca entity.
type Voca struct {
	ent.Schema
}

// Fields of the Voca.
func (Voca) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Default(lib.GenerateRandKey(5)),
		field.String("key"),
		field.String("value"),
		field.String("example"),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Voca.
func (Voca) Edges() []ent.Edge {
	return nil
}
