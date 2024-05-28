package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectIfAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("auth_user") //ve sau chac chan hon query db voi thong tin user trong session
		if user != nil {
			c.Redirect(http.StatusFound, "/home")
			return
		}

		c.Next()
	}
}
