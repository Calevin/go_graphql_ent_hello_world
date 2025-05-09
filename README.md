# go_graphql_ent_hello_world
Example GraphQL + Ent (ORM)

## Añadir dependencias
```sh
# Ent (ORM)
go get entgo.io/ent
# gqlgen (GraphQL codegen)
go get github.com/99designs/gqlgen
# driver de PostgreSQL
github.com/jackc/pgx/v5
# router HTTP
github.com/go-chi/chi/v5
```

## Instalar generadores
```sh
# Ent codegen
go install entgo.io/ent/cmd/ent@latest
# gqlgen
go install github.com/99designs/gqlgen@latest
```

## Definir el esquema de datos con Ent
```sh
ent new User --target ./ent/schema
```

## Crear ent/generate.go con el siguiente contenido:
```
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```

Le dice a Go que ejecute el generador de Ent (entgo.io/ent/cmd/ent generate) apuntando al directorio ./schema (que contiene los archivos de esquema como user.go). 
Se usa -mod=mod para asegurar que la resolución de dependencias sea reproducible.

## Ejecutar generate de Ent
```
go generate ./ent
```
