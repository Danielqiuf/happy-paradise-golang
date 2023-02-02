package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	_ "happy-paradise-golang/model"
	"happy-paradise-golang/utils"
	"time"
)

type CpBuilder struct {
}

type Pouring struct {
	pouring []byte
	len     int
	cap     int
}

type CpStack struct {
	CpBuilder
	umi     string
	pouring utils.List[int]
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		Println("-----middleware")
		status := c.Writer.Status()
		Println("status:", status)
		t2 := time.Since(t)
		Println("time:", t2)
	}
}

func main() {
	var ch chan int
	/** pouring为新分配的内存地址，指向List结构体 */
	pouring := utils.MakeList[int](7, 10)

	/** 追加数组 **/
	pouring.Append(422)
	pouring.Append(566)

	cp := CpStack{
		umi:     "Umi",
		pouring: *pouring, // 使用指针访问地址
	}

	cp.pouring.Append(988)
	cp.pouring.Append(123)

	Println("cp --- ", cp)

	ch = make(chan int, 1)
	ch <- 100
	close(ch)

	func() {
		for {
			ret, ok := <-ch
			if !ok {
				Println("closed channel")
				break
			}
			Println("ccc go ==> ", ret, ok)
		}

	}()

	//r := gin.Default()
	//
	//r.Use(MiddleWare())
	//
	//d := login.UserInfo{UserName: "danielqiu", Age: 12, LastLoginTime: 123332}
	//
	//r.GET("/login", func(c *gin.Context) {
	//	// 取值
	//	phone := c.DefaultQuery("phone", "100")
	//	c.JSON(200, gin.H{"message": phone, "userName": d.UserName, "age": d.Age})
	//})
	//
	//r.Run(":8000")
}
