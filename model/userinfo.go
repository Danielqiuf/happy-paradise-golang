package login

import . "fmt"

type UserInfo struct {
	UserName      string
	Age           int16
	LastLoginTime int64
}

func init() {
	Println("userinfo initialize!!")
}
