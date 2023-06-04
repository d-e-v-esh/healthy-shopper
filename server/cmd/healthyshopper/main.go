package main

import (
	"context"
	"database/sql"
	"healthyshopper"
	"healthyshopper/ent"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const defaultPort = "8080"
const databaseConnectionString = "postgres://postgres:postgres@localhost:5432/HealthyShopper?sslmode=disable"

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open(databaseConnectionString)

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8080.
	srv := handler.NewDefaultServer(healthyshopper.NewSchema(client))
	http.Handle("/",
		playground.Handler("User", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
