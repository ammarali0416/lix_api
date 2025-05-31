package supabase

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/supabase-community/supabase-go"
)

func GetSupabaseClient() *supabase.Client {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	client, err := supabase.NewClient(supabaseUrl, supabaseKey,
		&supabase.ClientOptions{
			Schema: "public",
		},
	)
	if err != nil {
		log.Fatalf("Failed to create supabase client: %v", err)
	}
	return client
}

func GetPGClient() *sql.DB {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	database := os.Getenv("PG_DATABASE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to create postgres client: %v", err)
	}
	return db
}
