package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		// many to many 'Two Type'
		// group can have many user
		// Group(ID-PK, ...) & User-Group(group-ID-PK, user-ID-PK, ....) 'new table'
		edge.To("users", User.Type),
	}
}
