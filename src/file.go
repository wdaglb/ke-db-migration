package src

import "strings"

// 获取路径内的文件名
func getFilename(val string) string {
	temps := strings.Split(val, "/")
	if len(temps) == 0 {
		return val
	}
	return temps[len(temps)-1]
}
