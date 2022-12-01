/**
 * Created by PhpStorm.
 * @file   system.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 02:10
 * @desc   system.go
 */

package config

// System 系统配置
type System struct {
	Mode        string `yaml:"mode"`
	TempDir     string `yaml:"tempDir"`
	PwdPrefix   string `yaml:"pwdPrefix"`
	SystemTitle string `yaml:"systemTitle"`
}

// Upload 图片设置
type Upload struct {
	UploadDir string `yaml:"uploadDir"`
	TempDir   string `yaml:"tempDir"`
}

// Log 日志设置
type Log struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
	Mode string `yaml:"mode"`
}

// Server 服务配置
type Server struct {
	Protocol string `yaml:"protocol"`
	Domain   string `yaml:"domain"`
	Port     int    `yaml:"port"`
	Ip       string `yaml:"ip"`
}
type Avatar struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type View struct {
	PageCount int `yaml:"pageCount"`
}