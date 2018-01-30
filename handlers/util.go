package handlers

import (
	"github.com/labstack/echo"
	"strconv"
)

func page(c echo.Context) int {
	page := c.QueryParam("p")
	var p int
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 0
	}
	return p
}
