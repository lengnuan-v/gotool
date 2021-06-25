// +----------------------------------------------------------------------
// | string
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"html"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// string转成int
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// string转成int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// int转成string：
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// int64转成string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// string转float
// size 32 -> float32 64 -> float64
func StringToFloat(str string, size int) (float64, error) {
	return strconv.ParseFloat(str, size)
}

// float转string
// size 32 -> float32 64 -> float64
func FloatToString(f float64 , size int) string {
	return strconv.FormatFloat(f, 'f', -1, size)
}

// 移除字符串两侧的空白字符或其他预定义字符
// str 要检查的字符串
// charlist 从字符串中删除哪些字符
func Trim(str string, charlist ...string) string {
	mask := ""
	if len(charlist) == 0 {
		mask = " \\t\\n\\r\\x0B"
	} else {
		mask = charlist[0]
	}
	return strings.Trim(str, mask)
}

// 移除字符串左侧的空白字符或其他预定义字符
// str 要检查的字符串
// charlist 从字符串中删除哪些字符
func Ltrim(str string, charlist ...string) string {
	mask := ""
	if len(charlist) == 0 {
		mask = " \\t\\n\\r\\x0B"
	} else {
		mask = charlist[0]
	}
	return strings.TrimLeft(str, mask)
}

// 移除字符串右侧的空白字符或其他预定义字符
// str 要检查的字符串
// charlist 从字符串中删除哪些字符
func Rtrim(str string, charlist ...string) string {
	mask := ""
	if len(charlist) == 0 {
		mask = " \\t\\n\\r\\x0B"
	} else {
		mask = charlist[0]
	}
	return strings.TrimRight(str, mask)
}

// 搜索字符串在另一字符串中的第一次出现
// str 被搜索的字符串
// search 所搜索的字符串
// Strstr("xxx@gmail.com", "@")
func Strstr(str string, search string) string {
	if search == "" {
		return ""
	}
	idx := strings.Index(str, search)
	if idx == -1 {
		return ""
	}
	return str[idx+len(search):]
}

// 获取字符串长度 (单个汉字 = 3)
// str 需要计算长度的字符串
func Strlen(str string) int {
	return len(str)
}

// 获取字符串的长度 (单个汉字 = 1)
// str 需要计算长度的字符串
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// 返回字符串的一部分
// str 要检查的字符串
// start 字符串的何处开始
// length 要返回的字符串长度
// Substr("abc", 0, 2)
func Substr(str string, start int, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

// 把字符串中每个单词的首字符转换为大写
// str 要转换的字符串
func Ucwords(str string) string {
	return strings.Title(str)
}

// 首字符转换为小写
// str 要转换的字符串
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return str
}

// 首字符转换为大写
// str 要转换的字符串
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return str
}

// 字符转换为大写
// str 要转换的字符串
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

// 字符转换为小写
// str 要转换的字符串
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// 以其他字符替换字符串中的一些字符（区分大小写）
// search 要查找的值
// replace 替换 search 中的值
// str 被搜索的字符串
// count 对替换数进行计数的变量, 替换的数量没有限制使用 -1
// StrReplace("a","b","abcd", -1)
func StrReplace(search, replace, str string, count int) string {
	return strings.Replace(str, search, replace, count)
}

