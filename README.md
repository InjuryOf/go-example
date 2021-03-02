# go-example
  
 ### 1.jaeger trace接入方案 example
 ```go
   //启动jaeer服务（基于docker镜像-all-in-one模式）
   1、 docker run -d --name jaeger \
    -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
    -p 5775:5775/udp \
    -p 6831:6831/udp \
    -p 6832:6832/udp \
    -p 5778:5778 \
    -p 16686:16686 \
    -p 14268:14268 \
    -p 9411:9411 \
   jaegertracing/all-in-one:lates 
   
   容器启动后可以通过访问：http://localhost:16686/来查看jaeger_ui界面
   
   // start api server
   2、 go run server.go

   // start api client
   3、go run client.go 
   
   **jaeger端口号说明**
   
   agent 暴露如下端口
   端口号	协议	功能
   5775	UDP	通过兼容性 thrift 协议，接收 zipkin thrift 类型的数据
   6831	UDP	通过二进制 thrift 协议，接收 jaeger thrift 类型的数据
   6832	UDP	通过二进制 thrift 协议，接收 jaeger thrift 类型的数据
   5778	HTTP	可用于配置采样策略
   
   collector 暴露如下端口
   端口号	协议	功能
   14267	TChannel	用于接收 jaeger-agent 发送来的 jaeger.thrift 格式的 span
   14268	HTTP	能直接接收来自客户端的 jaeger.thrift 格式的 span
   9411	HTTP	能通过 JSON 或 Thrift 接收 Zipkin spans，默认关闭
    
   query 暴露如下端口
   端口号	协议	功能
   16686	HTTP	1. /api/* - API 端口路径 2. / - Jaeger UI 路径
 ```
