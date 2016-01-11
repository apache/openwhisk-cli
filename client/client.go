package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"url"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "https://whisk.com" // TODO :: insert real url
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	// TODO :: put state in here
	// authToken string // etc.
	// version string
	// verbose bool

	Sdk        *SdkService
	Trigger    *TriggerService
	Action     *ActionService
	Rule       *RuleService
	Activation *ActivationService
	Package    *PackageService
}

func NewClient(httpClient *http.Client) (c *Client) {

	if httpClient == nil {
		httpClient = http.defaultClient
	}
	baseURL := url.Parse(defaultBaseURL)

	c := &Client{
		client:  httpClient,
		baseURL: baseURL,
	}

	c.Sdk = &SdkService{client: c}
	c.Trigger = &TriggerService{client: c}
	c.Action = &ActionService{client: c}
	c.Rule = &RuleService{client: c}
	c.Activation = &ActivationService{client: c}
	c.Package = &PackageService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaTypeV3)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	return response, err
}

// Auth performs authorization operation --> stores token in client
func (c *Client) Auth(authKey string) error {
	// Does auth, stores token in client
}

// Clean resets object state (cache + auth)
func (c *Client) Clean() {

}

// Version returns the version of the API
func (c *Client) Version() string {

}

//List returns lists of all actions, triggers, rules, and activations.
func (c *Client) List() (actions []Action, triggers []Trigger, rules []Rule, activations []Activation, err error) {
	actions, err = c.ActionService.List()
	if err != nil {
		return
	}

	triggers, err = c.TriggerService.List()
	if err != nil {
		return
	}

	rules, err = c.RuleService.List()
	if err != nil {
		return
	}

	activations, err = c.ActivationService.List()
	if err != nil {
		return
	}

	return
}

// CrudService is a generic CRUD mixin for other services to provide basic CRUD functionality.
type crudService struct {
	resource string
}

// TODO :: crud service functions (Create, update, list etc...)
