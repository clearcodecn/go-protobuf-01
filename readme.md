### Grpc Starter 

This is a simple starter for grpc with golang and protobuf.

欢迎加入学习群
![https://github.com/clearcodecn/leveldb-demo/raw/master/images/qr.png](https://github.com/clearcodecn/leveldb-demo/raw/master/images/qr.png)

#### Resources 

* [golang 语言](https://golang.org/)
* [protobuf 介绍](https://developers.google.com/protocol-buffers/)
* [grpc 官网](https://grpc.io/)
* [golang-protobuf 套件](https://github.com/golang/protobuf)
* [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* [grpc-web](https://github.com/improbable-eng/grpc-web)
* [grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)
* [awesome-grpc](https://github.com/grpc-ecosystem/awesome-grpc#go)


#### Installation

##### install protoc

* 打开 [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases) 下载最新版本
* 目录说明
```javascript
☁  protobuf  tree -l
.
├── bin
│ └── protoc                # 二进制文件
├── include
│ └── google
│     └── protobuf          # protobuf 相关的 proto 文件
│         ├── any.proto
│         ├── api.proto
│         ├── compiler
│         │ └── plugin.proto
│         ├── descriptor.proto
│         ├── duration.proto
│         ├── empty.proto
│         ├── field_mask.proto
│         ├── source_context.proto
│         ├── struct.proto
│         ├── timestamp.proto
│         ├── type.proto
│         └── wrappers.proto
└── readme.txt             # 说明文件
5 directories, 14 files
```
* 执行命令, 检查安装情况. 
```shell
☁  protobuf  ./bin/protoc --version
libprotoc 22.3
```

* 将 protoc 所在目录添加到环境变量
```javascript
> 查看所在目录 
$ pwd
> /tmp/protobuf

$ export PATH=$PATH:/tmp/protobuf/bin   # 临时添加到当前终端 , 永久添加则需要修改 ~/.bash_profile 文件
$ protoc --version
> libprotoc 22.3
```


##### 安装 go protobuf 相关套件. 

* 打开 [https://protobuf.dev/reference/go/go-generated/](https://protobuf.dev/reference/go/go-generated/)
* 安装:
```shell 
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

##### 编译protobuf:

```javascript
make build-proto
```


#### 使用 GRPC 

1. 安装 [文档](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```