package tracer

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func ExtractUser(orderNo string, ctx context.Context) {
	// 创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "Client_request")
	defer span.Finish()

	// 发起http请求
	requestUrl := "http://localhost:9001/getOrderInfo"
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// http请求（traceing）
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, requestUrl)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	flag = make(chan bool)
	SendRequest(req, ctx, flag)

	<- flag

	// 请求其他方法
	//ExtractOrder(orderNo, ctx)
}


