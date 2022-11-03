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

// @Summary List comments
// @Description Get all comments
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Success 200 {array} comments.Comments
// @Router /comments [get]
func ReturnAllComments(c echo.Context) error {
	result := comment.Repository.GetAllComments()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}

	return c.XML(http.StatusOK, result)
}

// @Summary Comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Param id path int true "ID of comment to return"
// @Success 200 {object} comments.Comments
// @Failure 400 "Record_not_found"
// @Router /comments/{id} [get]
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

// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body comments.Comments true "Data for comment to create"
// @Success 200 "OK"
// @Failure 400 "Bad_request"
// @Router /restricted/comments [post]
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

// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of comment to update"
// @Param comment body comments.Comments true "Data for comment to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /comments/{id} [put]
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

// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of comment to update"
// @Param comment body comments.Comments true "Data for comment to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /comments/{id} [put]
func DeleteComment(c echo.Context) error {
	comment.Repository.DeleteComment(c.Param("id"))

	return c.JSON(http.StatusOK, nil)
}
