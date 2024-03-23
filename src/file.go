package src

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取路径内的文件名
func getFilename(val string) string {
	temps := strings.Split(val, string(os.PathSeparator))
	if len(temps) == 0 {
		return val
	}
	return temps[len(temps)-1]
}

// 获取绝对路径
func getAbsDir(dir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	wd, err := os.Getwd()
	if err != nil {
		return dir
	}
	return wd + string(os.PathSeparator) + dir
}
