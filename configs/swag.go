package configs

// swaggo文档通用的API注释
type swag struct {
	// 必填 应用程序的名称
	Title string `yaml:"Title"`

	// 描述
	Description string `yaml:"Description"`

	// 必填 提供应用程序API的版本
	Version string `yaml:"Version"`

	// 运行API的主机
	Host string `yaml:"Host"`

	// 运行API的基本路径
	BasePath string `yaml:"BasePath"`

	// 请求的传输协议
	Schemes []string `yaml:"Schemes"`
}
