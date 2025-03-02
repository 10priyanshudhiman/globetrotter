package datastore

import (
	"common/config"
	"common/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
)

func convertRawJson(rawData string) ([]model.City, error) {
	var cities []model.City
	if err := json.Unmarshal([]byte(rawData), &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func storeMappingsToRedis(redisClient *redis.Client, redisKey string, data map[string]string) error {
	ctx := context.Background()
	for id, value := range data {
		if _, err := redisClient.HSet(ctx, redisKey, id, value).Result(); err != nil {
			return fmt.Errorf("failed to set %s for key %s: %v", id, redisKey, err)
		}
	}
	return nil
}

func getMappingsFromRedis(redisClient *redis.Client, redisKey string) (map[string]string, error) {
	ctx := context.Background()
	return redisClient.HGetAll(ctx, redisKey).Result()
}

func loadMappings(redisClient *redis.Client, cities []model.City) error {
	env := config.C.Server.Env

	mappingKeys := map[string]string{
		"clues":  env + config.C.Redis.Key.CluesMapping,
		"facts":  env + config.C.Redis.Key.FunFactMapping,
		"trivia": env + config.C.Redis.Key.TriviaMapping,
		"cities": env + config.C.Redis.Key.CountryCity,
	}

	mappings := map[string]map[string]string{
		"clues":  {},
		"facts":  {},
		"trivia": {},
		"cities": {},
	}

	counters := map[string]int{"clues": 1, "facts": 1, "trivia": 1}

	for _, city := range cities {
		for key, items := range map[string][]string{"clues": city.Clues, "facts": city.FunFact, "trivia": city.Trivia} {
			for _, item := range items {
				id := fmt.Sprintf("%s_%d", key, counters[key])
				mappings[key][id] = item
				counters[key]++
			}
		}

		country := city.Country
		if existingCities, exists := mappings["cities"][country]; exists {
			mappings["cities"][country] = existingCities + "#" + city.City
		} else {
			mappings["cities"][country] = city.City
		}
	}

	for key, data := range mappings {
		if err := storeMappingsToRedis(redisClient, mappingKeys[key], data); err != nil {
			return err
		}
	}
	return nil
}

func loadCountryData(redisClient *redis.Client, cities []model.City) error {
	env := config.C.Server.Env

	countryKeys := map[string]string{
		"clues":  env + config.C.Redis.Key.CountryClues,
		"facts":  env + config.C.Redis.Key.CountryFunFact,
		"trivia": env + config.C.Redis.Key.CountryTrivia,
	}

	countryCityKeys := map[string]string{
		"clues":  env + config.C.Redis.Key.CountryCityClues,
		"facts":  env + config.C.Redis.Key.CountryCityFunfact,
		"trivia": env + config.C.Redis.Key.CountryCityTrivia,
	}

	redisMappings := map[string]map[string]string{
		"clues":  {},
		"facts":  {},
		"trivia": {},
	}

	for key, redisKey := range map[string]string{
		"clues":  env + config.C.Redis.Key.CluesMapping,
		"facts":  env + config.C.Redis.Key.FunFactMapping,
		"trivia": env + config.C.Redis.Key.TriviaMapping,
	} {
		mapping, err := getMappingsFromRedis(redisClient, redisKey)
		if err != nil {
			return fmt.Errorf("error fetching %s mappings: %v", key, err)
		}
		redisMappings[key] = mapping
	}

	redisValueToIDs := map[string]map[string][]string{
		"clues":  {},
		"facts":  {},
		"trivia": {},
	}

	for key, mapping := range redisMappings {
		for id, value := range mapping {
			redisValueToIDs[key][value] = append(redisValueToIDs[key][value], id)
		}
	}

	countryData := map[string]map[string]string{
		"clues":  {},
		"facts":  {},
		"trivia": {},
	}

	countryCityData := map[string]map[string]string{
		"clues":  {},
		"facts":  {},
		"trivia": {},
	}

	for _, city := range cities {
		country := city.Country
		countryCity := fmt.Sprintf("%s#%s", country, city.City)

		for key, items := range map[string][]string{
			"clues":  city.Clues,
			"facts":  city.FunFact,
			"trivia": city.Trivia,
		} {
			for _, item := range items {
				ids, exists := redisValueToIDs[key][item]
				if !exists {
					log.Printf("No mapping found for %s: %s", key, item)
					continue
				}

				idsStr := strings.Join(ids, "#")

				if existing, ok := countryData[key][country]; ok {
					countryData[key][country] = existing + "#" + idsStr
				} else {
					countryData[key][country] = idsStr
				}

				if existing, ok := countryCityData[key][countryCity]; ok {
					countryCityData[key][countryCity] = existing + "#" + idsStr
				} else {
					countryCityData[key][countryCity] = idsStr
				}
			}
		}
	}

	for key, data := range countryData {
		if err := storeMappingsToRedis(redisClient, countryKeys[key], data); err != nil {
			return err
		}
	}

	for key, data := range countryCityData {
		if err := storeMappingsToRedis(redisClient, countryCityKeys[key], data); err != nil {
			return err
		}
	}

	return nil
}

func LoadCityDataToRedis(redisClient *redis.Client) error {
	rawData := RawJSON
	cities, err := convertRawJson(rawData)
	if err != nil {
		return err
	}

	if err := loadMappings(redisClient, cities); err != nil {
		return err
	}

	if err := loadCountryData(redisClient, cities); err != nil {
		return err
	}

	return nil
}

func DeleteMatchingKeys(redisClient *redis.Client, pattern string) error {
	ctx := context.Background()
	var cursor uint64
	var keys []string
	var err error

	// Use SCAN to find keys matching the pattern
	for {
		keys, cursor, err = redisClient.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			log.Println("Error scanning keys:", err)
			return err
		}

		if len(keys) > 0 {
			if err := redisClient.Del(ctx, keys...).Err(); err != nil {
				log.Println("Failed to delete keys:", keys, err)
				return err
			} else {
				log.Println("Deleted keys:", keys)
			}
		}

		// Stop when cursor is 0 (end of iteration)
		if cursor == 0 {
			break
		}
	}

	return nil
}
