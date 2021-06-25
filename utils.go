// +----------------------------------------------------------------------
// | utils
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"encoding/binary"
	"net"
	"os"
	"reflect"
	"runtime"
)

// 检查一个变量是否为空
func IsEmpty(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// 三元表达式
func If(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}

// 获取当前操作系统
func Goos() string {
	return runtime.GOOS
}

// 获取当前系统HOME
func GetUserHome() string {
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}
	return home
}

// 获取主机名
func Gethostname() (string, error) {
	return os.Hostname()
}

// 将长整数地址转换为（IPv4）Internet标准点分格式的字符串
// IP2long("8.8.8.8")
func IP2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// 将长整数地址转换为（IPv4）Internet标准点分格式的字符串
// Long2ip(134744072)
func Long2ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}