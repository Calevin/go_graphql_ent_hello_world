package main

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib" // Importa el driver de pgx para database/sql
	"go_graphql_ent_hello_world/ent"
	"log"
)

func main() {
	// Cadena de conexión a la DB Docker
	// "postgres://user:password@host:port/dbname?sslmode=disable"
	databaseURL := "postgres://gql_ent_db:gql_ent_db@localhost:5436/gql_ent_db?sslmode=disable"

	// Abrir conexión con database/sql usando el driver pgx
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	// Crear un driver de Ent a partir de la conexión sql
	drv := entsql.OpenDB(dialect.Postgres, db)
	// Crear el cliente Ent
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	// Ejecutar migraciones
	// !!! IMPORTANTE: Esto es para desarrollo rápido.
	// En producción, usa la herramienta de migración de Ent
	// (go run entgo.io/ent/cmd/ent migrate create --env dsn=...) o Goose.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Database schema created/updated successfully.")

	u, _ := client.User.Create().
		SetName("Juan").
		SetEmail("juan@test.com").
		Save(ctx)

	log.Printf("User save name %v\n", u.Name)
}
