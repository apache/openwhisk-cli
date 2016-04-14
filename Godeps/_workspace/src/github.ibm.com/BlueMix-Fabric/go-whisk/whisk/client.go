package whisk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://whisk.stage1.ng.bluemix.net:443/api/"
)

type Client struct {
	client *http.Client
	*Config

	Sdks        *SdkService
	Triggers    *TriggerService
	Actions     *ActionService
	Rules       *RuleService
	Activations *ActivationService
	Packages    *PackageService
	Namespaces  *NamespaceService
	Info        *InfoService
}

type Config struct {
	Namespace string // NOTE :: Default is "_"
	AuthToken string
	BaseURL   *url.URL // NOTE :: Default is "https://whisk.stage1.ng.bluemix.net:443/api/"
	Version   string
	Verbose   bool
}

func NewClient(httpClient *http.Client, config *Config) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var err error
	if config.BaseURL == nil {
		config.BaseURL, err = url.Parse(defaultBaseURL)
		if err != nil {
			return nil, err
		}
	}

	if len(config.Namespace) == 0 {
		config.Namespace = "_"
	}

	if len(config.Version) == 0 {
		config.Version = "v1"
	}

	c := &Client{
		client: httpClient,
		Config: config,
	}

	c.Sdks = &SdkService{client: c}
	c.Triggers = &TriggerService{client: c}
	c.Actions = &ActionService{client: c}
	c.Rules = &RuleService{client: c}
	c.Activations = &ActivationService{client: c}
	c.Packages = &PackageService{client: c}
	c.Namespaces = &NamespaceService{client: c}
	c.Info = &InfoService{client: c}

	return c, nil
}

///////////////////////////////
// Request/Utility Functions //
///////////////////////////////

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	urlStr = fmt.Sprintf("%s/namespaces/%s/%s", c.Config.Version, c.Config.Namespace, urlStr)

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
	if req.Body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	c.addAuthHeader(req)

	return req, nil
}

func (c *Client) addAuthHeader(req *http.Request) {
	if c.Config.AuthToken != "" {
		encodedAuthToken := base64.StdEncoding.EncodeToString([]byte(c.Config.AuthToken))
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedAuthToken))
	}
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// TODO :: clean up verbose scheme
	if c.Verbose {
		fmt.Printf("\n[%s]\t%s\n", req.Method, req.URL)
		if len(req.Header) > 0 {
			fmt.Println("Headers")
			printJSON(req.Header)
		}
		if req.Body != nil {
			printJSON(req.Body)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		if c.Verbose {
			printJSON(err)
		}
		return resp, err
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

	if c.Verbose {
		fmt.Printf("\n[%d]\t%s\n", resp.StatusCode, resp.Status)
		printJSON(v)
	}

	return resp, err
}

////////////
// Errors //
////////////

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}

type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

////////////////////////////
// Basic Client Functions //
////////////////////////////
