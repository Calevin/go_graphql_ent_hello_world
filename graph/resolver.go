package graph

import (
	"go_graphql_ent_hello_world/ent" // Se importa el paquete Ent
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client // Cliente Ent
}
