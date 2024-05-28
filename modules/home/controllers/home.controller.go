package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (us *HomeController) Home(c *gin.Context) {
	data := c.Keys["templateData"].(gin.H)
	data["icon"] = "(^_^)"
	c.HTML(http.StatusOK, "home/home.html", data)
}
