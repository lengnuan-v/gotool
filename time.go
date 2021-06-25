// +----------------------------------------------------------------------
// | time
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import "time"

// 停顿、暂停
// t 秒数
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// 获取当前时间戳
func Time() int64 {
	return time.Now().Unix()
}

// 获取当前日期时间
// format 格式 2006-01-02 15:04:05 必须是这个时间点, 据说是go诞生之日（返回的格式）
func Date(format string) string {
	return time.Now().Format(format)
}

// 日期时间转换时间戳
// format 格式 2006-01-02 15:04:05 必须是这个时间点, 据说是go诞生之日
// strtime 需要转化时间戳的日期时间
func StrtoTime(format, strtime string) (int64, error) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(format, strtime, loc)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// 时间戳转日期时间
// format 格式 2006-01-02 15:04:05 必须是这个时间点, 据说是go诞生之日（返回的格式）
// timestamp 需要转化的时间戳
func DateTime(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}