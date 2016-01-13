package client

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addRouteOptions(route string, options interface{}) (string, error) {
	v := reflect.ValueOf(options)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return route, nil
	}

	u, err := url.Parse(route)
	if err != nil {
		return route, err
	}

	qs, err := query.Values(options)
	if err != nil {
		return route, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
