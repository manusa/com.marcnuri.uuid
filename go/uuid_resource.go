package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
)

const bulkMin = 1
const bulkMax = 20000

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
			os.Exit(1)
		}
	}()
	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if len(allowedOrigins) != 0 {
		config := cors.DefaultConfig()
		config.AllowOrigins = strings.Split(allowedOrigins, ",")
		router.Use(cors.New(config))
	}
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
