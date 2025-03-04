package controller

import (
	"context"
	"errors"
	"globetrotter_api/model"
	"globetrotter_api/service"
	"globetrotter_api/utility"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type gameController struct {
	gameService service.GameService
}

type GameController interface {
	GetAllCountries(c echo.Context) error
	GetInitialGameDetails(c echo.Context) error
	VerifyUserAction(c echo.Context) error
}

func NewGameController(gs service.GameService) GameController {
	return &gameController{gs}
}

func (gs *gameController) GetInitialGameDetails(c echo.Context) error {
	ctx := context.Background()

	country := c.QueryParam("country")
	log.Println("RECIEVED Country", country)

	user, err := gs.gameService.GetInitialGameDetails(ctx, country)
	if err != nil {
		log.Println("error in fetching game details", err)
		return c.JSON(http.StatusInternalServerError, &model.Message{Msg: "Failed to fetch game Details", Data: nil})
	}
	return c.JSON(http.StatusOK, &model.Message{Msg: "Successfully fetched game details", Data: user})
}

func (gs *gameController) VerifyUserAction(c echo.Context) error {
	ctx := context.Background()

	var userAction []model.VerifyAction
	if err := c.Bind(&userAction); !errors.Is(err, nil) {
		return c.JSON(http.StatusBadRequest, &model.Message{Msg: utility.ErrBadRequest, Data: nil})
	}

	log.Println("RECIEVED verify user action", userAction)

	user, err := gs.gameService.VerifyUserAction(ctx, userAction)
	if err != nil {
		log.Println("error in verify user action", err)
		return c.JSON(http.StatusInternalServerError, &model.Message{Msg: "Failed to verify user action", Data: nil})
	}
	return c.JSON(http.StatusOK, &model.Message{Msg: "Successfully verify user action", Data: user})
}

func (gs *gameController) GetAllCountries(c echo.Context) error {
	ctx := context.Background()

	data, err := gs.gameService.GetAllCountries(ctx)
	if err != nil {
		log.Println("error in getting all countries", err)
		return c.JSON(http.StatusInternalServerError, &model.Message{Msg: "Failed to get all countries", Data: nil})
	}
	return c.JSON(http.StatusOK, &model.Message{Msg: "Successfully get all countries", Data: data})
}
