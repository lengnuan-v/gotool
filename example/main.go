// +----------------------------------------------------------------------
// | example
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package main

import (
	"fmt"
	"gotool"
)

func main() {
	db := gotool.DB{
		Dsn:    "root:1234567890@tcp(127.0.0.1:3306)/crxl?charset=utf8",
		Driver: "mysql",
		Prefix: "ob_",
	}
	fmt.Println(db.Update("action_log", map[string]interface{}{"username":2}, "id=1951"))
}
