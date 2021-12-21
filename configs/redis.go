package configs

// 缓存
type redis struct {

	// 数据库
	Db1 redisConfig `yaml:"Db1"`
}

type redisConfig struct {
	// 地址
	Host string `yaml:"Host"`

	// 端口
	Port string `yaml:"Port"`

	// 密码
	Auth string `yaml:"Auth"`

	// 数据库
	Select int `yaml:"Select"`

	// 连接方式
	Network string `yaml:"Network"`

	// 最大连接数
	PoolSize int `yaml:"PoolSize"`
}
