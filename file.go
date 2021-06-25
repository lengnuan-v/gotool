// +----------------------------------------------------------------------
// | file
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

// 读取一个文件
// filename	规定要读取的文件。
func ReadFile(filename string) ([]byte, error) {
	if data, err := ioutil.ReadFile(filename); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

// 创建并写入文件
// content 写入的内容
// filename 规定要写入的文件。
// keep true 已有的数据会被保留 false 已有的数据会被清除
func Tracefile(content []byte, filename string, keep bool) (int, error) {
	dir, _ := path.Split(filename)
	_ = os.MkdirAll(dir, 0777)
	var line string
	var file *os.File
	if keep == true {
		line = "\n"
		file, _ = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	} else {
		file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	}
	defer file.Close()
	fileContent := strings.Join([]string{string(content), line}, "")
	return file.Write([]byte(fileContent))
}

// 检查文件或目录是否存在
// filename 指定的文件或目录
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断给定文件名是否是一个目录
// filename 指定的文件或目录
func IsDir(filename string) (bool, error) {
	if fd, err := os.Stat(filename); err != nil {
		return false, err
	} else {
		fm := fd.Mode()
		return fm.IsDir(), nil
	}
}

// 检测目录是否存在，不存在就创建
// path 指定的目录
func IsDirCreate(path string) {
	if e := FileExists(path); e == false {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Fatalln(err)
		}
	}
}

// 以数组的形式返回文件路径的信息
// path 文件路径
// options -1: all; 1: dirname; 2: basename; 4: extension; 8: filename
// Pathinfo("/home/go/php2go.go.go", -1)
func Pathinfo(path string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}
	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(path)
	}
	if (options & 2) == 2 {
		info["basename"] = filepath.Base(path)
	}
	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""
		if (options & 2) == 2 {
			basename, _ = info["basename"]
		} else {
			basename = filepath.Base(path)
		}
		p := strings.LastIndex(basename, ".")
		filename, extension := "", ""
		if p > 0 {
			filename, extension = basename[:p], basename[p+1:]
		} else if p == -1 {
			filename = basename
		} else if p == 0 {
			extension = basename[p+1:]
		}
		if (options & 4) == 4 {
			info["extension"] = extension
		}
		if (options & 8) == 8 {
			info["filename"] = filename
		}
	}
	return info
}

// 检查指定的文件是否是常规的文件
func IsFile(filename string) bool {
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 返回指定文件的大小
func FileSize(filename string) (int64, error) {
	if info, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

// 删除文件
func Unlink(filename string) error {
	return os.Remove(filename)
}

// 拷贝文件
// source 要复制的文件
// dest 复制文件的目的地
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// 判断指定文件名是否可读
func IsReadable(filename string) bool {
	if _, err := syscall.Open(filename, syscall.O_RDONLY, 0); err != nil {
		return false
	}
	return true
}

// 判断指定的文件是否可写
func IsWriteable(filename string) bool {
	if _, err := syscall.Open(filename, syscall.O_WRONLY, 0); err != nil {
		return false
	}
	return true
}

// 重命名文件或目录
// oldname 要重命名的文件或目录
// newname 文件或目录的新名称
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// 获取当前工作目录
func Getcwd() (string, error) {
	return os.Getwd()
}

// 返回绝对路径
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// 返回路径中的文件名部分
func Basename(path string) string {
	return filepath.Base(path)
}

// 改变文件模式
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// 返回匹配指定模式的文件名或目录
func Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// 递归路径
func RecursiveListPath(path string, slice *[]string) {
	if path == "/" {
		return
	}
	path2 := filepath.Dir(fmt.Sprintf("/%s", strings.TrimLeft(path, "/")))
	RecursiveListPath(path2, slice)
	*slice = append(*slice, path)
}