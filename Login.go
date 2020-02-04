package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context,) error {

	player := user{}
	err = c.Bind(&player)
	email:=player.Email


	db.Where("email = ?", email).First(&userobj)
	dbpwd:=userobj.Password

	pwd:=[]byte(player.Password)


	pwdMatch := comparePasswords(dbpwd, pwd)
	fmt.Println("Passwords Match?", pwdMatch)

	if pwdMatch == true {
		return c.JSON(http.StatusCreated, "user login successfully")
	}
	return c.JSON(http.StatusCreated, "Invalid login credentials")

}
