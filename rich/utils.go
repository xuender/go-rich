package rich

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mozillazg/go-pinyin"
	"github.com/mozillazg/go-slugify"
)

// GetURL 获取网址
func GetURL(address string, tls bool) (string, error) {
	ip, err := GetIP()
	if err != nil {
		return "", err
	}
	if tls {
		return fmt.Sprintf("https://%s%s", ip, address), nil
	}
	return fmt.Sprintf("http://%s%s", ip, address), nil
}

// GetIP 获取IP地址
func GetIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !strings.HasPrefix(ipnet.IP.String(), "172") {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("未找到IP")
}

// Mkdir 创建目录
func Mkdir(path string) error {
	if fi, err := os.Stat(path); err == nil {
		if !fi.IsDir() {
			return fmt.Errorf("创建目录失败： %s 已经存在，并且不是目录。", path)
		}
	} else if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

var args = pinyin.NewArgs()

func init() {
	slugify.Separator = " "
	args.Heteronym = true
}

// PY 字符串拼音
func PY(str string) string {
	arr := []string{slugify.Slugify(str)}
	for _, s := range PC(pinyin.Pinyin(str, args)) {
		arr = append(arr, s)
	}
	sort.Strings(arr)
	ret := ""
	for i, k := range arr {
		if k == "" || (i > 0 && k == arr[i-1]) {
			continue
		}
		ret = ret + k + " ; "
	}
	return strings.Trim(ret, " ; ")
}

// Initial 首字母
func Initial(str string) []string {
	arr := []string{strings.ToUpper(slugify.Slugify(str))}
	for _, s := range PC(pinyin.Pinyin(str, args)) {
		arr = append(arr, strings.ToUpper(s))
	}
	sort.Strings(arr)
	ret := []string{}
	for i, k := range arr {
		if k == "" || (i > 0 && k[0] == arr[i-1][0]) {
			continue
		}
		ret = append(ret, string(k[0]))
	}
	return ret
}

// PC 排列组合
func PC(array [][]string) []string {
	return pc("", array)
}
func pc(s string, array [][]string) []string {
	if len(array) == 0 {
		return []string{s}
	}
	last := array[1:]
	ret := []string{}
	for _, i := range array[0] {
		larray := pc(i, last)
		for _, str := range larray {
			ret = append(ret, strings.Trim(s+" "+str, " "))
		}
	}
	return ret
}

var salt = []byte("ender weihai 2018-09-12")

// Pass 密码生成
func Pass(str string) []byte {
	h := md5.New()
	h.Write(salt)
	h.Write([]byte(str))
	h.Write(salt)
	return h.Sum(nil)
}

// Axis 装环Excel坐标
func Axis(col, row int) string {
	return fmt.Sprintf("%c%d", col+65, row+1)
}

// ReadXlsx 读取Excel文件
func ReadXlsx(file string, add func(row []string)) error {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return err
	}
	for _, name := range xlsx.GetSheetMap() {
		rows := xlsx.GetRows(name)
		for _, row := range rows {
			add(row)
		}
	}
	return nil
}
