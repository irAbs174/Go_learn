package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"id": 174,
			"name": "abbas",
			"family": "damerchilo",
			"skils": gin.H{
				"0": "Full Stack developer",
				"1": "Network",
				"2": "Ecommerce",
				"3": "Cyber SEC",
			},
		})
	})
	r.GET("/bye", func(c *gin.Context){
		c.String(200, "Bye !!")
	})
	r.Run(":2020")
}