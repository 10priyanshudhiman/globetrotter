package service

import (
	"context"
	"globetrotter_api/model"
	"globetrotter_api/repository"
)

type gameService struct {
	gameRepository repository.GameRepository
}

type GameService interface {
	GetAllCountries(ctx context.Context) (map[string]string, error)
	GetInitialGameDetails(ctx context.Context, country string) (model.Game, error)
	VerifyUserAction(ctx context.Context, actions []model.VerifyAction) (model.ActionResponse, error)
}

func NewGameService(gr repository.GameRepository) GameService {
	return &gameService{gr}
}

func (gs *gameService) GetInitialGameDetails(ctx context.Context, country string) (model.Game, error) {
	game, err := gs.gameRepository.GetInitialGameDetails(ctx, country)
	if err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (gs *gameService) VerifyUserAction(ctx context.Context, actions []model.VerifyAction) (model.ActionResponse, error) {

	game, err := gs.gameRepository.VerifyUserAction(ctx, actions)
	if err != nil {
		return model.ActionResponse{}, err
	}
	return game, nil
}

func (gs *gameService) GetAllCountries(ctx context.Context) (map[string]string, error) {
	data, err := gs.gameRepository.GetAllCountries(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
