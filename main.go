package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"go_graphql_ent_hello_world/ent"   // Tu paquete Ent
	"go_graphql_ent_hello_world/graph" // Tu paquete graph

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib" // Importa el driver de pgx para database/sql
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// --- Configurar la Base de Datos (Ent) ---

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

	// --- Configurar el Servidor GraphQL (gqlgen) ---

	// Crear el handler GraphQL
	// Pasamos el cliente Ent a los resolvers
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: client}}))

	// Crear el handler del Playground (UI para probar GraphQL)
	playgroundHandler := playground.Handler("GraphQL playground", "/graphql")

	// --- Configurar el Router HTTP (Chi) ---

	router := chi.NewRouter()

	// Rutas
	router.Handle("/", playgroundHandler) // UI en la raíz
	router.Handle("/graphql", srv)        // Endpoint GraphQL

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
