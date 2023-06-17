package service

import (
	"context"
	"database/sql"
	"log"
	"userManagementSystem/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	// privacy "entgo.io/ent/examples/privacy"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func OpenDB() *ent.Client {
	// db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=ent-practise password=admin sslmode=disable")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to postgres: %v", err)
	// }
	// defer db.Close()

	client := Open("postgresql://postgres:admin@localhost:5432/ent-practise")

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// client := ent.NewClient(ent.Driver(dialect.Postgres), ent.DB(db))

	// client := ent.NewClient(dialect.Postgres, ent.(db))
	// if err != nil {
	// 	log.Fatalf("failed to create Ent client: %v", err)
	// }
	// defer client.Close()

	return client
}

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
