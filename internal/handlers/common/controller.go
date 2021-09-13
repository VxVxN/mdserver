package common

import (
	e "github.com/VxVxN/mdserver/pkg/error"
)

type Controller struct {
	e.ErrResponseController
}

func NewController() *Controller {
	ctrl := &Controller{}
	return ctrl
}
