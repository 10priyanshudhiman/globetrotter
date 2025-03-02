package router

import (
	"globetrotter_api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	//user details
	e.GET("/globetrotter-v1/user/details", func(context echo.Context) error { return c.UserRegistry.GetUserDetails(context) })
	e.POST("/globetrotter-v1/user/update", func(context echo.Context) error { return c.UserRegistry.UpdateUserDetails(context) })

	// Game
	e.GET("/globetrotter-v1/game/country", func(context echo.Context) error { return c.Game.GetAllCountries(context) })
	e.GET("/globetrotter-v1/game/details", func(context echo.Context) error { return c.Game.GetInitialGameDetails(context) })
	e.POST("/globetrotter-v1/game/verify", func(context echo.Context) error { return c.Game.VerifyUserAction(context) })
	return e
}
