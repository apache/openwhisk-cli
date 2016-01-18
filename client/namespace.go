package client

import "net/http"

type Namespace string

type NamespaceService struct {
	client *Client
}

func (s *NamespaceService) List() ([]Namespace, *http.Response, error) {
	// make a request to c.BaseURL

	url := s.client.BaseURL.String() // "whisk.stage1.ng.bluemix.net"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	s.client.addAuthHeader(req)

	var namespaces []Namespace
	resp, err := s.client.Do(req, &namespaces)
	if err != nil {
		return nil, resp, err
	}

	return namespaces, resp, nil
}
