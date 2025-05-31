package supabase

import (
	"github.com/supabase-community/supabase-go"
)

func LoadStagingDataSupabase(client *supabase.Client, data interface{}, metadata interface{}) error {
	insertData := map[string]interface{}{
		"data":     data,
		"metadata": metadata,
	}
	_, _, err := client.From("lix_staging_json").Insert(insertData, false, "", "", "").Execute()
	return err
}
