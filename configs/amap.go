package configs

// 高德地图
type amap struct {

	// Web服务端
	WebServerKey string `yaml:"WebServerKey"`

	// 逆地理编码API服务地址
	HostRegeo string `yaml:"HostRegeo"`

	// 坐标转换
	HostCoordinate string `yaml:"HostCoordinate"`
}
