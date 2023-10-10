package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const bulkMin = 1
const bulkMax = 20000

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()
	router.Run("0.0.0.0:8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", get)
	return router
}

func get(c *gin.Context) {
	quantityParam := c.DefaultQuery("quantity", "1")
	quantity, err := strconv.Atoi(quantityParam)
	if err != nil {
		c.String(400, "Invalid quantity "+quantityParam)
		return
	}
	if (quantity < bulkMin) || (quantity > bulkMax) {
		c.String(400, fmt.Sprintf("Quantity must be between %d and %d", bulkMin, bulkMax))
		return
	}
	uuids := make([]string, quantity)
	for i := 0; i < quantity; i++ {
		uuids[i] = NewUUID()
	}
	c.IndentedJSON(200, uuids)
}
