package lixapi

func GetDailyAllowance(apiKey string) ([]byte, error) {
	return getDailyAllowance(apiKey)
}

func GetConnections(apiKey, viewerID string, count, start int) ([]byte, error) {
	return getConnections(apiKey, viewerID, count, start)
}

func GetPostsSearch(apiKey, url string, start int, viewerID, sequenceID string) ([]byte, error) {
	return getPostsSearch(apiKey, url, start, viewerID, sequenceID)
}
