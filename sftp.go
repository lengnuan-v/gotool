// +----------------------------------------------------------------------
// | sftp
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"path"
	"strings"
	"time"
)

type SFTP struct {
	Host     string // 地址
	Username string // 用户名
	Password string // 密码
	Port     int    // 端口号
}

func (s *SFTP) SftpClient() (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	clientConfig = &ssh.ClientConfig{
		User:    s.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		// 需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", s.Host, s.Port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// SFTP 上传文件
// filePathName 文件路径名
// content 文件内容
func (s *SFTP) FileUpload(filePathName string, content []byte) error {
	if client, err := s.SftpClient(); err != nil {
		return err
	} else {
		defer client.Close()
		// 递归创建目录
		pathName := strings.Replace(filePathName, path.Base(filePathName), "", 1)
		if err := client.MkdirAll(strings.TrimRight(pathName, "/")); err != nil {
			return err
		}
		// 创建文件
		if file, err := client.Create(filePathName); err != nil {
			return err
		} else {
			defer file.Close()
			_, err = file.Write(content)
			return err
		}
	}
}

// SFTP 下载文件
// filePath 本地目录
// remotePath 远程文件路径名
func (s *SFTP) FileDownload(filePath, remotePathName string) error {
	if client, err := s.SftpClient(); err != nil {
		return err
	} else {
		defer client.Close()
		if file, err := client.Open(remotePathName); err != nil {
			return err
		} else {
			defer file.Close()
			dstFile, err := os.Create(path.Join(filePath, path.Base(remotePathName)))
			if err != nil {
				return err
			}
			defer dstFile.Close()
			_, err = file.WriteTo(dstFile)
			return err
		}
	}
}

// FTP 删除文件
// filePathName 删除文件
func (s *SFTP) FileDelete(filePathName string) error {
	if client, err := s.SftpClient(); err != nil {
		return err
	} else {
		defer client.Close()
		results, _ := client.Open(filePathName)
		if results != nil {
			return client.Remove(filePathName)
		}
	}
	return nil
}