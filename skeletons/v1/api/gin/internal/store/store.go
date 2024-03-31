package store

import (
	"context"

	v1 "github.com/jaguar/apiskeleton/pkg/api/apiserver/v1"
	metav1 "github.com/jaguar/apiskeleton/pkg/api/meta/v1"
)

var _client Factory

// Factory defines the idm platform storage interface.
type Factory interface {
	Users() UserStore
	Close() error
}

// Client returns the store client instance.
func Client() Factory {
	return _client
}

// SetClient set the idm store client.
func SetClient(factory Factory) {
	_client = factory
}

// UserStore defines the user storage interface.
type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteList(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
}
