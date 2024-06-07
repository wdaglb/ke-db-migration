package utils

import "strings"

// FilterSpaceLine 过滤空行
func FilterSpaceLine(s string) string {
	// 将输入字符串按行分割
	lines := strings.Split(s, "\n")

	// 创建一个新的切片保存非空行
	var filteredLines []string

	// 遍历所有行并过滤掉空行
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filteredLines = append(filteredLines, line)
		}
	}

	// 将过滤后的行连接成一个新的字符串
	return strings.Join(filteredLines, "\n")
}
