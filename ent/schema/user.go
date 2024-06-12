package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.Int("age").Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// one to many
		// user can have many pets
		edge.To("pets", Pet.Type),

		// many to many
		// user can have many groups
		edge.From("groups", Group.Type).Ref("users"),

		// one to one
		// user can have only one cart
		edge.To("card", Card.Type),
	}
}

/*
go run -mod=mod entgo.io/ent/cmd/ent new User
go generate ./ent
*/
