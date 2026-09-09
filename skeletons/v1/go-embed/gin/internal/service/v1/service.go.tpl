package v1

import "{{ .App.ModuleName }}/internal/store"

// Interface defines functions used to return resource interface.
type Interface interface {
	Users() UserService
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Interface {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserService {
	return newUsers(s)
}
