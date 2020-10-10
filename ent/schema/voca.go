package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Voca holds the schema definition for the Voca entity.
type Voca struct {
	ent.Schema
}

// Fields of the Voca.
func (Voca) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("key"),
		field.String("value"),
		field.String("example").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Voca.
func (Voca) Edges() []ent.Edge {
	return nil
}
