package common

import (
	"basic_framework/configs"
	"basic_framework/core/log"
	"fmt"
	"time"
)

type BaseService struct {
}

// 开始执行日志
func (this BaseService) ExeStartTimeLog(title string, mark interface{}, startTime int64) {
	st := time.Now().UnixNano() / 1e6
	log.Info(fmt.Sprintf("%v %v 执行开始时间: 毫秒=%v, 格式时间=%v", title, mark, startTime, time.Unix(st/1000, 0).Format(configs.Yaml.Time.Y_M_D_H_I_S)))
}

// 开始执行日志
func (this BaseService) ExeEndTimeLog(title string, mark interface{}, startTime int64) {
	et := time.Now().UnixNano() / 1e6
	log.Info(fmt.Sprintf("%v %v 执行结束时间: 毫秒=%v, 格式时间=%v，执行总毫秒数：%v", title, mark, et, time.Unix(et/1000, 0).Format(configs.Yaml.Time.Y_M_D_H_I_S), et-startTime))
}
