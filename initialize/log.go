/**
 * Created by PhpStorm.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 01:18
 * @desc   log.go
 */

package system

import (
	"github.com/druidcaesa/gotool"
	"github.com/druidcaesa/gotool/openfile"
	"github.com/go-utils-module/module/config"
	"github.com/go-utils-module/module/utils"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

// InitializeLogger 初始化日志配置
func InitializeLogger(config config.Log) *logrus.Logger {
	if !gotool.FileUtils.Exists(config.Path) {
		err := os.MkdirAll(config.Path, os.ModePerm)
		if err != nil {
			panic("init system error. make log data err.path:" + config.Path)
		}
	}
	// 日志文件
	fileName := path.Join(config.Path, config.File)
	if !gotool.FileUtils.Exists(fileName) {
		openfile.Create(fileName)
		if !gotool.FileUtils.Exists(fileName) {
			panic("init system error. create log file err. log file:" + fileName)
		}
	}
	logger := new(utils.Log).InitLogger(config.Path, config.File, config.Mode)
	return logger
}
