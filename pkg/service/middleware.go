package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(GoodSearchService) GoodSearchService

type loggingMiddleware struct {
	logger log.Logger
	next   GoodSearchService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GoodSearchService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GoodSearchService) GoodSearchService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SearchCompany(ctx context.Context, companyName string) (err error) {
	defer func() {
		l.logger.Log("method", "SearchCompany", "companyName", companyName, "err", err)
	}()
	return l.next.SearchCompany(ctx, companyName)
}
