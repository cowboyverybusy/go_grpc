【⭐go：使用grpc】
正如其他 RPC 系统，gRPC 基于如下思想：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。
gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构。

《执行步骤》
1.定义gRPC服务（写proto文件）。
2.生成客户端和服务器代码（protoc --go_out=plugins=grpc:. helloworld.proto）。生成代码后，执行 go mod tidy,下载缺少的包。
3.实现gRPC服务。
4.实现gRPC客户端。

《正式部署》
1、在同一台服务器运行的话，只需要开两个终端，分别进入server和client文件夹执行go run main.go，启动服务端和客户端。
2、不同服务器运行的话，修改客户端的监听地址为服务端所在服务器的IP地址加端口号。然后分别执行下面命令，生成二进制文件：
cd server;go build -o grpc_server main.go
cd client;go build -o grpc_client main.go
把两个文件复制到不同服务器上，启动即可：
chmod u+x grpc_client
./grpc_client


《遇到的bug》
在proto文件所在文件夹下执行如下命令：
protoc --go_out=plugins=grpc:. helloworld.proto
提示报错：
protoc-gen-go: unable to determine Go import path for "helloworld.proto"
Please specify either:
	• a "go_package" option in the .proto source file, or
	• a "M" argument on the command line.

解决方法：在helloworld.proto文件中指定生成.pb.go文件的路径：
syntax = "proto3";
package hellowrld;
option go_package ="../helloworld";//必须指定路径生成.pb.go文件的路径。
因为proto文件本来就在helloworld文件夹中，我们这里返回到上一级再回到helloworld文件夹，这样生成的.pb.go文件就和proto文件同目录了。.pb.go文件也不会缺少package包。


