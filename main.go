package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mwdchang/uhack2018/parse"
)

func main() {
	r := gin.Default()

	// set up CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081"}
	r.Use(cors.New(config))

	r.GET("/facilities", facilitiesHandler)
	r.GET("/water", waterHandler)
	r.Run(":3333")

}

func facilitiesHandler(c *gin.Context) {
	c.JSON(200, parse.Pollutants())
}

func waterHandler(c *gin.Context) {
	c.JSON(200, parse.TestSites())
}
