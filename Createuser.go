
package main

import (
"github.com/labstack/echo"
"net/http"
)

func CreateUser(c echo.Context) error {

	//migrate(db)
	player := user{}
	err = c.Bind(&player)

	name := player.Name
	b := []byte(player.Password)
	email := player.Email
	hashPass := hashAndSalt(b)
	num := player.PhNum
	cmpname:=player.CompanyName


	db.Where("account_name = ?", cmpname).First(&accobj)
	aid:=accobj.AccId

	err :=db.Create(&user{Name: name, Email: email, Password: hashPass, PhNum: num,Aid:aid,CompanyName:cmpname}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,"duplicate entry for key Email/company does not exist")
	}
	return c.JSON(http.StatusCreated, "new user created successfully")

}


