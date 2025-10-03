package tmdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiKey string
	http   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) TMDBRequest(method, url string, body interface{}) (map[string]interface{}, error) {
	var reqBodyReader *bytes.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBodyReader = bytes.NewReader(b)
	} else {
		reqBodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, url, reqBodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDb returned status %d", res.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
