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
		// one to one 'Two Types'
		// user can have only one cart
		// User(ID-PK, ...)
		edge.To("card", Card.Type),

		// one to one 'Bidirectional'
		// user can have only one spouse
		// User(ID-PK, ..., Spouse-ID)
		edge.To("spouse", User.Type).Unique(),

		// one to many 'Two Type'
		// user can have many pets
		// User(ID-PK, ...)
		edge.To("pets", Pet.Type).Unique(),

		// many to many 'Two Type'
		// group can have many user
		// User(ID-PK, ...) & User-Group(user-ID-PK, group-ID-PK, ....) 'new table'
		edge.From("groups", Group.Type).Ref("users"),

		// many to many 'Same Type'
		// user can follow many users and can have many followers
		// User(ID-PK, ...) & Followers(User-ID+Following-ID-PK)
		edge.To("following", User.Type).From("followers"),
	}
}

/*
go run -mod=mod entgo.io/ent/cmd/ent new User
go generate ./ent
*/
