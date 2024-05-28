package controllers

import (
	"github.com/gin-gonic/gin"
	"learn-go/helpers"
	"learn-go/modules/auth/requests"
	"learn-go/modules/auth/services"
	"net/http"
)

type AuthController struct {
	Service *services.Auth
}

func NewAuthController(auth *services.Auth) *AuthController {
	return &AuthController{Service: auth}
}

func (us *AuthController) LoginForm(c *gin.Context) {
	errorMsg, _ := helpers.GetSession(c, "error", true)
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"error": errorMsg,
		"old":   helpers.GetOldData(c),
	})
}

func (us *AuthController) Login(c *gin.Context) {
	var form requests.LoginForm
	if form.Validate(c) {
		if us.Service.CheckAuth(c, form.Username, form.Password) {
			c.Redirect(http.StatusFound, "/home")
		} else {
			helpers.SetOldData(c, form)
			helpers.SetSession(c, "error", "Tài khoản hoặc mật khẩu không đúng")
			referer := c.Request.Referer()
			c.Redirect(http.StatusFound, referer)
		}
	}
}

func (us *AuthController) Logout(c *gin.Context) {
	helpers.RemoveSession(c, "auth_user")
	c.Redirect(http.StatusFound, "/auth/login")
}

func (us *AuthController) RegisterFrom(c *gin.Context) {
	errorMsg, _ := helpers.GetSession(c, "error", true)
	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"error": errorMsg,
		"old":   helpers.GetOldData(c),
	})
}

func (us *AuthController) Register(c *gin.Context) {
	var form requests.RegisterForm
	if form.Validate(c) {
		if !us.Service.CheckExists(form.Username) {
			if us.Service.RegisterUser(form.Username, form.Password) {
				c.Redirect(http.StatusFound, "/auth/login")
				return
			}
		}

		helpers.SetOldData(c, form)
		helpers.SetSession(c, "error", "Tài khoản đã tồn tại")
		referer := c.Request.Referer()
		c.Redirect(http.StatusFound, referer)
	}
}
