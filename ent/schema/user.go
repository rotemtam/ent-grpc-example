package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("username").NotEmpty().Unique().Annotations(entproto.Field(2)),
		field.Text("first_name").NotEmpty().Annotations(entproto.Field(3)),
		field.Text("last_name").NotEmpty().Annotations(entproto.Field(4)),
		field.String("email").NotEmpty().Unique().Annotations(entproto.Field(5)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin_of", Group.Type).Annotations(entproto.Field(6)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
