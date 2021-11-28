package log

import (
	"basic_framework/configs"
	"basic_framework/core/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var logs = logrus.New()
var nowDay string

// 加载文件
func loadFile() {

	// 当天
	nowDayT := time.Now().Format(configs.Yaml.Time.YMD)

	if nowDayT != nowDay {

		dir, _ := os.Getwd()

		// 输出日志文件地址
		outPutFile := file.OpenAppendFile(fmt.Sprintf("%s-%s.%s", configs.Yaml.Log.Prefix, time.Now().Format(configs.Yaml.Time.Y_M_D), configs.Yaml.Log.Suffix), path.Join(dir, configs.Yaml.Log.Path))

		if configs.Yaml.RunMode == gin.DebugMode || configs.Yaml.RunMode == gin.TestMode {

			// 设置日志最低等级
			logs.SetLevel(logrus.DebugLevel)

			// 输出位置 - 文件 & Stdout
			logs.SetOutput(io.MultiWriter(os.Stdout, outPutFile))
		} else {

			// 设置日志最低等级
			logs.SetLevel(logrus.ErrorLevel)

			// 输出位置 - 文件
			logs.SetOutput(outPutFile)
		}

		// 输出文件名和方法信息
		// logs.SetReportCaller(true)
		// 输出样式
		logs.SetFormatter(&logrus.JSONFormatter{TimestampFormat: configs.Yaml.Time.Y_M_D_H_I_S})

		// 标识
		nowDay = nowDayT
	}

}

// 设置产生日志的文件位置和行号
func setLocation() string {
	// skip如果是0，返回当前调用Caller函数的函数名、文件、程序指针PC，1是上一层函数，以此类推
	_, f, line, ok := runtime.Caller(2)
	// 日志输出格式
	logInfo := ""

	if ok {
		logInfo = fmt.Sprintf("[%s:%d]", filepath.Base(f), line)
	}

	return logInfo
}

func Trace(v ...interface{}) {
	loadFile()
	logs.Trace(v, setLocation())
}

func Debug(v ...interface{}) {
	loadFile()
	logs.Debug(v, setLocation())
}

func Info(v ...interface{}) {
	loadFile()
	logs.Info(v, setLocation())
}

func Print(v ...interface{}) {
	loadFile()
	logs.Print(v, setLocation())
}

func Warn(v ...interface{}) {
	loadFile()
	logs.Warn(v, setLocation())
}

func Warning(v ...interface{}) {
	loadFile()
	logs.Warning(v, setLocation())
}

func Error(v ...interface{}) {
	loadFile()
	logs.Error(v, setLocation())
}

func Fatal(v ...interface{}) {
	loadFile()
	logs.Fatal(v, setLocation())
}

func Panic(v ...interface{}) {
	loadFile()
	logs.Panic(v, setLocation())
}
