package tracer

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func ExtractOrder(orderNo string, ctx context.Context) {
	// 创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_order_1_1")
	defer func() {
		span.SetTag("orderNo", orderNo)
		span.Finish()
	}()
}