package user

import (
	svcv1 "github.com/jaguar/apiskeleton/internal/service/v1"
	"github.com/jaguar/apiskeleton/internal/store"
)

// Controller create a user handler used to handle request for user resource.
type Controller struct {
	svc svcv1.Interface
}

// New creates a user Controller.
func New(store store.Factory) *Controller {
	return &Controller{
		svc: svcv1.NewService(store),
	}
}
