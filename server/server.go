package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/d-e-v-esh/healthy-shopper/graph"
	"github.com/jackc/pgx/v5"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const defaultPort = "8080"
const databaseConnectionString = "postgres://postgres:postgres@localhost:5432/HealthyShopper?sslmode=disable"

func runMigrations() {

	absPath, pathErr := filepath.Abs("../server/migrations")

	fmt.Println(absPath)
	if pathErr != nil {
		log.Fatal("PathError: ", pathErr)
	}

	m, err := migrate.New(
		"file://migrations",
		databaseConnectionString)

	if err != nil {
		log.Fatal("MigrateError: ", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	// -------------------------------------------

	// db, dbErr := sql.Open("postgres", databaseConnectionString)
	// if dbErr != nil {
	// 	log.Fatal("DBError: ", dbErr)
	// }

	// driver, driverErr := postgres.WithInstance(db, &postgres.Config{})

	// fmt.Println(driver)

	// if driverErr != nil {
	// 	log.Fatal("DriverError: ", driverErr)
	// }

	// absPath, pathErr := filepath.Abs("../server/migrations")

	// if pathErr != nil {
	// 	log.Fatal("PathError: ", pathErr)
	// }

	// m, migrateErr := migrate.NewWithDatabaseInstance(
	// 	absPath,
	// 	"postgres", driver)

	// if migrateErr != nil {
	// 	log.Fatal("MigrateError: ", migrateErr)
	// }

	// m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}

func main() {
	conn, err := pgx.Connect(context.Background(), databaseConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	runMigrations()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
