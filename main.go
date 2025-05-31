package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"lix_api/lixapi"
	"lix_api/supabase"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("LIX_API_KEY")
	if apiKey == "" {
		log.Fatalf("LIX_API_KEY not set in .env file")
	}

	if len(os.Args) < 2 {
		fmt.Println("expected 'daily-allowance' or 'connections' subcommands")
		os.Exit(1)
	}

	subcmd := os.Args[1]

	switch subcmd {
	case "daily-allowance":
		dailyAllowanceCmd := flag.NewFlagSet("daily-allowance", flag.ExitOnError)
		outfile := dailyAllowanceCmd.String("outfile", "", "output file (optional)")
		dailyAllowanceCmd.Parse(os.Args[2:])

		resp, err := lixapi.GetDailyAllowance(apiKey)
		if err != nil {
			log.Fatalf("API error: %v", err)
		}
		if *outfile != "" {
			err = os.WriteFile(*outfile, resp, 0644)
			if err != nil {
				log.Fatalf("Failed to write file: %v", err)
			}
			fmt.Printf("Response saved to %s\n", *outfile)
		} else {
			fmt.Println(string(resp))
		}
	case "connections":
		connectionsCmd := flag.NewFlagSet("connections", flag.ExitOnError)
		viewerID := connectionsCmd.String("viewerid", "", "viewer_id (required)")
		count := connectionsCmd.Int("count", 1000, "count (default 1000)")
		start := connectionsCmd.Int("start", 0, "start (default 0)")
		outfile := connectionsCmd.String("outfile", "", "output file (optional)")
		connectionsCmd.Parse(os.Args[2:])

		if *viewerID == "" {
			log.Fatalf("viewerid is required")
		}
		resp, err := lixapi.GetConnections(apiKey, *viewerID, *count, *start)
		if err != nil {
			log.Fatalf("API error: %v", err)
		}
		if *outfile != "" {
			err = os.WriteFile(*outfile, resp, 0644)
			if err != nil {
				log.Fatalf("Failed to write file: %v", err)
			}
			fmt.Printf("Response saved to %s\n", *outfile)
		} else {
			fmt.Println(string(resp))
		}
	case "posts-search":
		postsSearchCmd := flag.NewFlagSet("posts-search", flag.ExitOnError)
		urlFlag := postsSearchCmd.String("url", "", "url-encoded LinkedIn search URL (required)")
		start := postsSearchCmd.Int("start", 0, "start offset (default 0)")
		viewerID := postsSearchCmd.String("viewerid", "", "viewer_id (optional)")
		sequenceID := postsSearchCmd.String("sequenceid", "", "sequence_id (optional)")
		outfile := postsSearchCmd.String("outfile", "", "output file (optional)")
		postsSearchCmd.Parse(os.Args[2:])

		if *urlFlag == "" {
			log.Fatalf("url is required")
		}
		resp, err := lixapi.GetPostsSearch(apiKey, *urlFlag, *start, *viewerID, *sequenceID)
		if err != nil {
			log.Fatalf("API error: %v", err)
		}
		if *outfile != "" {
			err = os.WriteFile(*outfile, resp, 0644)
			if err != nil {
				log.Fatalf("Failed to write file: %v", err)
			}
			fmt.Printf("Response saved to %s\n", *outfile)
		} else {
			fmt.Println(string(resp))
		}
	case "test-db":
		fmt.Println("Testing database connections...")
		// Test Postgres
		pg := supabase.GetPGClient()
		var one int
		err := pg.QueryRow("SELECT 1").Scan(&one)
		if err != nil {
			fmt.Printf("Postgres connection failed: %v\n", err)
		} else {
			fmt.Println("Postgres connection successful, SELECT 1 =>", one)
		}
		pg.Close()

		// Test Supabase
		sb := supabase.GetSupabaseClient()
		// Supabase client does not expose direct SQL, so just check client is not nil
		if sb == nil {
			fmt.Println("Supabase client creation failed")
		} else {
			fmt.Println("Supabase client creation successful")
		}
		return
	default:
		fmt.Println("expected 'daily-allowance' or 'connections' subcommands")
		os.Exit(1)
	}
}
