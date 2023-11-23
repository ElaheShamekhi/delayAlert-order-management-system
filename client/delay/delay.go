package delay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const url = "https://run.mocky.io/v3/122c2796-5df4-461c-ab75-87c1192b17f7"

type Client struct {
	net *http.Client
}

func New() *Client {
	client := &Client{
		net: &http.Client{Timeout: time.Second * 30},
	}
	return client
}

func (c *Client) GetNewEstimatedDelay() (*time.Time, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.net.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		var e any
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.Join(err, fmt.Errorf("unknown status error: %s", res.Status))
		}
		return nil, errors.New("unknown status error")
	}
	var decodeTo any
	if err := json.NewDecoder(res.Body).Decode(decodeTo); err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to decode '%s' response", url))
	}
	return nil, nil
}
