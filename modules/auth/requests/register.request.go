package requests

import (
	"github.com/gin-gonic/gin"
	"learn-go/helpers"
	"net/http"
)

type RegisterForm struct {
	Username        string `form:"username" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
}

func (form *RegisterForm) Validate(c *gin.Context) bool {
	if err := c.ShouldBind(&form); err != nil {
		helpers.SetOldData(c, form)
		helpers.SetSession(c, "error", err.Error())
		referer := c.Request.Referer()
		c.Redirect(http.StatusFound, referer)
		return false
	}
	return true
}
