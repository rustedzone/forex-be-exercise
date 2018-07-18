package route

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	rate "forex-be-exercise/rate/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
	})
	r.Use(AccessControl())
	r.Static("/css", "./templates/assets/css")
	r.Static("/js", "./templates/assets/js")
	r.Static("/fonts", "./templates/assets/fonts")
	r.Static("/img", "./templates/assets/img")

	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLFiles("templates/views/index.html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/pages/:folder/:filename", func(c *gin.Context) {

		r.LoadHTMLFiles("templates/views/pages/" + c.Param("folder") + "/" + c.Param("filename"))
		c.HTML(http.StatusOK, c.Param("filename"), gin.H{
			"title": "This Is Dashboard",
		})
	})

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": true, "data": "Forex Be Exercise's Endpoint"})
	})

	list := r.Group("api/list")
	{
		list.POST("", rate.NewList_)
		list.DELETE("", rate.RemoveFromList_)
	}

	daily := r.Group("api/daily")
	{
		daily.POST("", rate.NewTx_)
		daily.GET("", rate.TrackDailyTx_)
	}

	return r
}
