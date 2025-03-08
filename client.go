package yahoofinanceapi

import (
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	client  *http.Client
	cookies []*http.Cookie
	crumb   string
}

var instance *Client
var once sync.Once

func GetClient() *Client {
	once.Do(func() {
		instance = &Client{client: &http.Client{}, cookies: []*http.Cookie{}, crumb: ""}
	})
	return instance
}

func (c *Client) Get(url string, params url.Values) (*http.Response, error) {
	c.getCrumb()
	return c.get(url, params)
}

func (c *Client) get(url string, params url.Values) (*http.Response, error) {
	if c.crumb != "" {
		params.Add("crumb", c.crumb)
	}
	url = fmt.Sprintf("%s?%s", url, params.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Failed to create request", "err", err)
		return nil, err
	}

	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", USER_AGENTS[rand.Intn(len(USER_AGENTS))])
	resp, err := c.client.Do(req)
	if err != nil {
		slog.Error("Failed to get data from Yahoo Finance API", "err", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) getCookie() {
	if len(c.cookies) > 0 {
		return
	}

	endpoint := "https://fc.yahoo.com"
	resp, err := c.get(endpoint, url.Values{})
	if err != nil {
		slog.Error("Failed to get cookie", "err", err)
		return
	}

	c.cookies = resp.Cookies()
}

func (c *Client) getCrumb() {
	slog.Info("crumb", "crumb", c.crumb)
	if c.crumb != "" {
		return
	}

	c.getCookie()
	endpoint := fmt.Sprintf("%s/v1/test/getcrumb", BASE_URL)
	resp, err := c.get(endpoint, url.Values{})
	if err != nil {
		slog.Error("Failed to get crumb", "err", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body:", "err", err)
		return
	}

	c.crumb = string(body)
}
