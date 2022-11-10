/**
 * Created by Goland.
 * @file   display.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 16:17
 * @desc   display.go
 */

package utils

import (
	"fmt"
	"github.com/go-utils-module/module/utils/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 响应结构体
type ResponseData struct {
	Code fmt.Stringer `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

// ApiResponse 异常通知
func ApiResponse(context *gin.Context, errorCode fmt.Stringer, data ...interface{}) {
	response := ResponseData{
		Code: errorCode,
		Msg:  errorCode.String(),
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	context.JSON(http.StatusOK, response)
}

// WebResponse 异常通知
func WebResponse(context *gin.Context, errorCode interface{}, data ...interface{}) {
	msg := ""
	var errCode code.ErrCode = 201
	if code, ok := errorCode.(code.ErrCode); ok {
		msg = code.String()
		errCode = code
	} else if err, ok := errorCode.(error); ok {
		msg = err.Error()
	} else {
		msg = errorCode.(string)
	}
	response := ResponseData{
		Code: errCode,
		Msg:  msg,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	context.HTML(http.StatusOK, "notice.html", gin.H{"notice": response})
}
