package main

// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"flutter_call_go/dart_api_dl"
	"fmt"
	"log"
	"time"
	"unsafe"
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

// 异步回调
var (
	callbackFunc func(string)      // 由flutter注册异步回调时赋值
	stopChan     = make(chan bool) // 异步任务停止管道
	isStop       = true            // 异步任务是否停止状态
)

//export InitializeDartApi
func InitializeDartApi(api unsafe.Pointer) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("用于注册回调异常", err)
		}
	}()
	// 用于注册回调使用
	dart_api_dl.Init(api)
}

// 注册回调
//export SetCallback
func SetCallback(port int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("注册回调异常", err)
		}
	}()
	go runCallback()
	log.Println("注册订阅回调")
	callbackFunc = func(data string) {
		dart_api_dl.SendStringToPort(port, data)
	}
}

// 停止异步回调
//export StopCallback
func StopCallback() {
	log.Println("停止回调任务")
	stopChan <- true
}

// 开启定时异步回调，模拟异步go传递数据给flutter
func runCallback() {
	if !isStop {
		return
	}
	isStop = false
	defer func() {
		isStop = true
	}()

	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t.C:
			sendDataToFlutter()
		case <-stopChan:
			t.Stop()
			return
		}
	}
}

// 异步发送消息
func sendDataToFlutter() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("异步发送消息异常", err)
		}
	}()
	if callbackFunc == nil {
		return
	}
	data := fmt.Sprintf("go异步回调flutter:%s", time.Now().Format(time.RFC3339))
	callbackFunc(data)
}

func main() {
	if DEBUG != "true" {
		fmt.Println("flutter call go")
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
