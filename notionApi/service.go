package notionApi

import "net/http"

type DatabaseService interface {
	Get(databaseId string) *http.Response
	Query(databaseId string, body interface{}) *http.Response
}

type databaseClient struct {
	apiClient *Client
}
