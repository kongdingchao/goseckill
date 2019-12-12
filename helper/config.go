package helper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

//配置文件的全局map
var ConfigMap sync.Map

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//配置初始化
func InitConfig(path string) {
	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	checkErr(err)
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		ConfigMap.Store(key, value)
	}
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@param			def				default，即默认值
//@return           val             值
func GetConfig(cate ConfigCate, key string, def ...string) string {
	if val, ok := ConfigMap.Load(fmt.Sprintf("%v.%v", cate, key)); ok {
		return val.(string)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@return           val             值
func GetConfigBool(cate ConfigCate, key string) (val bool) {
	value := GetConfig(cate, key)
	if value == "true" || value == "1" {
		val = true
	}
	return
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@return           val             值
func GetConfigInt64(cate ConfigCate, key string) (val int64) {
	val, _ = strconv.ParseInt(GetConfig(cate, key), 10, 64)
	return
}

//获取配置
//@param            cate            配置分类
//@param            key             键
//@return           val             值
/*func GetConfigFloat64(cate ConfigCate, key string) (val float64) {
	val, _ = strconv.ParseFloat(GetConfig(cate, key), 64)
	return
}*/
