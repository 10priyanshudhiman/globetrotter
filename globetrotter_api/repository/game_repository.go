package repository

import (
	"context"
	"fmt"
	"globetrotter_api/config"
	"globetrotter_api/model"
	"globetrotter_api/utility"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
)

type gameRepository struct {
	redisClient *redis.Client
}

type GameRepository interface {
	GetAllCountries(ctx context.Context) (map[string]string, error)
	GetInitialGameDetails(ctx context.Context, country string) (model.Game, error)
	VerifyUserAction(ctx context.Context, actions []model.VerifyAction) (model.ActionResponse, error)
}

func NewGameRepository(redisClient *redis.Client) GameRepository {
	return &gameRepository{redisClient}
}

func (gr *gameRepository) GetInitialGameDetails(ctx context.Context, country string) (model.Game, error) {

	game := model.Game{
		Clues:  make(map[string]string),
		Cities: []string{},
	}
	env := config.C.Server.Env
	countryCluesKey := env + config.C.Redis.Key.CountryClues
	cluesKey := env + config.C.Redis.Key.CluesMapping
	countryCityKey := env + config.C.Redis.Key.CountryCity

	clues, err := gr.redisClient.HGet(ctx, countryCluesKey, country).Result()
	if err != nil {
		return model.Game{}, err
	}

	cluesMap, err := gr.redisClient.HGetAll(ctx, cluesKey).Result()
	if err != nil {
		return model.Game{}, err
	}

	countryCitiesMap, err := gr.redisClient.HGetAll(ctx, countryCityKey).Result()
	if err != nil {
		return model.Game{}, err
	}

	clueArr := strings.Split(clues, "#")
	selectedClues := utility.PickRandomCities(clueArr, 2)

	for _, clueId := range selectedClues {
		if clue, ok := cluesMap[clueId]; !ok {
			log.Println("mapping does not exists for clueId", clueId)
		} else {
			game.Clues[clueId] = clue
		}
	}

	var correctCities []string
	if citiesStr, exists := countryCitiesMap[country]; exists {
		correctCities = utility.PickRandomCities(strings.Split(citiesStr, "#"), 5)
	}

	incorrectCities := utility.GetRandomCitiesFromOtherCountries(countryCitiesMap, country, 10-len(correctCities))
	game.Cities = append(correctCities, incorrectCities...)

	return game, nil

}

func (gr *gameRepository) VerifyUserAction(ctx context.Context, actions []model.VerifyAction) (model.ActionResponse, error) {

	countryCityClues := config.C.Server.Env + config.C.Redis.Key.CountryCityClues
	countryCityFacts := config.C.Server.Env + config.C.Redis.Key.CountryCityFunfact
	factsKey := config.C.Server.Env + config.C.Redis.Key.FunFactMapping
	countryFacts := config.C.Server.Env + config.C.Redis.Key.CountryFunFact

	actionResponse := model.ActionResponse{
		Facts: make(map[string]string),
	}

	mapping, err := gr.redisClient.HGetAll(ctx, countryCityClues).Result()
	if err != nil {
		return model.ActionResponse{}, nil
	}

	funFacts, err := gr.redisClient.HGetAll(ctx, factsKey).Result()
	if err != nil {
		return model.ActionResponse{}, nil
	}

	isCorrect := true

	for _, action := range actions {
		key := fmt.Sprintf("%s#%s", action.Country, action.City)

		ids, exists := mapping[key]
		idsArr := strings.Split(ids, "#")
		if !exists {
			isCorrect = false
			log.Println("verify false", key, idsArr)
			break
		} else {
			isExists := false
			for _, id := range idsArr {
				if id == action.ClueId {
					isExists = true
				}
			}
			if !isExists {
				isCorrect = false
				log.Println("verify false", key, idsArr, action.ClueId)
				break
			}
		}
	}
	actionResponse.IsCorrect = isCorrect

	switch isCorrect {
	case true:
		for _, action := range actions {
			key := action.Country + "#" + action.City
			factsId, err := gr.redisClient.HGet(ctx, countryCityFacts, key).Result()
			if err != nil {
				return model.ActionResponse{}, nil
			}

			factsArr := strings.Split(factsId, "#")
			for _, factId := range factsArr {
				if fact, ok := funFacts[factId]; !ok {
					log.Println("factId mapping not exists", factId)
				} else {
					actionResponse.Facts[factId] = fact
				}
			}
		}
	case false:
		for _, action := range actions {
			key := action.Country
			factsId, err := gr.redisClient.HGet(ctx, countryFacts, key).Result()
			if err != nil {
				return model.ActionResponse{}, nil
			}

			factsArr := strings.Split(factsId, "#")
			for _, factId := range factsArr {
				if fact, ok := funFacts[factId]; !ok {
					log.Println("factId mapping not exists", factId)
				} else {
					actionResponse.Facts[factId] = fact
				}
			}

		}
	}

	return actionResponse, nil

}

func (gr *gameRepository) GetAllCountries(ctx context.Context) (map[string]string, error) {

	key := config.C.Server.Env + config.C.Redis.Key.CountryCity

	data, err := gr.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}
