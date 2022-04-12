package notionApi

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	apiURL        = "https://api.notion.com"
	apiVersion    = "v1"
	notionVersion = "2021-08-16"
)

type Client struct {
	httpClient    *http.Client
	baseUrl       *url.URL
	apiVersion    string
	notionVersion string
	Token         string

	Database DatabaseService
}

func NewClient(token string) (*Client, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}
	if token == "" {
		log.Fatal("Empty Token")
		return nil, err
	}
	c := &Client{
		httpClient:    http.DefaultClient,
		baseUrl:       u,
		apiVersion:    apiVersion,
		notionVersion: notionVersion,
		Token:         token,
	}

	c.Database = &databaseClient{apiClient: c}

	return c, nil
}

func (c Client) Request(url, method string, requestBody string) (*http.Response, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", c.apiVersion, url))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != "" {
		//body, err := json.Marshal(requestBody)
		//if err != nil {
		//	return nil, err
		//}
		buf = bytes.NewBuffer([]byte(requestBody))
	}

	uString := u.String()

	req, err := http.NewRequest(method, uString, buf)

	if err != nil {
		return nil, fmt.Errorf("got error %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Add("Notion-Version", c.notionVersion)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("got error %s", err.Error())
	}
	//defer response.Body.Close()

	return response, nil
}
