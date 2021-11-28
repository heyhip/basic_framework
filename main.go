package main

import (
	"basic_framework/configs"
	"basic_framework/core/log"
	"basic_framework/docs"
	_ "basic_framework/embeds"
	"basic_framework/middlewares"
	"basic_framework/routes"
	"basic_framework/scripts"
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"html/template"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 模板
//go:embed assets/views
var tmpl embed.FS

// 静态文件
//go:embed assets/statics
var statics embed.FS

func main() {

	log.Info("开始启动...")

	// 设置运行模式
	gin.SetMode(configs.Yaml.RunMode)

	r := gin.New()

	//// 加载模板路径，两层
	tp, _ := template.ParseFS(tmpl, "assets/views/**/*")
	r.SetHTMLTemplate(tp)
	//r.LoadHTMLGlob("./assets/views/**/*")

	//// 加载静态文件
	r.StaticFS("/statics", http.FS(statics))
	//r.StaticFile("/favicon.ico", "statics/assets/statics/image/favicon.ico")
	//r.StaticFS("/statics", http.Dir("./assets/statics"))

	// 加载路由
	routes.InitRoutes(r)

	// 设置中间件，恐慌时恢复
	r.Use(gin.Recovery())

	// 自定义基础中间件
	r.Use(middlewares.Basic())

	if configs.Yaml.RunMode != gin.ReleaseMode {
		// 控制台输出请求日志
		r.Use(gin.Logger())

		// 调试工具
		go func() {
			http.ListenAndServe(configs.Yaml.App.Pprof, nil)
		}()

		// swagger动态配置
		sw := configs.Yaml.Swag
		docs.SwaggerInfo.Title = sw.Title
		docs.SwaggerInfo.Description = sw.Description
		docs.SwaggerInfo.Version = sw.Version
		docs.SwaggerInfo.Host = sw.Host
		docs.SwaggerInfo.BasePath = sw.BasePath
		docs.SwaggerInfo.Schemes = sw.Schemes

		// 文档
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", configs.Yaml.App.Host, configs.Yaml.App.Port),
		Handler:      r,
		ReadTimeout:  configs.Yaml.App.ReadTimeout * time.Second,
		WriteTimeout: configs.Yaml.App.WriteTimeout * time.Second,
		// HTTP请求的头域最大允许长度,1MB
		MaxHeaderBytes: configs.Yaml.App.MaxHeaderBytes,
	}

	// 使用协程开始监听
	go func() {
		if err := s.ListenAndServe(); err != nil {
			//panic("监听端口失败")
			log.Error("监听端口失败: "+time.Now().Format(configs.Yaml.Time.Y_M_D_H_I_S), err)
		} else {
			log.Info("启动成功: " + time.Now().Format(configs.Yaml.Time.Y_M_D_H_I_S))
		}
	}()

	// 定时任务
	scripts.Execute()

	// 创建一个通道，类型为信号接口
	quit := make(chan os.Signal)
	// notify方法用来监听收到的信号，将监听的信号放入通道
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// 获取通道信号，如果没有，会阻塞
	<-quit

	// 关机日志
	log.Info("开始关机..." + time.Now().Format(configs.Yaml.Time.Y_M_D_H_I_S))

	/**
	context.Background()返回一个空的上下文，这个空的Context一般用于整个Context树的根节点
	context.WithTimeout超时自动取消方法，传入上下文，设置5秒时间超时，然后取消协程
	返回上下文和一个可执行函数
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 返回的是一个函数，是用来释放关联的资源，提前主动调用，可以尽快的释放，避免等待过期时间之间的浪费，在调用WithTimeout之后defer cancel()
	defer cancel()

	// 使用Shutdown关机
	if err := s.Shutdown(ctx); err != nil {
		log.Error("关机失败: "+time.Now().Format(configs.Yaml.Time.Y_M_D_H_I_S), err)
	}

	log.Info("关机成功: " + time.Now().Format(configs.Yaml.Time.Y_M_D_H_I_S))

}
