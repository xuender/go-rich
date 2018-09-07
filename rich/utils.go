package rich

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mozillazg/go-slugify"
)

// 获取网址
func GetUrl(port string) (string, error) {
	ip, err := GetIp()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s%s", ip, port), nil
}

// 获取IP地址
func GetIp() (string, error) {
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

func init() {
	slugify.Separator = " "
}
func py(str string) string {
	return slugify.Slugify(str)
}
