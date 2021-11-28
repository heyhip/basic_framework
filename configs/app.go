package configs

import "time"

// 应用基础配置
type app struct {

	// 地址
	Host string `yaml:"Host"`

	// 端口
	Port string `yaml:"Port"`

	// 读超时时间，秒
	ReadTimeout time.Duration `yaml:"ReadTimeout"`

	// 写超时时间，秒
	WriteTimeout time.Duration `yaml:"WriteTimeout"`

	// 请求最大字节数，0为默认DefaultMaxHeaderBytes大小1MB
	MaxHeaderBytes int `yaml:"MaxHeaderBytes"`

	// pprof性能分析工具地址
	Pprof string `yaml:"Pprof"`

	// 静态文件地址
	StaticsUrl string `yaml:"StaticsUrl"`
}
