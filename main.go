package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
	"net/http"
	"tonic/generator"
)

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	userRoute := r.Group("/image")
	{
		userRoute.GET("/circles", func(c *gin.Context) {
			file := generator.DrawOne("circles")
			c.Header("Content-Type", "image/png")
			c.File(file)
		})
		userRoute.GET("/:generator", func(c *gin.Context) {
			engine := c.Param("generator")
			file := generator.DrawOne(engine)
			c.Header("Content-Type", "image/png")
			c.File(file)
		})

	}

	listRoute := r.Group("/list")
	{
		listRoute.GET("/simple", func(c *gin.Context) {
			c.HTML(http.StatusOK, "simple.tmpl", gin.H{
				"keys": maps.Keys(generator.DRAWINGS),
			})
		})

		return r
	}
}

func main() {
	router().Run()
}
