package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"learn-go/middleware"
	authController "learn-go/modules/auth/controllers"
	homeController "learn-go/modules/home/controllers"
)

type Router struct {
	AuthController *authController.AuthController
	HomeController *homeController.HomeController
}

func NewRouter(auth *authController.AuthController, home *homeController.HomeController) *Router {
	return &Router{
		HomeController: home,
		AuthController: auth,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("quyennv-secret"))
	router.Use(sessions.Sessions("quyennv", store))

	router.LoadHTMLGlob("modules/**/resources/**/*.html")
	router.Static("/static", "./static")

	group := router.Group("/").Use(middleware.AuthMiddleware())
	{
		group.GET("/home", r.HomeController.Home)
		group.GET("/auth/logout", r.AuthController.Logout)
	}

	authGroup := router.Group("/").Use(middleware.RedirectIfAuthenticatedMiddleware())
	{
		authGroup.GET("/auth/login", r.AuthController.LoginForm)
		authGroup.POST("/auth/login", r.AuthController.Login)
		authGroup.GET("/auth/register", r.AuthController.RegisterFrom)
		authGroup.POST("/auth/register", r.AuthController.Register)
	}

	return router
}
