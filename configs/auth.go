package configs

// 限制
type auth struct {

	// 是否开启
	Enable bool `yaml:"Enable"`

	// 允许访问ip
	Ip []string `yaml:"Ip"`
}
