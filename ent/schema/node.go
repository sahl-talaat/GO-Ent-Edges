package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value"),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		// one to one 'Same Type'
		// each node have next and prev from the same node type
		// Node(ID-PK, ...., prev-id-FK'refere to p-k in same table' nullable)
		edge.To("next", Node.Type).Unique().From("prev").Unique(),
	}
}
