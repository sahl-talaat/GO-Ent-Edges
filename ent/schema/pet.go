package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		// one to many 'Two Type'
		// pet has one owner
		// Pet(ID-PK, ..., owner-ID-FK)
		edge.To("pets", Pet.Type).Unique(),
		edge.From("owner", User.Type).Ref("pets").Unique(),
	}
}
