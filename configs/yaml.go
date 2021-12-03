package configs

var Yaml appYaml

// 基础配置
type AppBasic struct {
	// gin运行模式 debug test release
	RunMode string `yaml:"RunMode"`
}

type appYaml struct {

	// 基础配置
	AppBasic

	// 应用基础配置
	App app `yaml:"App"`

	// 限制
	Auth auth `yaml:"Auth"`

	// 时间戳格式化
	Time timeFomat `yaml:"TimeFomat"`

	// 日志
	Log log `yaml:"Log"`

	// 数据库
	Db db `yaml:"Db"`

	// 缓存
	Redis redis `yaml:"Redis"`

	// 文档
	Swag swag `yaml:"Swag"`

	// 上传
	Upload upload `yaml:"Upload"`

	// 消息队列
	Amqp amqp `yaml:"Amqp"`

	// 高德
	Amap amap `yaml:"Amap"`

	// Jwt
	Jwt jwt `yaml:"Jwt"`

	// Wxpay 微信支付
	Wxpay wxpay `yaml:"Wxpay"`
}
