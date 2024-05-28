package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, msgKey string, msg string) {
	session := sessions.Default(c)
	session.Set(msgKey, msg)
	session.Save()
}

func GetSession(c *gin.Context, msgKey string, isFlash bool) (string, bool) {
	session := sessions.Default(c)
	if msg := session.Get(msgKey); msg != nil {
		if isFlash {
			session.Delete(msgKey)
			session.Save()
		}

		return ToString(msg), true
	}

	return "", false
}

func RemoveSession(c *gin.Context, msgKey string) {
	session := sessions.Default(c)
	session.Delete(msgKey)
	session.Save()
}

func SetOldData(c *gin.Context, form any) {
	path := c.Request.URL.Path
	SetSession(c, path, ToJsonString(form))
}

func GetOldData(c *gin.Context) map[string]interface{} {
	path := c.Request.URL.Path
	jsonData, _ := GetSession(c, path, true)

	return StringToMap(jsonData)
}
