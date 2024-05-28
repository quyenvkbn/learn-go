package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"learn-go/helpers"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("auth_user") //ve sau chac chan hon query db voi thong tin user trong session
		if user == nil {
			helpers.SetSession(c, "error", "Authentication required")
			c.Redirect(http.StatusFound, "/auth/login")
			return
		}

		c.Keys["templateData"] = gin.H{
			"user": helpers.StringToMap(helpers.ToString(user)),
		}

		c.Next()
	}
}
