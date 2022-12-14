/**
 * Created by PhpStorm.
 * @file   api_middleware.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/1 17:32
 * @desc   api_middleware.go
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/global"
	"github.com/go-utils-module/module/utils"
	"github.com/golang-module/carbon"
	"strconv"
)

type InnerApiMiddleware struct {
	BaseMiddleware
	serverConfig config.Server
	apiConfig    config.Api
}

func NewInnerApiMiddleware(serverConfig config.Server, apiConfig config.Api) *InnerApiMiddleware {
	middle := new(InnerApiMiddleware)
	middle.serverConfig = serverConfig
	middle.apiConfig = apiConfig
	return middle
}

func (a *InnerApiMiddleware) Middleware() gin.HandlerFunc {
	return a.checkSign
}

// 检查参数签名
func (a *InnerApiMiddleware) checkSign(context *gin.Context) {
	ts := context.Query("ts")
	sign := context.Query("sign")
	if ts == "" || sign == "" {
		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
	}
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		utils.ApiResponse(context, global.NoSignParamsErr)
		context.Abort()
	}
	path := a.getBaseUri(context, a.serverConfig.Domain)
	if a.isApi(path) { // 不是api 请求
		// 接口请求超时超过系统超时
		if carbon.Now().Timestamp()-timestamp > a.apiConfig.Overtime {
			utils.ApiResponse(context, global.RequestOvertimeErr)
			context.Abort()
		}
		newSign := utils.RequestSign(ts, a.apiConfig.Secret)
		if newSign != sign {
			utils.ApiResponse(context, global.SignErr)
			context.Abort()
		}
	}
}
