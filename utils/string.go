package utils

import (
	"regexp"
	"strings"
)

// FilterSpaceLine 过滤空行
func FilterSpaceLine(s string) string {
	s = removeComments(s)
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

// 从 SQL 语句中移除注释
func removeComments(sql string) string {
	// 正则表达式模式，用于匹配 SQL 注释
	singleLineComment := regexp.MustCompile(`--.*\n`)
	multiLineComment := regexp.MustCompile(`/\*.*?\*/`)

	// 移除单行注释
	sql = singleLineComment.ReplaceAllString(sql, "")

	// 移除多行注释
	sql = multiLineComment.ReplaceAllString(sql, "")

	return sql
}
