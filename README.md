# go-frame 
基于gin搭建的go框架

1、通过GOPROXY来控制代理;通过GOPRIVATE控制私有库不走代理。(version > 1.13)
   
   设置GOPROXY代理：
       
       go env -w GOPROXY=https://goproxy.cn,direct
       
   设置GOPRIVATE私有库，比如常用的Gitlab或Gitee，中间使用逗号分隔：
       
       go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
       
       这里指定域名为私有仓库，单个go get遇到此域名下的所有依赖时，会直接通过git进行
       
   如果在运行go mod vendor时，提示Get https://sum.golang.org/lookup/xxxxxx: dial tcp 216.58.200.49:443: i/o timeout，
   则是因为Go设置了默认的GOSUMDB=sum.golang.org，这个网站是被墙了的，用于验证包的有效性，可以通过如下命令关闭：
       
       go env -w GOSUMDB=off
    
   
   可以设置 GOSUMDB="sum.golang.google.cn"， 这个是专门为国内提供的sum 验证服务。
   
       go env -w GOSUMDB="sum.golang.google.cn"
   
2、mod常用指令
   go mod init	     生成 go.mod 文件
   go mod download	 下载 go.mod 文件中指明的所有依赖
   go mod tidy	     整理现有的依赖
   go mod graph	     查看现有的依赖结构
   go mod edit	     编辑 go.mod 文件
   go mod vendor	 导出项目所有的依赖到vendor目录
   go mod verify	 校验一个模块是否被篡改过
   go mod why	     查看为什么需要依赖某模块
   
   
   go mod init xxx  指定了模块导入路径为xxx
   
   go.mod文件中的module 用于定义当前项目的模块路径
   
   #GOPROXY
   这个环境变量主要是用于设置 Go 模块代理（Go module proxy），其作用是用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式，直接通过镜像站点来快速拉取。
   
   GOPROXY 的默认值是：https://proxy.golang.org,direct，这有一个很严重的问题，就是 proxy.golang.org 在国内是无法访问的，
   因此这会直接卡住你的第一步，所以你必须在开启 Go modules 的时，同时设置国内的 Go 模块代理，执行如下命令：
   go env -w GOPROXY=https://goproxy.cn,direct
   GOPROXY 的值是一个以英文逗号 “,” 分割的 Go 模块代理列表，允许设置多个模块代理，假设你不想使用，也可以将其设置为 “off” ，
   这将会禁止 Go 在后续操作中使用任何 Go 模块代理。
   
   #direct是什么
   “direct” 是一个特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取（比如 GitHub 等），
   场景如下：当值列表中上一个 Go 模块代理返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，
   也就是回到源地址去抓取，而遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误。
   
   
  