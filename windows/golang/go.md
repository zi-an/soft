 


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
