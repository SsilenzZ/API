package handler

import (
	"Api/pkg/db/repository/post"
	"Api/pkg/service"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func ReturnAllPosts(c echo.Context) error {
	result := post.Repository.GetAllPosts()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

func ReturnPost(c echo.Context) error {
	result, err := post.Repository.GetPost(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

func CreatePost(c echo.Context) error {
	user_ := c.Get("user").(*jwt.Token)
	claims := user_.Claims.(*service.JWTCustomClaims)
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["user"] = claims.ID
	post.Repository.CreatePost(request)
	return c.JSON(http.StatusOK, nil)
}

func UpdatePost(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = post.Repository.UpdatePost(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func DeletePost(c echo.Context) error {
	post.Repository.DeletePost(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}
