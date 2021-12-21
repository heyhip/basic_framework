package configs

// 数据库
type databases struct {

	// 数据库
	Db1 dbConfig `yaml:"Db1"`
}

type dbConfig struct {

	// 类型
	Type string `yaml:"Type"`

	// 地址
	Host string `yaml:"Host"`

	// 用户名
	Username string `yaml:"Username"`

	// 密码
	Password string `yaml:"Password"`

	// 端口
	Port string `yaml:"Port"`

	// 数据库名称
	Database string `yaml:"Database"`

	// 前缀
	Prefix string `yaml:"Prefix"`

	// 最大空闲连接数
	MaxIdleConns int `yaml:"MaxIdleConns"`

	// 最大连接数量
	MaxOpenConns int `yaml:"MaxOpenConns"`

	// 可复用最大时间
	MaxLifetime int `yaml:"MaxLifetime"`

	// 慢sql时间
	SlowThreshold int `yaml:"SlowThreshold"`

	// 日志等级 1,2,3,4
	LogLevel int `yaml:"LogLevel"`
}
