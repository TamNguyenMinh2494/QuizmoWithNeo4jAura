package apis

import (
	"quizmo/bl"
	"quizmo/core"

	"github.com/labstack/echo/v4"
)

func Movies(server *core.Server) *echo.Group {
	api := server.Echo.Group("/api/v1")
	business := bl.NewMovieBusinessLogic(server)

	api.GET("/", func(c echo.Context) error {
		result, err := business.RelationshipBetweenPeople()
		if err != nil {
			return err
		}
		return c.JSON(200, result)
	})
	return api
}
