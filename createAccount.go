package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func CreateAccount(c echo.Context) error {

	player := Account{}
	err = c.Bind(&player)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	acctype := player.AccType
	accname := player.AccountName

	err:=db.Create(&Account{AccType: acctype, AccountName: accname}).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Account name should be unique")
	}
	return c.JSON(http.StatusCreated, player)
}
