package main

import (
	"common/config"
	"common/datastore"
	"fmt"
	"globetrotter_api/registry"
	"globetrotter_api/router"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ReadConfig("../common/config")

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s", config.C.Log.Path, config.C.Log.ApiName),
		rotatelogs.WithRotationTime(time.Duration(config.C.Log.RotationTime)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(config.C.Log.MaxAge)),
	)
	if err != nil {
		log.Fatalf("Failed to initialize log file rotation: %v", err)
	}

	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(logrus.StandardLogger().Out)

	redisClient := datastore.RedisClient()
	defer redisClient.Close()

	err = datastore.LoadCityDataToRedis(redisClient)
	if err != nil {
		log.Println(err)
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowMethods},
	}))

	r := registry.NewRegistry(redisClient)

	e = router.NewRouter(e, r.NewAppController())

	log.Println("API Server listening on port", config.C.Server.ApiPort)
	if err := e.Start(config.C.Server.ApiPort); err != nil {
		log.Fatalln(err)
	}

}
