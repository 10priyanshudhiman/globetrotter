package registry

import (
	"globetrotter_api/controller"

	"github.com/redis/go-redis/v9"
)

type registry struct {
	redisClient *redis.Client
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(redisClient *redis.Client) Registry {
	return &registry{redisClient: redisClient}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		UserRegistry: r.NewUserRegistryController(),
		Game:         r.NewGameController(),
	}
}
