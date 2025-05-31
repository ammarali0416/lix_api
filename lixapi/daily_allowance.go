package lixapi

import (
	"fmt"
	"io"
	"net/http"
)

func getDailyAllowance(apiKey string) ([]byte, error) {
	endpoint := "https://api.lix-it.com/v1/account/allowances/daily"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
