package requests

import (
	"github.com/gin-gonic/gin"
	"learn-go/helpers"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (form *LoginForm) Validate(c *gin.Context) bool {
	if err := c.ShouldBind(&form); err != nil {
		helpers.SetOldData(c, form)
		helpers.SetSession(c, "error", err.Error())
		referer := c.Request.Referer()
		c.Redirect(http.StatusFound, referer)

		return false
	}

	return true
}
