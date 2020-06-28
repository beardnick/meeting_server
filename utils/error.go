package utils

import (
	"runtime"
	"strconv"

	"github.com/pkg/errors"
)

// 业务逻辑错误
func BizError(err error) error {
	if err == nil {
		return nil
	}
	//_, srcName, line, _ := runtime.Caller(1) // (1)
	// [源文件名:行号]
	//prefix := "[" + srcName + ":" + strconv.Itoa(line) + "] "
	return errors.Wrap(err, "biz")
}

// 系统错误
func SysError(err error) error {
	if err == nil {
		return nil
	}
	_, srcName, line, _ := runtime.Caller(1) // (1)
	// [源文件名:行号]
	prefix := "[" + srcName + ":" + strconv.Itoa(line) + "] "
	return errors.Wrap(err, prefix)
}

// 输入参数错误
func ParamError(err error) error {
	if err == nil {
		return nil
	}
	// #NOTE: 20-06-28 wrap同时添加prefix和stacktrace,
	// 只有用%+v打印的时候才可以看到stacktrace
	return errors.Wrap(err, "参数错误")
}
