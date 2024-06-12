package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Unique(),
		field.String("password").MaxLen(6).MinLen(6).Immutable(),
		field.Float("cach"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		// one to one
		// card can have only one user
		edge.From("owner", User.Type).Ref("card").Unique().Required(),
		/* .SetOwnerid()'user */
	}
}
