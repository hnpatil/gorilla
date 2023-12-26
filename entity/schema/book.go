package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Book struct {
	ent.Schema
}

func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.String("discription"),
	}
}

func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("books").
			Unique().
			Required(),
	}
}
