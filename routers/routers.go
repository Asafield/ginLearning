package routers

import (
	"ginLearning/config"
	"ginLearning/handler/users"
	"ginLearning/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//创建一个*gin.EnginE类型的对象
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Cors())
	gin.SetMode(config.RunMode)

	user := r.Group("/user")
	user.POST("/login", users.SignInHandler)
	return r
}
