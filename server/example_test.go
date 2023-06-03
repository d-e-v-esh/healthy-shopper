package healthyshopper

import (
	"context"
	"fmt"
	"healthyshopper/ent"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_User() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	task1, err := client.User.Create().SetUsername("dvsh").SetEmailAddress("dev@dev.com").SetPassword("123456").SetFirstName("Devesh").SetLastName("K").Save(ctx)

	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", task1)

	items, err := client.User.Query().All(ctx)

	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}

	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Username)
	}

	// Output:
	// 1: "dvsh"
}
