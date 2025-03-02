package registry

import (
	"globetrotter_api/controller"
	"globetrotter_api/repository"
	"globetrotter_api/service"
)

func (r *registry) NewUserRegistryController() controller.UserRegistryController {
	return controller.NewUserRegistryController(r.NewUserRegistryService())
}

func (r *registry) NewUserRegistryService() service.UserRegistryService {
	return service.NewUserRegisrtyService(r.NewUserRegistryRepository())
}

func (r *registry) NewUserRegistryRepository() repository.UserRegistryRepository {
	return repository.NewUserRegistryRepository(r.redisClient)
}
