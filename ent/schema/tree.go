package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tree holds the schema definition for the Tree entity.
type Tree struct {
	ent.Schema
}

// Fields of the Tree.
func (Tree) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value"),
	}
}

// Edges of the Tree.
func (Tree) Edges() []ent.Edge {
	return []ent.Edge{
		// one to many 'same Type'
		// each node can have many children bu children can have only one parent
		// Node(ID-PK, ..., parent-ID-K)
		edge.To("children", Tree.Type).From("parent").Unique(),
	}
}
