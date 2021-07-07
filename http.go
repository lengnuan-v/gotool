// +----------------------------------------------------------------------
// | http
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"crypto/tls"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Method GET POST
// Body bytes.NewReader()
// Proxy http://IP:Proxy or socks5://Username:Password@IP:Proxy
type ClientOptions struct {
	Url     string
	Method  string
	Body    io.Reader
	Header  map[string]string
	Proxy   []byte
	Timeout int64
}

func (c *ClientOptions) HttpRequest() ([]byte, []byte, error) {
	var err error
	var request *http.Request
	if request, err = http.NewRequest(c.Method, c.Url, c.Body); err != nil {
		return nil, nil, err
	}
	var client *http.Client
	// 代理 Client
	if IsEmpty(c.Proxy) == false { client = c.proxyClient()} else { client = &http.Client{} }
	// Header
	if IsEmpty(c.Header) == false { for k, v := range c.Header { request.Header.Set(k, v) }}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, nil, err
	}
// 	defer response.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	return body, CookiesString(response.Cookies()), nil
}

// 代理 Client
func (c *ClientOptions) proxyClient() *http.Client {
	if IsEmpty(c.Timeout) == true { c.Timeout = 30 }
	proxyURL, _ := url.Parse(string(c.Proxy))
	return &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout:   time.Second * time.Duration(c.Timeout), // 超时时间
	}
}

// Cookie 解析 字符串
func CookiesString(cookies []*http.Cookie) []byte {
	if len(cookies) <= 0 {
		return nil
	}
	var c []string
	for _,  cookie := range cookies {
		var list = Explode(";",  cookie.String())
		for _, v := range list {
			v = Trim(v, " ")
			if b, _ := InArray(v, c); b == false {
				c = append(c, v)
			}
		}
	}
	return []byte(Implode("; ", c))
}

// 随机UA (选择不同的浏览器 UA)
func UA(source ...string) string {
	var name = If(IsEmpty(source) == true, []string{"default"}, source)
	switch name.([]string)[0] {
	case "Chrome":
		return browser.Chrome()
	case "InternetExplorer":
		return browser.InternetExplorer()
	case "Firefox":
		return browser.Firefox()
	case "Safari":
		return browser.Safari()
	case "Android":
		return browser.Android()
	case "MacOSX":
		return browser.MacOSX()
	case "IOS":
		return browser.IOS()
	case "Linux":
		return browser.Linux()
	case "IPhone":
		return browser.IPhone()
	case "IPad":
		return browser.IPad()
	case "Computer":
		return browser.Computer()
	default:
		return browser.Random()
	}
}
