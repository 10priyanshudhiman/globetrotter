package registry

import (
	"globetrotter_api/controller"
	"globetrotter_api/repository"
	"globetrotter_api/service"
)

func (r *registry) NewGameController() controller.GameController {
	return controller.NewGameController(r.NewGameService())
}

func (r *registry) NewGameService() service.GameService {
	return service.NewGameService(r.NewGameRepository())
}

func (r *registry) NewGameRepository() repository.GameRepository {
	return repository.NewGameRepository(r.redisClient)
}
