package service

import "context"

type LoginRes struct{}

// GokitusercenterService describes the service.
type GokitusercenterService interface {
	// Add your methods here
	Login(ctx context.Context, username, password string) (rs LoginRes, err error)
}

type basicGokitusercenterService struct{}

func (b *basicGokitusercenterService) Login(ctx context.Context, username string, password string) (rs LoginRes, err error) {
	// TODO implement the business logic of Login
	return rs, err
}

// NewBasicGokitusercenterService returns a naive, stateless implementation of GokitusercenterService.
func NewBasicGokitusercenterService() GokitusercenterService {
	return &basicGokitusercenterService{}
}

// New returns a GokitusercenterService with all of the expected middleware wired in.
func New(middleware []Middleware) GokitusercenterService {
	var svc GokitusercenterService = NewBasicGokitusercenterService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
