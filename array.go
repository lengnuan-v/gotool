// +----------------------------------------------------------------------
// | array
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"bytes"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// 字符串转数组
// separator 规定在哪里分割字符串。
// string 要分割的字符串。
// Explode(",", "a,b,c")
func Explode(separator, str string) []string {
	return strings.Split(str, separator)
}

// 数组转字符串
// separator 规定数组元素之间放置的内容
// array 要组合为字符串的数组。
// Implode(",", []string{"a", "b", "c"})
func Implode(separator string, array []string) string {
	var buf bytes.Buffer
	l := len(array)
	for _, str := range array {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(separator)
		}
	}
	return buf.String()
}
// 数组分割数组块
// array 规定要使用的数组。
// size	整数值，规定每个新数组包含多少个元素。
func ArrayChunk(array []interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("size: 不能小于1")
	}
	length := len(array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, array[i*size:end])
		i++
	}
	return n
}

// 搜索数组中是否存在指定的值
// search 规定要在数组搜索的值。
// array 规定要搜索的数组。
// InArray(1, [2]interface{}{"a", 1})
func InArray(search interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(search, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

// 移除数组中的重复的值，并返回结果数组
// array 规定数组。
// ArrayUnique([]string("a", "b", "a"))
func ArrayUnique(array []interface{}) (newArr []interface{}) {
	newArr = make([]interface{}, 0)
	for i := 0; i < len(array); i++ {
		repeat := false
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, array[i])
		}
	}
	return
}

// 返回输入数组中某个单一列的值
// array 规定要使用的多维数组
// key 需要返回值的列
// ArrayColumn([][]interface{}{{"id":"a", "name":"b"}, {"id":"c", "name":"d"}}, "id")
func ArrayColumn(array []interface{}, key string) []interface{} {
	rt := reflect.TypeOf(array)
	rv := reflect.ValueOf(array)
	if rt.Kind() == reflect.Slice { //切片类型
		var sliceColumn []interface{}
		elemt := rt.Elem() //获取切片元素类型
		for i := 0; i < rv.Len(); i++ {
			inxv := rv.Index(i)
			if elemt.Kind() == reflect.Struct {
				for i := 0; i < elemt.NumField(); i++ {
					if elemt.Field(i).Name == key {
						strf := inxv.Field(i)
						switch strf.Kind() {
						case reflect.String:
							sliceColumn = append(sliceColumn, strf.String())
						case reflect.Float64:
							sliceColumn = append(sliceColumn, strf.Float())
						case reflect.Int, reflect.Int64:
							sliceColumn = append(sliceColumn, strf.Int())
						default:
							//do nothing
						}
					}
				}
			}
		}
		return sliceColumn
	}
	return nil
}

// 返回包含数组中所有键名的一个新数组
// array 数组
func ArrayKeys(array map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(array))
	for key := range array {
		keys[i] = key
		i++
	}
	return keys
}

// 返回一个包含给定数组中所有键值的数组，但不保留键名
// elements 数组
func ArrayValues(array map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(array))
	for _, val := range array {
		vals[i] = val
		i++
	}
	return vals
}

// 把多个数组合并为一个数组
// array 数组
func ArrayMerge(array ...[]interface{}) []interface{} {
	n := 0
	for _, v := range array {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range array {
		s = append(s, v...)
	}
	return s
}

// 在数组中根据条件取出一段值
// array 数组
// offset 取出元素的开始位置
// length 返回数组的长度
func ArraySlice(array []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(array)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(array)) {
		return array[offset:end]
	}
	return array[offset:]
}

// 返回数组中的随机键名
// array 数组
func ArrayRand(array []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]interface{}, len(array))
	for i, v := range r.Perm(len(array)) {
		n[i] = array[v]
	}
	return n
}

// 在数组尾部添加一个或多个元素
// array 数组
// elements 添加的值
// ArrayPush(&s1, "u", "v")
func ArrayPush(array *[]interface{}, elements ...interface{}) int {
	*array = append(*array, elements...)
	return len(*array)
}

// 删除数组中的最后一个元素
// array 数组
func ArrayPop(array *[]interface{}) interface{} {
	if len(*array) == 0 {
		return nil
	}
	ep := len(*array) - 1
	e := (*array)[ep]
	*array = (*array)[:ep]
	return e
}

// 删除数组中第一个元素，并返回被删除元素的值
// array 数组
func ArrayShift(array *[]interface{}) interface{} {
	if len(*array) == 0 {
		return nil
	}
	f := (*array)[0]
	*array = (*array)[1:]
	return f
}

// 检查某个数组中是否存在指定的键名
// ArrayKeyExists("a", map[interface{}]interface{"a":"a","b":b})
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

// 合并两个数组来创建一个新数组，其中的一个数组元素为键名，另一个数组的元素为键值
func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	if len(s1) != len(s2) {
		panic("the number of elements for each slice isn't equal")
	}
	m := make(map[interface{}]interface{}, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}
	return m
}

// 相反的元素顺序返回数组
func ArrayReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}