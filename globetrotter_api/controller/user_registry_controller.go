package controller

import (
	"common/model"
	"common/utility"
	"context"
	"errors"
	"globetrotter_api/service"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userRegistryController struct {
	userRegistryService service.UserRegistryService
}

type UserRegistryController interface {
	GetUserDetails(c echo.Context) error
	UpdateUserDetails(c echo.Context) error
}

func NewUserRegistryController(urs service.UserRegistryService) UserRegistryController {
	return &userRegistryController{urs}
}

func (urc *userRegistryController) GetUserDetails(c echo.Context) error {
	ctx := context.Background()

	userId := c.QueryParam("userId")
	log.Println("RECIEVED userId", userId)

	user, err := urc.userRegistryService.GetUserDetails(ctx, userId)
	if err != nil {
		log.Println("error in fetching user details", err)
		return c.JSON(http.StatusInternalServerError, &model.Message{Msg: "Failed to fetch userDetails", Data: nil})
	}
	return c.JSON(http.StatusOK, &model.Message{Msg: "Successfully fetched user details", Data: user})

}

func (urc *userRegistryController) UpdateUserDetails(c echo.Context) error {
	ctx := context.Background()

	var userAction model.UserAction
	if err := c.Bind(&userAction); !errors.Is(err, nil) {
		return c.JSON(http.StatusBadRequest, &model.Message{Msg: utility.ErrBadRequest, Data: nil})
	}

	log.Println("RECIEVED user action", userAction)

	user, err := urc.userRegistryService.UpdateUserDetails(ctx, userAction)
	if err != nil {
		log.Println("error in updating user details", err)
		return c.JSON(http.StatusInternalServerError, &model.Message{Msg: "Failed to Update userDetails", Data: nil})
	}
	return c.JSON(http.StatusOK, &model.Message{Msg: "Successfully updated user details", Data: user})
}
