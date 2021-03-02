package tracer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"io/ioutil"
	"net/http"
	"strings"
)


var (
	tracer opentracing.Tracer
	Tracers opentracing.Tracer
	flag = make(chan bool)
)

func AddTrace(ctx *gin.Context) {

	// 创建span
	span := opentracing.StartSpan("span_order")
	// 在函数返回的时候调用finish结束这个span
	defer span.Finish()
	//创建上下文，使用上下文来传递span
	ctxn := opentracing.ContextWithSpan(context.Background(), span)

	// 业务逻辑
	orderNo := ctx.DefaultQuery("orderNo", "")
	span.SetTag("orderNo", orderNo)
	// 将ctx上下文传到调用的函数里
	ExtractUser(strings.TrimSpace(orderNo), ctxn)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": ctx.GetHeader("X-B3-TraceId"),
	})
}

func ChildTrace(ctx *gin.Context) {
	// 获取请求span信息
	fmt.Printf("%+v",ctx.Request.Header)
	spanCtx, _ := Tracers.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	span := Tracers.StartSpan("getOrderInfo", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	// 业务逻辑
	orderNo := ctx.DefaultQuery("orderNo", "")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": orderNo,
	})
}

func saveResponse(response []byte) error {
	err := ioutil.WriteFile("response.txt", response, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SendRequest(req *http.Request, ctx context.Context, flag chan bool) {
	reqPrepareSpan, _ := opentracing.StartSpanFromContext(ctx, "Client_sendRequest")
	defer reqPrepareSpan.Finish()

	go func(req *http.Request) {
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Printf("Do send requst failed(%s)\n", err)
			return
		}
		fmt.Printf("response (%v)\n", resp)

		respSpan, _ := opentracing.StartSpanFromContext(ctx, "Client_response")
		defer respSpan.Finish()

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ReadAll error(%s)\n", err)
			return
		}

		if resp.StatusCode != 200 {
			return
		}

		fmt.Printf("Response:%s\n", string(body))

		respSpan.LogFields(
			log.String("event", "getResponse"),
			log.String("value", string(body)),
		)

		saveResponse(body)
		flag <- true
	}(req)
}
