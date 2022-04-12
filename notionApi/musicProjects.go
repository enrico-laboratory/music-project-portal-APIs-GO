package notionApi

import (
	"encoding/json"
	"net/http"
)

func (db *databaseClient) Get(databaseId string) *http.Response {
	url := "databases/" + databaseId
	m := http.MethodGet
	resp, _ := db.apiClient.Request(url, m, "")
	return resp
}

func (db *databaseClient) Query(databaseId string, body interface{}) *http.Response {
	url := "databases/" + databaseId + "/query"
	m := http.MethodPost

	if body != nil {
		sb, _ := json.Marshal(body)
		resp, _ := db.apiClient.Request(url, m, string(sb))
		return resp
	}

	resp, _ := db.apiClient.Request(url, m, "")
	return resp

}
