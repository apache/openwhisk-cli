package client

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
	Start        string `json:"start,omitempty"`
	End          string `json:"end,omitempty"`
	Result       `json:"result,omitempty"`
	Logs         string `json:"logs,omitempty"`
}

type ActivationListOptions struct {
	Name  string `url:"name,omitempty"`
	Limit string `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Since int    `url:"since,omitempty"`
	Upto  int    `url:"upto,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}

type Result struct {
	Status string                 `json:"status,omitempty"`
	Value  map[string]interface{} `json:"value,omitempty"`
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

func (s *ActivationService) Fetch(activationID string) (*Activation, *http.Response, error) {
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

func (s *ActivationService) Result(activationID string) (*Result, *http.Response, error) {
	route := fmt.Sprintf("activations/%s", activationID)

	req, err := s.client.NewRequest("get", route, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(Result)
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil

}
