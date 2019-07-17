package level

import (
	"strings"
)

type Level int

// 日志输出级别
const (
	Emergency Level = 600
	Alert           = 550
	Critical        = 500
	Error           = 400
	Warning         = 300
	Notice          = 250
	Info            = 200
	Debug           = 100
)

// GetLevelNameAbbreviation 获取日志级别缩写
func GetLevelNameAbbreviation(level Level) string {
	switch level {
	case Emergency:
		return "EMCY"
	case Alert:
		return "ALER"
	case Critical:
		return "CRIT"
	case Error:
		return "EROR"
	case Warning:
		return "WARN"
	case Notice:
		return "NOTI"
	case Info:
		return "INFO"
	case Debug:
		return "DEBG"
	}

	return "UNON"
}

// GetLevelName 获取日志级别名称
func GetLevelName(level Level) string {
	switch level {
	case Emergency:
		return "EMERGENCY"
	case Alert:
		return "ALERT"
	case Critical:
		return "CRITICAL"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Notice:
		return "NOTICE"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	}

	return "UNKNOWN"
}

// GetLevelByName 使用名称获取Level真实的数值
func GetLevelByName(levelName string) Level {
	switch strings.ToUpper(levelName) {
	case "EMERGENCY":
		return Emergency
	case "ALERT":
		return Alert
	case "CRITICAL":
		return Critical
	case "ERROR":
		return Error
	case "WARNING":
		return Warning
	case "NOTICE":
		return Notice
	case "INFO":
		return Info
	case "DEBUG":
		return Debug
	}

	return 0
}
