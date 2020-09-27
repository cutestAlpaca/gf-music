# 基于 GoFrame 框架
 https://goframe.org
 ****
 
 ##限制
 golang版本 >= 1.11
 
 ####代理地址使用
 https://goproxy.io
 
 ## 架构
 <div align=center>
 <img src="https://goframe.org/images/arch.png?v=12"/>
 </div>
 
 ## 项目结构
 
 ```undefined
 ├── app
 │   ├── api
 │   ├── model
 │   └── service
 ├── boot
 ├── config
 ├── docker
 ├── document
 ├── i18n
 ├── library
 ├── public
 ├── router
 ├── template
 ├── vendor
 ├── Dockerfile
 ├── go.mod
 └── main.go
 ```
 #
| 目录/文件名称 | 说明       | 描述                                                         |
| :------------ | :--------- | :----------------------------------------------------------- |
| `app`         | 业务逻辑层 | 所有的业务逻辑存放目录。                                     |
| - `api`       | 业务接口   | 接收/解析用户输入参数的入口/接口层。                         |
| - `model`     | 数据模型   | 数据管理层，仅用于操作管理数据，如数据库操作。               |
| - `service`   | 逻辑封装   | 业务逻辑封装层，实现特定的业务需求，可供不同的包调用。       |
| `boot`        | 初始化包   | 用于项目初始化参数设置，往往作为`main.go`中第一个被`import`的包。 |
| `config`      | 配置管理   | 所有的配置文件存放目录。                                     |
| `docker`      | 镜像文件   | Docker镜像相关依赖文件，脚本文件等等。                       |
| `document`    | 项目文档   | Document项目文档，如: 设计文档、帮助文档等等。               |
| `i18n`        | I18N国际化 | I18N国际化配置文件目录。                                     |
| `library`     | 公共库包   | 公共的功能封装包，往往不包含业务需求实现。                   |
| `public`      | 静态目录   | 仅有该目录下的文件才能对外提供静态服务访问。                 |
| `router`      | 路由注册   | 用于路由统一的注册管理。                                     |
| `template`    | 模板文件   | MVC模板文件存放的目录。                                      |
| `vendor`      | 第三方包   | 第三方依赖包存放目录(可选, 未来会被淘汰)。                   |
| `Dockerfile`  | 镜像描述   | 云原生时代用于编译生成Docker镜像的描述文件。                 |
| `go.mod`      | 依赖管理   | 使用`Go Module`包管理的依赖描述文件。                        |
| `main.go`     | 入口文件   | 程序入口文件。                                               |

## 初始化项目init

## 热编译运行项目run
gf run <项目/主文件>  

gf run main.go

## 交叉编译build
gf build FILE [OPTION]

工具链  
https://goframe.org/toolchain/cli