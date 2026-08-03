package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/learnda1ly/homeschool_tracker/internal/database"
)

func main() {
	db, err := sql.Open("pgx", "postgres://squinlan@localhost:5432/go_homeschool?sslmode=disable")
	if err != nil {
		log.Fatal("open db:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("ping db:", err)
	}

	ctx := context.Background()
	queries := database.New(db)

	family, err := queries.CreateFamily(ctx, "Test Family")
	if err != nil {
		log.Fatal("create family:", err)
	}
	fmt.Printf("created family: id=%d name=%s\n", family.ID, family.FamilyName)

	learner, err := queries.CreateUser(ctx, database.CreateUserParams{
		FamilyID:  family.ID,
		FirstName: "Test",
		LastName:  "User",
		Grade:     sql.NullString{String: "3rd", Valid: true},
	})
	if err != nil {
		log.Fatal("create user:", err)
	}
	fmt.Printf("created learner: id=%d family_id=%d name=%s %s grade=%s\n",
		learner.ID, learner.FamilyID, learner.FirstName, learner.LastName, learner.Grade.String)
}
