package main

import (
	"github.com/gin-gonic/gin"
	"go-example/tracer"
)

func main()  {

	// 生成jaeger trace
	tracerH := tracer.InitTracer("Trace-Server")
	tracer.Tracers = tracerH.Tracer
	defer tracerH.Closer.Close()

	// 启动http服务
	r := gin.Default()
	r.GET("/getOrderInfo", tracer.ChildTrace)
	r.Run(":9001")
}
