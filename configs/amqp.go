package configs

// 消息队列
type amqp struct {

	// 气象队列
	Weather weather `yaml:"Weather"`
}

// 气象队列
type weather struct {

	// 是否开启
	Enable bool `yaml:"Enable"`

	// 地址
	Address string `yaml:"Address"`

	// 账号key
	AccessKey string `yaml:"AccessKey"`

	// 账号秘钥
	AccessSecret string `yaml:"AccessSecret"`

	// 消费组ID
	ConsumerGroupId string `yaml:"ConsumerGroupId"`

	// 客户端ID，唯一标识
	ClientId string `yaml:"ClientId"`

	// 实例ID
	IotInstanceId string `yaml:"IotInstanceId"`

	// 队列名称
	QueueName string `yaml:"QueueName"`
}
