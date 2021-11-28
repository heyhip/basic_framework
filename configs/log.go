package configs

// 日志配置
type log struct {

	// 日志前缀
	Prefix string `yaml:"Prefix"`

	// 日志路径
	Path string `yaml:"Path"`

	// 日志文件后缀
	Suffix string `yaml:"Suffix"`

	// 日志等级 "DEBUG", "INFO", "WARN", "ERROR", "FATAL"
	Level string `yaml:"Level"`
}