// 查找字符串在另一字符串中最后一次出现的位置（区分大小写）
// str 被搜索的字符串
// needle 要查找的字符
// offset 在何处开始搜索 1
// Strrpos("hello word", "w", 1)
func Strrpos(str, needle string, offset int) int {
	pos, length := 0, len(str)
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		str = str[:offset+length+1]
	} else {
		str = str[offset:]
	}
	pos = strings.LastIndex(str, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// 查找字符串在另一字符串中最后一次出现的位置（不区分大小写）
// str 被搜索的字符串
// needle 要查找的字符
// offset 在何处开始搜索, 没有限制使用 1
// Strripos("hello word", "w", 1)
func Strripos(str, needle string, offset int) int {
	pos, length := 0, len(str)
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		str = str[:offset+length+1]
	} else {
		str = str[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(str), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// 查找字符串在另一字符串中第一次出现的位置（不区分大小写）
// str 被搜索的字符串
// needle 要查找的字符
// offset 在何处开始搜索 1
// Stripos("hello word", "w", 1)
func Stripos(str, needle string, offset int) int {
	length := len(str)
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	str = str[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(str), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// 查找字符串在另一字符串中第一次出现的位置（区分大小写）
// str 被搜索的字符串
// needle 要查找的字符
// offset 在何处开始搜索 1
// Strpos("hello word", "w", 1)
func Strpos(str, needle string, offset int) int {
	length := len(str)
	if length == 0 || offset > length || -offset > length {
		return -1
	}
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(str[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// 反转字符串
// str 要反转的字符串
func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 把字符串分割后添加指定end
// body 要分割的字符串
// chunklen 数字值，定义字符串块的长度
// end 定义在每个字符串块末端放置的内容
// ChunkSplit("abc", 1, "e") => aebece
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

// 转换字符串中特定的字符
// haystack 要转换的字符串
// params 要改变的字符, 要改变为的字符
// Strtr("baab", "ab", "01")
func Strtr(haystack string, params ...interface{}) string {
	ac := len(params)
	if ac == 1 {
		pairs := params[0].(map[string]string)
		length := len(pairs)
		if length == 0 {
			return haystack
		}
		oldnew := make([]string, length*2)
		for o, n := range pairs {
			if o == "" {
				return haystack
			}
			oldnew = append(oldnew, o, n)
		}
		return strings.NewReplacer(oldnew...).Replace(haystack)
	} else if ac == 2 {
		from := params[0].(string)
		to := params[1].(string)
		trlen, lt := len(from), len(to)
		if trlen > lt {
			trlen = lt
		}
		if trlen == 0 {
			return haystack
		}

		str := make([]uint8, len(haystack))
		var xlat [256]uint8
		var i int
		var j uint8
		if trlen == 1 {
			for i = 0; i < len(haystack); i++ {
				if haystack[i] == from[0] {
					str[i] = to[0]
				} else {
					str[i] = haystack[i]
				}
			}
			return string(str)
		}
		// trlen != 1
		for {
			xlat[j] = j
			if j++; j == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[from[i]] = to[i]
		}
		for i = 0; i < len(haystack); i++ {
			str[i] = xlat[haystack[i]]
		}
		return string(str)
	}
	return haystack
}

// 在每个双引号（"）前添加反斜杠
// str 要转义的字符串
func Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// 删除反斜杠
// str 要检查的字符串
func Stripslashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// 在预定义字符前添加反斜杠
// str 要检查的字符串
func Quotemeta(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// 字符转换为 HTML 实体
// str 要转换的字符串
func Htmlentities(str string) string {
	return html.EscapeString(str)
}

// 把 HTML 实体转换为字符
// str 要转换的字符串
func HTMLEntityDecode(str string) string {
	return html.UnescapeString(str)
}

// MD5加密
// str 要加密的字符串
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// 计算两个字符串的相似度，并返回匹配字符的数目
// first 比较的第一个字符串
// second 比较的第二个字符串
// percent 百分比相似度的变量名
// SimilarText("golang", "google", &percent)
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0
		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}
		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}
		return sum
	}
	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

// 使用base64对数据进行编码
// str 要编码的数据
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解码base64编码的数据
// str 要编码的数据
func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 获取字符串中2个字符串之间到内容
func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

// Unicode 转中文
func UnescapeUnicode(raw string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return html.UnescapeString(str), nil
}

// 删除标点符号
func DeletePunctuation(str string) string {
	return regexp.MustCompile("\\pP|\\pS").ReplaceAllString(str, "")
}