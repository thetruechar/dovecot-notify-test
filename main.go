package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	_ = router.POST("/notify", func(c *gin.Context) {
		log.Println("POST")
		rawData, e := c.GetRawData()
		if e == nil {
			log.Println("err")
		}
		log.Println(fmt.Sprintf("notify---body: %s", rawData, ))
	})
	_ = router.Run(":9002")
}