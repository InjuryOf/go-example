package tracer

import (
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// 采样率
const JaegerSampleParam = 1

// Jaeger-Agent上报地址
const JaegerReportingHost = "127.0.0.1:6831"

type TraceHandler struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func InitTracer(serviceName string) TraceHandler {
	// 初始化jaeger配置上下文对象
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: JaegerSampleParam,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: JaegerReportingHost,
		},
	}

	// 设置服务名称
	cfg.ServiceName = serviceName

	// 创建tracer
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Errorf("NewTracer失败")
	}
	return TraceHandler{
		Tracer: tracer,
		Closer: closer,
	}
}
