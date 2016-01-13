package client

import (
	"fmt"
	"net/http"
)

type PackageService struct {
	client *Client
}

type Package struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	Annotations `json:"annotations"`
	Parameters  `json:"parameters"`
	Binding     `json:"binding"`
}

type Binding struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type PackageListOptions struct {
	Public bool   `url:"public,omitempty"`
	Limit  string `url:"limit,omitempty"`
	Skip   int    `url:"skip,omitempty"`
	Since  int    `url:"since,omitempty"`
	Docs   bool   `url:"docs,omitempty"`
}

func (s *PackageService) List(options *PackageListOptions) ([]Package, *http.Response, error) {
	route := fmt.Sprintf("packages")
	route, err := addRouteOptions(route, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	var packages []Package
	resp, err := s.client.Do(req, &packages)
	if err != nil {
		return nil, resp, err
	}

	return packages, resp, err

}

func (s *PackageService) Create(x_package *Package, blocking bool) (*Package, *http.Response, error) {
	route := "packages"

	req, err := s.client.NewRequest("POST", route, x_package)
	if err != nil {
		return nil, nil, err
	}

	p := new(Package)
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil

}

func (s *PackageService) Fetch(packageName string) (*Package, *http.Response, error) {
	route := fmt.Sprintf("packages/%s", packageName)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	p := new(Package)
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil

}

func (s *PackageService) Update(x_package *Package, overwrite bool) (*Package, *http.Response, error) {
	route := fmt.Sprintf("packages/%s?overwrite=", x_package.Name, overwrite)

	req, err := s.client.NewRequest("POST", route, x_package)
	if err != nil {
		return nil, nil, err
	}

	p := new(Package)
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

func (s *PackageService) Delete(packageName string) (*http.Response, error) {
	route := fmt.Sprintf("packages/%s", packageName)

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
