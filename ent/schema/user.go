package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),  // Nombre único
		field.String("email").Unique(), // Email único
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil // Sin relaciones por ahora
}
