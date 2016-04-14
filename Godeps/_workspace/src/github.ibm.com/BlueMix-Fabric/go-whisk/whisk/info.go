package whisk

import (
	"net/http"
	"net/url"
)

type Info struct {
	Whisk   string `json:"whisk,omitempty"`
	Version string `json:"version,omitempty"`
	Build   string `json:"build,omitempty"`
}

type InfoService struct {
	client *Client
}

func (s *InfoService) Get() (*Info, *http.Response, error) {
	// make a request to c.BaseURL / v1

	ref, err := url.Parse(s.client.Config.Version)
	if err != nil {
		return nil, nil, err
	}

	u := s.client.BaseURL.ResolveReference(ref)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	info := new(Info)
	resp, err := s.client.Do(req, &info)
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil
}
