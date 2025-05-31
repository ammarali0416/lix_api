package lixapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getConnections(apiKey, viewerID string, count, start int) ([]byte, error) {
	endpoint := "https://api.lix-it.com/v1/connections"
	params := url.Values{}
	params.Set("viewer_id", viewerID)
	params.Set("count", fmt.Sprintf("%d", count))
	params.Set("start", fmt.Sprintf("%d", start))
	urlStr := endpoint + "?" + params.Encode()

	req, err := http.NewRequest("GET", urlStr, nil)
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
