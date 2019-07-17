package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	_ = router.GET("/notify", func(c *gin.Context) {
		rawData, e := c.GetRawData()
		if e == nil {
			log.Fatal(e)
		}
		log.Println(fmt.Sprintf("notify---body: %v", rawData, ))
	})
	_ = router.Run(":9002")
}