package lixapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func getPostsSearch(apiKey, urlStr string, start int, viewerID, sequenceID string) ([]byte, error) {
	endpoint := "https://api.lix-it.com/v1/li/linkedin/search/posts"
	params := url.Values{}
	params.Set("url", urlStr)
	params.Set("start", fmt.Sprintf("%d", start))
	if viewerID != "" {
		params.Set("viewer_id", viewerID)
	}
	if sequenceID != "" {
		params.Set("sequence_id", sequenceID)
	}
	fullURL := endpoint + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
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
