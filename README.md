# basic_framework
### gin基础框架

### 运行
go run main.go mode=release

### 目录
```
asset 静态文件
configs 对应配置文件，用于读取
core 核心组件
docs swagger生成的文件
embeds 配置文件
middlewares 中间件
routes 路由
scripts 定时脚本
tools 第三方工具
logs 日志目录
web 项目
 |- common 公共函数，错误提示
 |- controllers 控制器
 |- models 数据模型
 |- request 请求数据模型
 |- response 返回数据模型
 |- services 服务层

build 编译目录，自行创建
``` 

#### Pprof 性能分析，默认非Release可用，可在main中打开限制
```
http://localhost:6060/debug/pprof/
```

#### Swagger 文档，默认非Release可用，可在main中打开限制
```
// 安装命令行
go get -u github.com/swaggo/swag/cmd/swag
// 根目录执行，生成api文档
swag init
// 查看d地址
http://localhost:8888/swagger/index.html
```