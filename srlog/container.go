package srlog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slaterest"
)

// GetMiddlewareOk will try to retrieve the registered logging middleware
// for ok responses instance from the application service container.
func GetMiddlewareOk(c slate.ServiceContainer) (slaterest.Middleware, error) {
	instance, err := c.Get(ContainerOkID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(slaterest.Middleware)
	if !ok {
		return nil, errConversion(instance, "slaterest.Middleware")
	}
	return i, nil
}

// GetMiddlewareCreated will try to retrieve the registered logging middleware
// for created responses instance from the application service container.
func GetMiddlewareCreated(c slate.ServiceContainer) (slaterest.Middleware, error) {
	instance, err := c.Get(ContainerCreatedID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(slaterest.Middleware)
	if !ok {
		return nil, errConversion(instance, "slaterest.Middleware")
	}
	return i, nil
}

// GetMiddlewareNoContent will try to retrieve the registered logging middleware
// for no-content responses instance from the application service container.
func GetMiddlewareNoContent(c slate.ServiceContainer) (slaterest.Middleware, error) {
	instance, err := c.Get(ContainerNoContentID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(slaterest.Middleware)
	if !ok {
		return nil, errConversion(instance, "slaterest.Middleware")
	}
	return i, nil
}
