package service

import (
	"context"
	"globetrotter_api/model"
	"globetrotter_api/repository"
)

type userRegistryService struct {
	userRegistryRepository repository.UserRegistryRepository
}

type UserRegistryService interface {
	GetUserDetails(ctx context.Context, userId string) (model.User, error)
	UpdateUserDetails(ctx context.Context, userAction model.UserAction) (model.User, error)
}

func NewUserRegisrtyService(urr repository.UserRegistryRepository) UserRegistryService {
	return &userRegistryService{urr}
}

func (urs *userRegistryService) GetUserDetails(ctx context.Context, userId string) (model.User, error) {
	user, err := urs.userRegistryRepository.GetUserDetails(ctx, userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (urs *userRegistryService) UpdateUserDetails(ctx context.Context, userAction model.UserAction) (model.User, error) {
	updatedUser, err := urs.userRegistryRepository.UpdateUserDetails(ctx, userAction)
	if err != nil {
		return model.User{}, err
	}
	return updatedUser, nil
}
