// +----------------------------------------------------------------------
// | Handler
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package controller

import (
	"reflect"
)

func Method(f interface{}, params ...interface{}) interface{} {
	in := make([]reflect.Value, len(params))
	for i, item := range params {
		in[i] = reflect.ValueOf(item)
	}
	return reflect.ValueOf(f).Call(in)[0].Interface()
}
