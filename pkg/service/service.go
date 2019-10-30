package service

import "context"

// GoodSearchService describes the service.
type GoodSearchService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SearchCompany(ctx context.Context, companyName string) (err error)
}

type basicGoodSearchService struct{}

func (b *basicGoodSearchService) SearchCompany(ctx context.Context, companyName string) (err error) {
	// TODO implement the business logic of SearchCompany
	return err
}

// NewBasicGoodSearchService returns a naive, stateless implementation of GoodSearchService.
func NewBasicGoodSearchService() GoodSearchService {
	return &basicGoodSearchService{}
}

// New returns a GoodSearchService with all of the expected middleware wired in.
func New(middleware []Middleware) GoodSearchService {
	var svc GoodSearchService = NewBasicGoodSearchService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
