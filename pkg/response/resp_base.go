package response

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/piupuer/go-helper/pkg/utils"
	"happy-paradise-golang/pkg/global"
)

type Base struct {
	Id uint `json:"id"'`
	Time
}

type Time struct {
	CreatedAt carbon.DateTime `json:"createdAt" swaggertype:"string" example:"2019-01-01 00:00:00"` // create time
	UpdatedAt carbon.DateTime `json:"updatedAt" swaggertype:"string" example:"2019-01-01 00:00:00"` // update time
}

type Response struct {
	Code      int         `json:"code" enums:"200,400"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	RequestId string      `json:"requestId""`
}

type Pageable struct {
	Page
	Source interface{} `json:"source"`
}

type Page struct {
	CurrentPage uint  `json:"currentPage"`
	PageSize    uint  `json:"pageSize"`
	Total       int64 `json:"total"`
}

type PageData struct {
	Page
	List interface{} `json:"list"`
}

func GetSuccess() Response {
	return GetResult(global.Ok, map[string]interface{}{}, global.ErrorHandle[global.Ok])
}

func GetResult(code int, data interface{}, format interface{}, a ...interface{}) Response {
	var f string
	switch format.(type) {
	case string:
		f = format.(string)
	case error:
		f = fmt.Sprintf("%v", format.(error))
	}
	return Response{
		Code: code,
		Data: data,
		Msg:  fmt.Sprintf(f, a...),
	}
}

func GetSuccessWithData(data ...interface{}) Response {
	switch len(data) {
	case 1:
		return GetResult(global.Ok, data[0], global.ErrorHandle[global.Ok])
	case 2:
		utils.Struct2StructByJson(data[0], data[1])
		return GetResult(global.Ok, data[1], global.ErrorHandle[global.Ok])
	}
	return GetSuccess()
}

func SuccessData(data ...interface{}) {
	panic(GetSuccessWithData(data...))
}

func GetSuccessWithPageData(real, brief interface{}, page Page) Response {
	utils.Struct2StructByJson(real, brief)
	var rp PageData
	rp.Page = page
	rp.List = brief
	return GetResult(global.Ok, rp, global.ErrorHandle[global.Ok])
}

func GetFailWithMsg(format interface{}, a ...interface{}) Response {
	return GetResult(global.Error, map[string]interface{}{}, format, a...)
}

func GetFailWithCode(code int) Response {
	// default NotOk
	msg := global.ErrorHandle[global.Error]
	if val, ok := global.ErrorHandle[code]; ok {
		msg = val
	}
	return GetResult(code, map[string]interface{}{}, msg)
}

func GetFailWithCodeAndMsg(code int, format interface{}, a ...interface{}) Response {
	return GetResult(code, map[string]interface{}{}, format, a...)
}
