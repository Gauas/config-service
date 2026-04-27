package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	secretKey  string
	httpClient *http.Client
}

type Options struct {
	BaseURL   string
	SecretKey string
	Timeout   time.Duration
}

func New(opts Options) *Client {
	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &Client{
		baseURL:   opts.BaseURL,
		secretKey: opts.SecretKey,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Get(service, environment string) (Config, error) {
	url := fmt.Sprintf(
		"%s/v1/config?service=%s&env=%s",
		c.baseURL,
		service,
		environment,
	)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("config-sdk: create request: %w", err)
	}

	req.Header.Set("secret_key", c.secretKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("config-sdk: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("config-sdk: unexpected status %d from config-service", resp.StatusCode)
	}

	var result Config
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("config-sdk: decode response: %w", err)
	}

	return result, nil
}
