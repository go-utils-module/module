/**
 * Created by PhpStorm.
 * @file   aaa.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 00:44
 * @desc   aaa.go
 */

package server

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/module/utils"
)

type webHockFun func(engine *gin.Engine, template embed.FS, assets embed.FS)
type hockFun func(engine *gin.Engine)

type ServerConfig struct {
	Mode     string   `json:"mode"`
	Ip       string   `json:"ip"`
	Port     int      `json:"port"`
	Template embed.FS `json:"template"`
	Assets   embed.FS `json:"assets"`
}

// StartWebServer 启动web服务
func StartWebServer(config ServerConfig, hock webHockFun) error {
	//  创建路由
	engine := gin.Default()

	// 服务钩子
	hock(engine, config.Template, config.Assets)
	// 设置系统运行模式
	SetRunModel(config.Mode)
	// 开始启动服务
	addr := fmt.Sprintf("%s:%d", config.Ip, config.Port)
	err := engine.Run(addr)
	if utils.CheckErr(err) {
		utils.Logger.Fatalln("服务启动失败,错误信息:", err.Error())
		return err
	} else {
		utils.Logger.Info("服务启动成功. 服务地址:", addr)
	}
	return nil
}

// StartServer 启动后台服务
func StartServer(config ServerConfig, hock hockFun) error {
	//  创建路由
	engine := gin.Default()
	// 服务钩子
	hock(engine)
	// 设置系统运行模式
	SetRunModel(config.Mode)
	// 开始启动服务
	addr := fmt.Sprintf("%s:%d", config.Ip, config.Port)
	err := engine.Run(addr)
	if utils.CheckErr(err) {
		utils.Logger.Fatalln("服务启动失败,错误信息:", err.Error())
		return err
	} else {
		utils.Logger.Info("服务启动成功. 服务地址:", addr)
	}
	return nil
}

// SetRunModel 设置当前系统运行模式
func SetRunModel(mode string) {
	utils.Logger.Info("当前运行模式为:", mode)
	switch mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	}
}