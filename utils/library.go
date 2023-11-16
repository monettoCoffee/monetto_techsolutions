package utils

import (
	"context"
	"encoding/json"
	"fmt"
)

// CtxInfo print log with context
func LogInfo(ctx context.Context, format string, v ...interface{}) {
	// TODO easy impl ...
	fmt.Printf(format, v...)
}

// ToLogStr make object to string
func ToLogStr(param interface{}) string {
	if err, ok := param.(error); ok {
		return err.Error()
	}

	data, _ := json.Marshal(param)
	return string(data)
}

// GenOrderID 调用雪花算法生成订单 ID
func GenOrderID(ctx context.Context) string {
	// TODO no impl ...
	return "NO IMPL ..."
}
