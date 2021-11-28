package configs

// 上传
type upload struct {

	// 位置
	Cloud string `yaml:"Cloud"`

	// 路径
	Path string `yaml:"Path"`

	// 类型
	Type []string `yaml:"Type"`

	// 大小 MB
	MaxSize float64 `yaml:"MaxSize"`
}
