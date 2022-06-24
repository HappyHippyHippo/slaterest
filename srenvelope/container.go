package srenvelope

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slaterest"
)

// GetMiddlewareGenerator will try to retrieve the registered logging middleware
// for ok responses instance from the application service container.
func GetMiddlewareGenerator(c slate.ServiceContainer) (func(string) (slaterest.Middleware, error), error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(func(string) (slaterest.Middleware, error))
	if !ok {
		return nil, errConversion(instance, "func(string) (slaterest.Middleware, error)")
	}
	return i, nil
}
