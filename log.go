// +----------------------------------------------------------------------
// | log
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
	"time"
)

var logger = logging.MustGetLogger("logging")

var fileFormatter = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{shortfile}:%{level} ➤ %{message}`,
)

var consoleFormatter = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05} %{shortfile}:%{level} ➤ %{message} %{color:reset}`,
)

func Log(dirname ...string) *logging.Logger {
	if IsEmpty(dirname) == true {
		logging.SetBackend(console())
	} else {
		logging.SetBackend(console(), file(dirname[0]))
	}
	return logger
}

// 文件输出
func file(dirname string) logging.LeveledBackend {
	// 检测目录是否存在，不存在就创建
	IsDirCreate(dirname)
	fileLog, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", dirname, time.Now().Format("2006-01-02")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	formatter := logging.NewBackendFormatter(logging.NewLogBackend(fileLog, "", 0), fileFormatter)
	backend := logging.SetBackend(logging.AddModuleLevel(formatter))
	return backend
}

// 控制台输出
func console() logging.Backend{
	return logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), consoleFormatter)
}

func Info(args ...interface{})  {
	Log().Info(args...)
}

func Error(args ...interface{})  {
	Log().Error(args...)
}

func Notice(args ...interface{})  {
	Log().Notice(args...)
}

func Debug(args ...interface{})  {
	Log().Debug(args...)
}

func Warning(args ...interface{})  {
	Log().Warning(args...)
}

func Critical(args ...interface{})  {
	Log().Critical(args...)
}