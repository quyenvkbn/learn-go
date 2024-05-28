package services

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"learn-go/helpers"
	"learn-go/modules/users/models"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Auth struct {
	Model *models.User
}

func NewAuthService(model *models.User) *Auth {
	return &Auth{Model: model}
}

func (sv *Auth) CheckAuth(c *gin.Context, username string, password string) bool {
	user := sv.Model.GetUser(username, true)
	if user["password"] != nil && helpers.CompareHashAndPassword(helpers.ToString(user["password"]), password) {
		delete(user, "password")
		helpers.SetSession(c, "auth_user", helpers.ToJsonString(user))

		return true
	}
	return false
}

func (sv *Auth) CheckExists(username string) bool {
	if sv.Model.GetUser(username, false) != nil {

		return true
	}

	return false
}

func (sv *Auth) RegisterUser(username string, password string) bool {
	if sv.Model.CreateUser(username, password) {

		return true
	}

	return false
}
