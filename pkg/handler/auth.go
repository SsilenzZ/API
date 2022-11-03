package handler

import (
	"Api/pkg/db/repository/user"
	"Api/pkg/requests"
	"Api/pkg/service"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthHandler struct {
	Hasher service.HasherI
	Jwt    service.JwtI
}

func NewAuthHandler(hasher service.HasherI, jwt service.JwtI) AuthHandler {
	return AuthHandler{Hasher: hasher, Jwt: jwt}
}

// @Summary Register
// @Description Register an account with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.Users true "Data for user to create"
// @Success 200 "OK"
// @Failure 400 "Account_with_this_email_is_already_registred"
// @Router /signup [post]
func (h AuthHandler) SignUp(c echo.Context) error {
	var user_ requests.SignUp
	err := c.Bind(&user_)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	password, err := h.Hasher.HashPassword(user_.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	success := user.Repository.CreateUser(user_.Email, password, user_.Name)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// @Summary Sign in
// @Description Sign in with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.Users true "Data for user to login"
// @Success 200 "OK"
// @Failure 400 "Wrong_login_info"
// @Router /signin [post]
func (h AuthHandler) SignIn(c echo.Context) error {
	var user_ requests.SignIn
	err := c.Bind(&user_)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	id, err := user.Repository.GetHashedPassword(user_.Email, user_.Password)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	t := h.Jwt.GenerateAccessToken(id)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// @Summary Sign in throгgh google
// @Description Get google authorization link
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /signin/google [get]
func (h AuthHandler) GoogleSignIn(c echo.Context) error {
	result := service.GoogleOauthConfig.AuthCodeURL(service.State)
	return c.JSON(http.StatusOK, result)
}

func (h AuthHandler) GetAuthToken(c echo.Context) error {
	if c.FormValue("state") != service.State {
		return c.String(http.StatusOK, "state is not valid")
	}

	var (
		response *http.Response
		token    *oauth2.Token
		err      error
	)

	token, err = service.GoogleOauthConfig.Exchange(context.Background(), c.FormValue("code"))
	if err != nil {
		return c.String(http.StatusOK, "could not get the token")
	}
	response, err = http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.String(http.StatusOK, "failed getting user info")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return c.String(http.StatusOK, "failed reading response body")
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(content, &data)
	if err != nil {
		return c.String(http.StatusOK, "failed unmarshalling response body")
	}
	password, err := h.Hasher.HashPassword("password")
	if err != nil {
		return c.String(http.StatusOK, "failed hashing password")
	}
	u := user.Repository.GetUser(data["email"].(string), password)

	t := h.Jwt.GenerateAccessToken(u.ID)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
