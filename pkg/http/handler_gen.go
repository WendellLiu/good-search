// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/wendellliu/good_search/pkg/endpoint"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeSearchCompanyHandler(m, endpoints, options["SearchCompany"])
	return m
}
