package config

import (
	"bufio"
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"regexp"
	"strings"
)

var Config config

type config struct {
	Database     Database
	Notify       Notify `yaml:"notify"`
	MigrationDir string `yaml:"migrationDir"`
	MigrationDb  string `yaml:"migrationDb"`
	LogDir       string `yaml:"logDir"`
	EnableLog    bool   `yaml:"enableLog"`
	SkipError    bool   `yaml:"skipError"`
}

func init() {
	configPath := flag.String("config", "config.yml", "配置文件")
	skipError := flag.Bool("skip-error", false, "是否跳过错误文件")
	flag.Parse()

	dataBytes, err := parseConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(dataBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "skip-error" {
			Config.SkipError = *skipError
		}
	})
}

func parseConfig(configSrc string) ([]byte, error) {
	// 打开配置文件
	file, err := os.Open(configSrc)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建一个正则表达式，用于匹配 ${变量名} 格式的字符串
	re := regexp.MustCompile(`\${(\w+)}`)

	// 读取文件内容并进行变量替换
	var result strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		replacedLine := re.ReplaceAllStringFunc(line, func(match string) string {
			// 提取变量名
			varName := re.FindStringSubmatch(match)[1]
			// 获取环境变量值
			envValue := os.Getenv(varName)
			if envValue == "" {
				// 如果环境变量不存在，保留原始字符串
				return match
			}
			return envValue
		})
		result.WriteString(replacedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return []byte(result.String()), nil
}
