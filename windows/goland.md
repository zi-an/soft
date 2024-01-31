# GoLand

|平台|下载|
|-|-|
|windows|https://download.jetbrains.com/go/goland-2020.1.1.win.zip|
|linux|https://download.jetbrains.com/go/goland-2020.1.1.tar.gz|

* 必须2020.1.1版本才能用jetbrains-agent.jar与zh.201.16.jar
* [go version](https://dl.google.com/go/go1.16.15.windows-amd64.zip)不能高于16,会无法识别
* 最后支持XP系统的是32位的[GO 1.10.8](https://golang.google.cn/dl/go1.10.8.windows-386.zip)
* epel里支持centos的版本是go 18

## 环境变量GOPATH设置

最好设置GOPATH,不然go get会把包下载到用户目录内,不方便编译



 


# 查找包看GO111MODULE的用法
go env -w GO111MODULE=off


# 环境变量地址
C:\Users\root\AppData\Roaming\go下的env文件删除即可清空

# 下载gin模块
go get -u github.com/gin-gonic/gin

# go使用模块化编译
go mod init gin.mod
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct




# GoLand使用模块
在GoLand设置中，将这个Enable Go modules intergration勾选上，然后Apply即可
