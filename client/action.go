package client

import (
	"fmt"
	"net/http"
)

type ActionService struct {
	client *Client
}

type Action struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"` // NOTE :: this might not include in json if its false ... would be an issue if server default is true

	Exec        `json:"exec,omitempty"`
	Annotations `json:"annotations,omitempty"`
	Parameters  `json:"parameters,omitempty"`
	Limits      `json:"limits,omitempty"`
}

type ActionRequest struct {
	// Use this if POST /actions requires different parameters than above.
}

type Exec struct {
	Code  string `json:"code,omitempty"`
	Image string `json:"image,omitempty"`
	Init  string `json:"init,omitempty"`
}

type ActionListOptions struct {
	Limit string `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}

////////////////////
// Action Methods //
////////////////////

func (s *ActionService) List(options *ActionListOptions) ([]Action, *http.Response, error) {
	route := fmt.Sprintf("actions")
	route, err := addRouteOptions(route, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	var actions []Action
	resp, err := s.client.Do(req, &actions)
	if err != nil {
		return nil, resp, err
	}

	return actions, resp, err

}

func (s *ActionService) Create(action *Action, blocking bool) (*Action, *http.Response, error) {
	route := fmt.Sprintf("actions?blocking=%s", blocking)

	req, err := s.client.NewRequest("POST", route, action)
	if err != nil {
		return nil, nil, err
	}

	a := new(Action)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

func (s *ActionService) Fetch(actionName string) (*Action, *http.Response, error) {
	route := fmt.Sprintf("actions/%s", actionName)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Action)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

func (s *ActionService) Delete(actionName string) (*http.Response, error) {
	route := fmt.Sprintf("actions/%s", actionName)

	req, err := s.client.NewRequest("DELETE", route, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *ActionService) Invoke(actionName string, blocking bool) (*Action, *http.Response, error) {
	route := fmt.Sprintf("actions/%s?blocking=", actionName, blocking)

	req, err := s.client.NewRequest("POST", route, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Action)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil

}

func (s *ActionService) Update(action *Action, overwrite bool) (*Action, *http.Response, error) {
	route := fmt.Sprintf("actions/%s?overwrite=", action.Name, overwrite)

	req, err := s.client.NewRequest("POST", route, action)
	if err != nil {
		return nil, nil, err
	}

	a := new(Action)
	resp, err := s.client.Do(req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}
