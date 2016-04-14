package whisk

import (
	"fmt"
	"net/http"
)

type ActivationService struct {
	client *Client
}

type Activation struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	Subject      string `json:"subject,omitempty"`
	ActivationID string `json:"activationId,omitempty"`
	Cause        string `json:"cause,omitempty"`
	Start        int64  `json:"start,omitempty"`
	End          int64  `json:"end,omitempty"`
	Response     `json:"response,omitempty"`
	Logs         []Log `json:"logs,omitempty"`
}

type Response struct {
	Status     string `json:"status,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Result     `json:"result,omitempty"`
}

type Result map[string]interface{}

type ActivationListOptions struct {
	Name  string `url:"name,omitempty"`
	Limit int    `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Since int64  `url:"since,omitempty"`
	Upto  int64  `url:"upto,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}

type Log struct {
	Log    string `json:"log,omitempty"`
	Stream string `json:"stream,omitempty"`
	Time   string `json:"time,omitempty"`
}

func (s *ActivationService) List(options *ActivationListOptions) ([]Activation, *http.Response, error) {
	route := "activations"
	route, err := addRouteOptions(route, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	var activations []Activation
	resp, err := s.client.Do(req, &activations)
	if err != nil {
		return nil, resp, err
	}

	return activations, resp, err

}

func (s *ActivationService) Get(activationID string) (*Activation, *http.Response, error) {
	route := fmt.Sprintf("activations/%s", activationID)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Activation)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

func (s *ActivationService) Logs(activationID string) (*Activation, *http.Response, error) {
	route := fmt.Sprintf("activations/%s/logs", activationID)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	activation := new(Activation)
	resp, err := s.client.Do(req, &activation)
	if err != nil {
		return nil, resp, err
	}

	return activation, resp, nil
}

func (s *ActivationService) Result(activationID string) (*Response, *http.Response, error) {
	route := fmt.Sprintf("activations/%s", activationID)

	req, err := s.client.NewRequest("get", route, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Response)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil

}
