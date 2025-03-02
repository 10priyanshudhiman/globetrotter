package repository

import (
	"common/config"
	"common/model"
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

type userRegistryRepository struct {
	redisClient *redis.Client
}

type UserRegistryRepository interface {
	GetUserDetails(ctx context.Context, userId string) (model.User, error)
	UpdateUserDetails(ctx context.Context, userAction model.UserAction) (model.User, error)
}

func NewUserRegistryRepository(redisClient *redis.Client) UserRegistryRepository {
	return &userRegistryRepository{redisClient}
}

func (urr *userRegistryRepository) GetUserDetails(ctx context.Context, userId string) (model.User, error) {
	redisKey := config.C.Server.Env + config.C.Redis.Key.User
	user, err := urr.redisClient.HGet(ctx, redisKey, userId).Result()
	if err != nil && err != redis.Nil {
		return model.User{}, err
	}
	if err == redis.Nil {
		newUser := model.User{
			UserId: userId,
		}
		userJSON, err := json.Marshal(newUser)
		if err != nil {
			return model.User{}, err
		}

		_, err = urr.redisClient.HSet(ctx, redisKey, userId, userJSON).Result()
		if err != nil {
			return model.User{}, err
		}
		newUserDetails, err := urr.redisClient.HGet(ctx, redisKey, userId).Result()
		if err != nil {
			return model.User{}, err
		}

		var userModel model.User
		err = json.Unmarshal([]byte(newUserDetails), &userModel)
		if err != nil {
			return model.User{}, err
		}
		return userModel, nil

	} else {
		var userModel model.User
		err = json.Unmarshal([]byte(user), &userModel)
		if err != nil {
			return model.User{}, err
		}
		return userModel, nil
	}
}

func (urr *userRegistryRepository) UpdateUserDetails(ctx context.Context, userAction model.UserAction) (model.User, error) {
	redisKey := config.C.Server.Env + config.C.Redis.Key.User
	userId := userAction.UserId
	details, err := urr.redisClient.HGet(ctx, redisKey, userId).Result()
	if err != nil {
		return model.User{}, err
	}

	var userModel model.User
	err = json.Unmarshal([]byte(details), &userModel)
	if err != nil {
		return model.User{}, err
	}

	switch userAction.IsCorrect {
	case true:
		userModel.TotalScore += (userAction.Point)
		userModel.CorrectAnswers += 1
	case false:
		userModel.TotalScore += (userAction.Point)
		userModel.IncorrectAnswers += 1
	default:
		return model.User{}, errors.New("InValid userAction")
	}

	userModelJSON, err := json.Marshal(userModel)
	if err != nil {
		return model.User{}, err
	}

	_, err = urr.redisClient.HSet(ctx, redisKey, userId, userModelJSON).Result()
	if err != nil {
		return model.User{}, err
	}

	return userModel, nil

}
