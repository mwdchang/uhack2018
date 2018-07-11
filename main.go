package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

// MeasuredLocation is a measurment location with associated measurements
type MeasuredLocation struct {
	Lat      float64   `json:"lat"`
	Lon      float64   `json:"lon"`
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Data     []float64 `json:"data"`
	Type     string    `json:"type"`
	Chemical string    `json:"chemical"`
}

func main() {
	r := gin.Default()

	// set up CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	r.Use(cors.New(config))

	r.GET("/facilities", facilitiesHandler)
	r.GET("/water", waterHandler)
	r.Run(":3333")

}

func facilitiesHandler(c *gin.Context) {
	c.JSON(200, Pollutants())
}

func waterHandler(c *gin.Context) {
	c.JSON(200, TestSites())
}
