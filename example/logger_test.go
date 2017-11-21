// Copyright 2017 Blurt. All rights reserved.
// Use of this source code is governed by a GPL-3.0
// license that can be found in the LICENSE file.
//
// test go logger

package example

import (
	"errors"
	"runtime"
	"testing"
	"time"

	"github.com/blurtheart/go-logger/logger"
)

func log() {
	logger.Debug("hello world")
	logger.Info(1)
	logger.Warn("dangerous operation")
	logger.Error(errors.New("Method not allowed"))
	logger.Fatal("fatal error:", errors.New("Method not allowed"))
}

func Test(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//指定是否控制台打印，默认为true
	logger.SetConsole(true)
	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位
	//logger.SetRollingFile("d:/logtest", "test.log", 10, 5, logger.KB)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	logger.SetRollingDaily("d:/logtest", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	logger.SetLevel(logger.ERROR)

	for i := 10000; i > 0; i-- {
		go log()
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)
}
