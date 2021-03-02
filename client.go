package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-example/tracer"
)

var (
	traceH *tracer.TraceHandler
)

func main()  {
	// 生成jaeger trace
	traceH := tracer.InitTracer("Trace-Client")
	defer traceH.Closer.Close()
	// 设置为全局单例tracer
	opentracing.SetGlobalTracer(traceH.Tracer)

	// 启动http服务
	r := gin.Default()
	r.GET("/trace", tracer.AddTrace)
	r.Run(":9000")
}
