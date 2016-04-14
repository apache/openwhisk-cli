package whisk

import (
	"fmt"
	"net/http"
	"net/url"
)

type Namespace struct {
	Name     string `json:"name"`
	Contents struct {
		Actions  []Action  `json:"actions"`
		Packages []Package `json:"packages"`
		Triggers []Trigger `json:"triggers"`
		Rules    []Rule    `json:"rules"`
	} `json:"contents,omitempty"`
}

type NamespaceService struct {
	client *Client
}

// get a list of available namespaces
func (s *NamespaceService) List() ([]Namespace, *http.Response, error) {
	// make a request to c.BaseURL / namespaces

	urlStr := fmt.Sprintf("%s/namespaces", s.client.Config.Version)
	ref, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	u := s.client.BaseURL.ResolveReference(ref)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	s.client.addAuthHeader(req)

	var namespaceNames []string
	resp, err := s.client.Do(req, &namespaceNames)
	if err != nil {
		return nil, resp, err
	}

	var namespaces []Namespace
	for _, nsName := range namespaceNames {
		ns := Namespace{
			Name: nsName,
		}
		namespaces = append(namespaces, ns)
	}

	return namespaces, resp, nil
}

func (s *NamespaceService) Get(nsName string) (*Namespace, *http.Response, error) {

	// GET request to currently-set namespace (def. "_")

	if nsName == "" {
		nsName = s.client.Config.Namespace
	}

	urlStr := fmt.Sprintf("%s/namespaces/%s", s.client.Config.Version, nsName)
	ref, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	u := s.client.BaseURL.ResolveReference(ref)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	s.client.addAuthHeader(req)

	ns := &Namespace{
		Name: nsName,
	}
	resp, err := s.client.Do(req, &ns.Contents)
	if err != nil {
		return nil, resp, err
	}

	return ns, resp, nil
}
