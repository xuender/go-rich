package rich

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

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

// 创建目录
func mkdir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
}

var args = pinyin.NewArgs()

func init() {
	slugify.Separator = " "
	args.Heteronym = true
}

// py 拼音转换
func py(str string) string {
	m := map[string]bool{}
	m[slugify.Slugify(str)] = true
	for _, str := range PC(pinyin.Pinyin(str, args)) {
		m[str] = true
	}
	ret := ""
	for k := range m {
		ret = ret + k + " "
	}
	return strings.TrimRight(ret, " ")
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
