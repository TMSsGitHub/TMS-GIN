package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// 日志文件的根目录
	logsDir := "logs"
	// 获取当前日期
	currentDate := time.Now()
	year, month, day := currentDate.Date()

	// 日志文件名
	logsDir = fmt.Sprintf("%s/%d%d", logsDir, year, month)
	fmt.Println(logsDir)
	infoLogFile := fmt.Sprintf("%s/%d_info.log", logsDir, day)
	errorLogFile := fmt.Sprintf("%s/%d_error.log", logsDir, day)
	fmt.Println(infoLogFile)
	fmt.Println(errorLogFile)
}
