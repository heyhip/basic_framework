package embeds

import (
	"basic_framework/configs"
	_ "embed"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 配置文件
//go:embed config_pro.yaml
var releaseConfig []byte

// 初始化配置
func init() {

	// 解析配置文件
	var appBasic configs.AppBasic

	// 命令行启动参数 mode=release
	if len(os.Args) > 1 {
		modeS := strings.Split(os.Args[1], "=")
		appBasic.RunMode = modeS[1]
	}

	var data []byte
	var err error

	dir, _ := os.Getwd()

	if appBasic.RunMode == "" {
		// 读取总配置文件
		data, err = ioutil.ReadFile(path.Join(dir, "embeds/config.yaml"))
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(data, &appBasic)
	}

	// 读取多环境配置文件
	if appBasic.RunMode == gin.ReleaseMode {
		//data, err = ioutil.ReadFile(path.Join(dir, "/config_pro.yaml"))
		data = releaseConfig
	} else if appBasic.RunMode == gin.DebugMode {
		data, err = ioutil.ReadFile(path.Join(dir, "embeds/config_dev.yaml"))
	} else if appBasic.RunMode == gin.TestMode {
		data, err = ioutil.ReadFile(path.Join(dir, "embeds/config_test.yaml"))
	} else {
		panic(errors.New("配置错误"))
	}

	// 解析配置文件
	err = yaml.Unmarshal(data, &configs.Yaml)

	if err != nil {
		panic(err)
	}

	configs.Yaml.AppBasic = appBasic
}
