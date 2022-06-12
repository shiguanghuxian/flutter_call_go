package main

// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var (
	VERSION    = "0.0.0"
	BUILD_TIME = ""
	GO_VERSION = ""
	GIT_HASH   = ""
	DEBUG      = "false" // 可根据调试状态做相应测试逻辑
)

// 同步调用函数
//export SynchronousCall
func SynchronousCall(param1 *string, param2 int32) *C.char {
	defer func() {
		if err := recover(); err != nil {
			log.Println("同步调用异常")
		}
	}()

	// 模拟函数执行时长，暂停3秒
	time.Sleep(3 * time.Second)

	result := fmt.Sprintf("SynchronousCall: param1 = %s, param2 = %d", *param1, param2)
	return C.CString(result)
}

// 异步调用函数
//export AsynchronousCall
func AsynchronousCall(param1 *string, param2 int32) *C.char {
	defer func() {
		if err := recover(); err != nil {
			log.Println("异步调用异常")
		}
	}()

	// 模拟函数执行时长，暂停3秒
	time.Sleep(3 * time.Second)

	result := fmt.Sprintf("AsynchronousCall: param1 = %s, param2 = %d", *param1, param2)
	return C.CString(result)
}

func main() {
	if DEBUG != "true" {
		fmt.Println("质心im lib")
		return
	}
	version := map[string]string{
		"version":    VERSION,
		"build_time": BUILD_TIME,
		"go_version": GO_VERSION,
		"git_hash":   GIT_HASH,
	}
	data, _ := json.Marshal(version)
	fmt.Println(string(data))
}
