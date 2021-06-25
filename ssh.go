// +----------------------------------------------------------------------
// | ssh
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"os/exec"
	"time"
)

type SSH struct {
	Host     string // 地址
	Username string // 用户名
	Password string // 密码
	Port     int    // 端口号
}

func (s *SSH) SSHClient() (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	clientConfig = &ssh.ClientConfig{
		User:    s.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d", s.Host, s.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

// 执行远程SSH命令行
// cmd 执行的命令行
func (s *SSH) ExecSSH(c string) ([]byte, error) {
	if session, err := s.SSHClient(); err != nil {
		return nil, err
	} else {
		defer session.Close()
		buf, e := session.Output(c)
		return buf, e
	}
}

// 执行本地命令行
// cmd 执行的命令行
func ExecCommand(c string) (string, error) {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	in.WriteString(c + "\n")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}