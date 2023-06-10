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

type StoreKey struct{}

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
	databaseClient := Open(databaseConnectionString)

	// Handling Redis Sessions
	http.HandleFunc("/redis", func(resWriter http.ResponseWriter, req *http.Request) {

		redisStore := healthyshopper.NewRedisStore("localhost:6379", "", 0, resWriter, *req)

		ctx := context.WithValue(req.Context(), StoreKey{}, redisStore)

		redisStore.SetCookie("test", "test1212")
		redisStore.SetCookie("testNew", "test5656")

		println(redisStore.Get("test"))
		println(redisStore.Get("testNew"))

		// Your code. For example:

		redisContext := context.WithValue(ctx, "redis", redisStore)

		if err := databaseClient.Schema.Create(redisContext); err != nil {
			log.Fatal("opening ent client", err)
		}

	})

	graphqlServer := handler.NewDefaultServer(healthyshopper.NewSchema(databaseClient))

	http.Handle("/", graphqlServer)

	http.Handle("/graphql", playground.Handler("GraphQL Playground", "/"))

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http server terminated", err)
	}

	// log.Printf("connect to http://localhost%s/graphql for GraphQL playground", defaultPort)
	// err := http.ListenAndServe(defaultPort, nil)
	// if err != nil {
	// 	log.Fatal("http server terminated", err)
	// }

	// Configure the server and start listening on :8080.

	// // http.Handle("/graphql",
	// // 	playground.Handler("GraphQL Playground", "/graphql"),
	// // )
	// // // http.Handle("/graphql", server)

	// log.Println("listening on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("http server terminated", err)
	// }
}
