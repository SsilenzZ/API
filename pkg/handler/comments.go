package handler

import (
	"Api/pkg/db/repository/comment"
	"Api/pkg/db/repository/user"
	"Api/pkg/service"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ReturnAllComments(c echo.Context) error {
	result := comment.Repository.GetAllComments()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}

	return c.XML(http.StatusOK, result)
}

func ReturnComment(c echo.Context) error {
	result, err := comment.Repository.GetComment(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}

	return c.XML(http.StatusOK, result)
}

func CreateComment(c echo.Context) error {
	user_ := c.Get("user").(*jwt.Token)
	claims := user_.Claims.(*service.JWTCustomClaims)

	email, err := user.Repository.GetEmail(claims.ID)
	if err != nil {
		return err
	}
	request := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["email"] = email
	comment.Repository.CreateComment(request)

	return c.JSON(http.StatusOK, nil)
}

func UpdateComment(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = comment.Repository.UpdateComment(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteComment(c echo.Context) error {
	comment.Repository.DeleteComment(c.Param("id"))

	return c.JSON(http.StatusOK, nil)
}
