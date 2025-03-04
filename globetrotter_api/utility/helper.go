package utility

import (
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func PickRandomCities(cityList []string, num int) []string {
	if len(cityList) <= num {
		return cityList
	}

	rng.Shuffle(len(cityList), func(i, j int) { cityList[i], cityList[j] = cityList[j], cityList[i] })
	return cityList[:num]
}

func GetRandomCitiesFromOtherCountries(countryCitiesMap map[string]string, excludeCountry string, needed int) []string {
	var allCities []string
	for country, citiesStr := range countryCitiesMap {
		if country != excludeCountry {
			allCities = append(allCities, strings.Split(citiesStr, "#")...)
		}
	}

	return PickRandomCities(allCities, needed)
}
