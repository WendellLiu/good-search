package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/wendellliu/good_search/pkg/service"
)

// SearchCompanyRequest collects the request parameters for the SearchCompany method.
type SearchCompanyRequest struct {
	CompanyName string `json:"company_name"`
}

// SearchCompanyResponse collects the response parameters for the SearchCompany method.
type SearchCompanyResponse struct {
	Err error `json:"err"`
}

// MakeSearchCompanyEndpoint returns an endpoint that invokes SearchCompany on the service.
func MakeSearchCompanyEndpoint(s service.GoodSearchService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchCompanyRequest)
		err := s.SearchCompany(ctx, req.CompanyName)
		return SearchCompanyResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r SearchCompanyResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SearchCompany implements Service. Primarily useful in a client.
func (e Endpoints) SearchCompany(ctx context.Context, companyName string) (err error) {
	request := SearchCompanyRequest{CompanyName: companyName}
	response, err := e.SearchCompanyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SearchCompanyResponse).Err
}
