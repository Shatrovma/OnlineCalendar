package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/event", CreateEvent)
	//router.GET(("/event/:id"))
	//router.PUT("/event/:id")
	//router.DELETE("/event/:id")

	router.Run(":8084")
}

func CreateEvent(c *gin.Context) {
	var event Event
	if err := c.BindJSON(&event); err != nil {
		fmt.Printf("Не могу распознать объект Event", err.Error())
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	c.JSON(200, "Удачно получили объект с id")
}
