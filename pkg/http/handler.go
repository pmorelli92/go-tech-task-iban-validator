package http

import (
	"github.com/labstack/echo/v4"
	"iban-validator/pkg/domain"
	"net/http"
)

func Bootstrap(port string) error {
	e := echo.New()
	e.GET("/validate/:iban", func(c echo.Context) error {
		valid := domain.Validate(c.Param("iban"))
		return c.JSON(http.StatusOK, &Valid{IsValid: valid})
	})

	err := e.Start(":" + port)
	return err
}

type Valid struct {
	IsValid bool
}
