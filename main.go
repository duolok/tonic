package main

import (
    "github.com/gin-gonic/gin"
    "tonic/generator"
)

func router() *gin.Engine {
    r := gin.Default()
    userRoute := r.Group("/image")
    {
        userRoute.GET("/circles", func(c *gin.Context) {
            file := generator.DrawOne("circles")
            c.Header("Content-Type", "image/png")
            c.File(file)
        })
    }

    return r
}

func main() {
    router().Run()
}
