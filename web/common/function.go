package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"reflect"
	"strings"
	"time"
)

//@author:
//@function: ValidationPara
//@description: 获取不同语言的错误码信息
//@param: code int
//@param: languages string
//@return: error
func GetCodeMsg(code int, languages string) string {
	msg := ""

	if languages == "en" {
		// msg = GetEnMsgByCode(code)
	} else if languages == "zh" {
		// msg = GetZhMsgByCode(code)
	} else if languages == "zh_tw" {
		// msg = GetZhtwMsgByCode(code)
	} else if languages == "ja" {
		// msg = GetJaMsgByCode(code)
	} else {
		msg = GetZhMsgByCode(code)
	}
	return msg
}

//@author:
//@function: ValidationPara
//@description: 验证参数
//@param: obj interface{}
//@return: error
func ValidationPara(obj interface{}) error {
	// ps : 验证器验证validate:"required"，如果是数字，0会验证不通过，加上指针*，这样可以确保指针is not nil，而不是not 0
	v := validator.New()
	err := v.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil

}

//@author:
//@function: SliceContainsString
//@description: 切片是否包含元素
//@param: items []string
//@param: item string
//@return: bool
func SliceContainsString(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

//@author:
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

//@author:
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

//@author:
//@function: GetTodayTime
//@description: 获取当天起始和结束时间
//@param: time.Time, time.Time
//@return: time.Time, time.Time
func GetTodayTime() (time.Time, time.Time) {
	t := time.Now()

	// 起始
	startDate := t.Format("2006-01-02") + " 00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate, t.Location())

	// 结束
	endDate := t.Format("2006-01-02") + " 23:59:59"
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate, t.Location())

	return startTime, endTime
}

//@author:
//@function: GetHidePhone
//@description: 隐藏手机号码中间四位
//@param: phone string
//@return: string
func GetHidePhone(phone string) string {
	return phone[:3] + "****" + phone[7:]
}

//@author:
//@function: GetPage
//@description: 翻页
//@param: page int
//@param: limit int
//@return: int
func GetPage(page int, limit int) int {
	return (page - 1) * limit
}

//@author:
//@function: SliceUniqueInt
//@description: 去除slice重复元素
//@param: s []int
//@return: []int
func SliceUniqueInt(s []int) []int {
	ns := make([]int, 0, len(s))
	t := map[int]struct{}{}
	for _, i := range s {
		if _, ok := t[i]; !ok {
			t[i] = struct{}{}
			ns = append(ns, i)
		}
	}
	return ns
}

//@author:
//@function: SliceUniqueInt
//@description: 去除slice重复元素
//@param: s []uint64
//@return: []uint64
func SliceUniqueUint64(s []uint64) []uint64 {
	ns := make([]uint64, 0, len(s))
	t := map[uint64]struct{}{}
	for _, i := range s {
		if _, ok := t[i]; !ok {
			t[i] = struct{}{}
			ns = append(ns, i)
		}
	}
	return ns
}

//@author:
//@function: SliceUniqueInt
//@description: 去除slice重复元素
//@param: s []string
//@return: []string
func SliceUniqueString(s []string) []string {
	ns := make([]string, 0, len(s))
	t := map[string]struct{}{}
	for _, i := range s {
		if _, ok := t[i]; !ok {
			t[i] = struct{}{}
			ns = append(ns, i)
		}
	}
	return ns
}

//@author:
//@function: CreatePathDir
//@description: 根据路径创建文件夹
//@param: dirPath string
//@return:
func CreatePathDir(dirPath string) {
	// 返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(dirPath)
	switch {
	case os.IsNotExist(err):
		// 创建目录，完整目录路径，权限os.ModePerm为0777
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			panic(err)
		}
	case os.IsPermission(err):
		panic(err)
	}
}

//@author:
//@function: SliceToJoinString
//@description: 将Slice转为指定字符拼接的字符串
//@param: i interface{}
//@param: s string
//@return: string
func SliceToJoinString(i interface{}, s string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(i), "[]"), " ", s, -1)
}

//@author:
//@function: SliceChunkInt
//@description: 将切片分割多个数组
//@param: arr []int
//@param: num int
//@return: [][]int
func SliceChunkInt(arr []int, num int) [][]int {
	max := len(arr)

	if max <= num {
		return [][]int{arr}
	}

	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}

	var g [][]int
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num

		if i != quantity {
			g = append(g, arr[start:end])
		} else {
			g = append(g, arr[start:])
		}

		start = end
	}

	return g
}
