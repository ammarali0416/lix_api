package supabase

import (
	"database/sql"
)

func TruncateTable(db *sql.DB, tableName string) error {
	_, err := db.Exec("TRUNCATE TABLE public." + tableName)
	return err
}
