package utils

import (
	"runtime"
	"strconv"
	"time"
)

type ErrorLog struct {
	Path         string    `json:"path"`
	Line         string    `json:"line"`
	ErrorMessage string    `json:"error_message"`
	CreatedAt    time.Time `json:"created_at"`
}

type CustomLog struct {
	Data      interface{} `json:"data"`
	CreatedAt time.Time   `json:"created_at"`
}

func LogCustom(obj interface{}, esHelper ElasticsearchHelper) {
	go func() {
		customLog := CustomLog{
			Data:      obj,
			CreatedAt: time.Now(),
		}
		esHelper.Push("custom_logs", customLog)
	}()
}

//错误日志记录
func HandleError(logType string, err error, esHelper ElasticsearchHelper) (hasError bool) {
	hasError = false
	if err != nil {
		hasError = true
		_, errorFilePath, line, _ := runtime.Caller(1)

		newErrorLog := ErrorLog{
			Path:         errorFilePath,
			Line:         strconv.Itoa(line),
			ErrorMessage: err.Error(),
			CreatedAt:    time.Now(),
		}

		go func() {
			esHelper.Push(logType+"_runtime_errors", newErrorLog)
		}()
	}
	return
}
