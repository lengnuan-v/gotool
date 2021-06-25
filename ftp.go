// +----------------------------------------------------------------------
// | ftp
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

type FTP struct {
	Host     string // 地址
	Username string // 用户名
	Password string // 密码
	Port     int    // 端口号
}

func (f *FTP) FtpClient() (*ftp.ServerConn, error) {
	if conn, err := ftp.Dial(fmt.Sprintf("%s:%d", f.Host, f.Port), ftp.DialWithTimeout(10*time.Second)); err != nil {
		return nil, err
	} else {
		if err := conn.Login(f.Username, f.Password); err != nil {
			return nil, err
		}
		return conn, nil
	}
}

// FTP 上传文件
// filePathName 文件路径名
// content 文件内容
func (f *FTP) FileUpload(filePathName string, content []byte) error {
	if conn, err := f.FtpClient(); err != nil {
		return err
	} else {
		defer conn.Quit()
		// 递归创建目录
		var pathList []string
		dir := strings.Replace(filePathName, path.Base(filePathName), "", 1)
		RecursiveListPath(strings.TrimRight(dir, "/"), &pathList)
		for _, path := range pathList {
			_ = conn.MakeDir(path)
		}
		// 写入文件
		if err := conn.Stor(filePathName, bytes.NewReader(content)); err != nil {
			return err
		}
	}
	return nil
}

// FTP 下载文件
// filePath 本地目录
// remotePath 远程文件路径名
func (f *FTP) FileDownload(filePath, remotePathName string) error {
	if conn, err := f.FtpClient(); err != nil {
		return err
	} else {
		defer conn.Quit()
		res, err := conn.Retr(remotePathName)
		if err != nil {
			return err
		}
		defer res.Close()
		// 读取内容
		if buf, err := ioutil.ReadAll(res); err != nil {
			return err
		} else {
			// 保存文件
			fileName := path.Base(remotePathName)
			_, err = Tracefile(buf, fmt.Sprintf("%s/%s", strings.TrimRight(filePath, "/"), fileName), false)
			return err
		}
	}
}

// FTP 删除文件
// filePathName 删除文件
func (f *FTP) FileDelete(filePathName string) error {
	if conn, err := f.FtpClient(); err != nil {
		return err
	} else {
		defer conn.Quit()
		return conn.Delete(filePathName)
	}
}