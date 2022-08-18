package rest

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
}

type Config struct {
	Host               string
	Username           string
	Password           string
	Timeout            time.Duration
	MaxRetries         int
	RetryInterval      time.Duration
	BearerToken        string
	ContentType        string
	AcceptContextTypes string
}

type Client struct {
	base   *url.URL
	config *ClientConfig
	Client *gorequest.SuperAgent
}

type ClientConfig struct {
	Username           string
	Password           string
	AcceptContextTypes string
	ContentType        string
	BearerToken        string
}

func NewClient(baseURL *url.URL, config *ClientConfig, client *gorequest.SuperAgent) (*Client, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}

	base.RawQuery = ""
	base.Fragment = ""

	return &Client{
		base:   &base,
		config: config,
		Client: client,
	}, nil
}

func NewClientForConfig(config *Config) (*Client, error) {
	hostUrl, err := url.Parse(config.Host)
	if err != nil {
		return nil, err
	}

	client := gorequest.New().Timeout(config.Timeout).Retry(config.MaxRetries, config.RetryInterval, http.StatusInternalServerError)
	client.DoNotClearSuperAgent = true

	clientConfig := &ClientConfig{
		Username:           config.Username,
		Password:           config.Password,
		ContentType:        config.ContentType,
		AcceptContextTypes: config.AcceptContextTypes,
	}

	return NewClient(hostUrl, clientConfig, client)
}

func (c *Client) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

func (c *Client) Post() *Request {
	return c.Verb("POST")
}

func (c *Client) Put() *Request {
	return c.Verb("PUT")
}

func (c *Client) Get() *Request {
	return c.Verb("GET")
}

func (c *Client) Delete() *Request {
	return c.Verb("DELETE")
}

func (c *ClientConfig) HasBasicAuth() bool {
	return len(c.Username) != 0
}

func (c *ClientConfig) HasTokenAuth() bool {
	return len(c.BearerToken) != 0
}
