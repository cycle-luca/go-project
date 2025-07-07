package routes

import (
	"book-management-sqlite/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 静态文件和模板
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// 图书路由
	bookGroup := r.Group("/books")
	{
		bookGroup.GET("", handlers.GetBooks)
		bookGroup.POST("", handlers.CreateBook)
	}

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "图书管理系统"})
	})

	return r
}
