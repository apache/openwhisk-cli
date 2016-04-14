package whisk

import (
	"fmt"
	"net/http"
)

type TriggerService struct {
	client *Client
}

type Trigger struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	ID          string `json:"id"`
	Annotations `json:"annotations"`
	Parameters  `json:"parameters"`
	Limits      `json:"limits"`
}

type TriggerListOptions struct {
	Limit int  `url:"limit,omitempty"`
	Skip  int  `url:"skip,omitempty"`
	Docs  bool `url:"docs,omitempty"`
}

func (s *TriggerService) List(options *TriggerListOptions) ([]Trigger, *http.Response, error) {
	route := "triggers"
	route, err := addRouteOptions(route, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	var triggers []Trigger
	resp, err := s.client.Do(req, &triggers)
	if err != nil {
		return nil, resp, err
	}

	return triggers, resp, err

}

func (s *TriggerService) Insert(trigger *Trigger, overwrite bool) (*Trigger, *http.Response, error) {
	route := fmt.Sprintf("triggers/%s?overwrite=%s", trigger.Name, overwrite)

	req, err := s.client.NewRequest("POST", route, trigger)
	if err != nil {
		return nil, nil, err
	}

	t := new(Trigger)
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil

}

func (s *TriggerService) Get(triggerName string) (*Trigger, *http.Response, error) {
	route := fmt.Sprintf("triggers/%s", triggerName)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	t := new(Trigger)
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil

}

func (s *TriggerService) Delete(triggerName string) (*http.Response, error) {
	route := fmt.Sprintf("triggers/%s", triggerName)

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

func (s *TriggerService) Fire(triggerName string, payload map[string]interface{}) (*Trigger, *http.Response, error) {
	route := fmt.Sprintf("triggers/", triggerName)

	req, err := s.client.NewRequest("POST", route, payload)
	if err != nil {
		return nil, nil, err
	}

	t := new(Trigger)
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil

}
