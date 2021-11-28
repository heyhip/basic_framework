package configs

// 时间格式化
type timeFomat struct {
	Y_M_D_H_I_S string `yaml:"Y_M_D_H_I_S"`
	Y_M_D       string `yaml:"Y_M_D"`
	YMD         string `yaml:"YMD"`
	YM          string `yaml:"YM"`
}
