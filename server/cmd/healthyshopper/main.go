package main

import (
	"context"
	"database/sql"
	"errors"
	"healthyshopper"
	"healthyshopper/ent"
	"log"
	"net/http"

	resErr "healthyshopper/pkg/response_errors"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

	http.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {

		redisStore := healthyshopper.NewRedisStore("localhost:6379", "", 0, resWriter, *req)

		ctx := context.WithValue(context.Background(), "redisStore", redisStore)

		if err := databaseClient.Schema.Create(ctx); err != nil {
			log.Fatal("opening ent client", err)
		}

		graphqlServer := handler.NewDefaultServer(healthyshopper.NewSchema(databaseClient))

		graphqlServer.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			err := graphql.DefaultErrorPresenter(ctx, e)

			var myErr *resErr.ResponseError
			if errors.As(e, &myErr) {
				err.Message = myErr.MainError.Error()
				err.Extensions = map[string]interface{}{
					"field":   myErr.Field,
					"message": myErr.Message,
				}
			}

			return err

		})

		graphqlServer.ServeHTTP(resWriter, req.WithContext(ctx))

	})

	http.Handle("/graphql", playground.Handler("GraphQL Playground", "/"))

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http server terminated", err)
	}

}
