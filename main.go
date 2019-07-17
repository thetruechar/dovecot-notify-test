package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.GET("/notify", func(c *gin.Context) {
		log.Println(fmt.Sprintf("notify---body: %v", c.GetRawData()))
	})
	_ = router.Run(":9002")
}