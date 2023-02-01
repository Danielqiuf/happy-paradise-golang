package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy-paradise-golang/model"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("-----middleware")
		status := c.Writer.Status()
		fmt.Println("status:", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()

	r.Use(MiddleWare())

	d := login.UserInfo{UserName: "danielqiu", Age: 12, LastLoginTime: 123332}

	r.GET("/login", func(c *gin.Context) {
		// 取值
		phone := c.DefaultQuery("phone", "100")
		c.JSON(200, gin.H{"message": phone, "userName": d.UserName, "age": d.Age})
	})

	r.Run(":8000")
}
