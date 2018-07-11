package main

import "github.com/gin-gonic/gin"

// MeasuredLocation is a measurment location with associated measurements
type MeasuredLocation struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Data Data    `json:"data"`
	Type string  `json:"type"`
}

// Data represents measurement data for a
type Data struct {
	Lead     []float64 `json:"lead"`
	Arsenic  []float64 `json:"arsenic"`
	Chromium []float64 `json:"chromium"`
}

func main() {
	r := gin.Default()
	r.GET("/facilities", facilitiesHandler)
	r.GET("/water", waterHandler)
	r.Run()

}

func facilitiesHandler(c *gin.Context) {
	c.JSON(200, Pollutants())
}

func waterHandler(c *gin.Context) {
	c.JSON(200, TestSites())
}
