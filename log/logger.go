// Author: Qingshan Luo <edoger@qq.com>
package log

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/logx"
)

var defaultLogger Logger = logger(0)

func Log() Logger { return defaultLogger }
func New() Logger { return logger(0) }

type Logger interface {
	Print(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Println(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	Printf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type logger int

func (logger) Print(args ...interface{}) { logx.Info(args...) }
func (logger) Debug(args ...interface{}) { logx.Info(args...) }
func (logger) Info(args ...interface{})  { logx.Info(args...) }
func (logger) Warn(args ...interface{})  { logx.Info(args...) }
func (logger) Error(args ...interface{}) { logx.Error(args...) }
func (logger) Fatal(args ...interface{}) { logx.Error(args...) }
func (logger) Panic(args ...interface{}) { logx.Error(args...) }

func (logger) Println(args ...interface{}) { logx.Info(sprintln(args...)) }
func (logger) Debugln(args ...interface{}) { logx.Info(sprintln(args...)) }
func (logger) Infoln(args ...interface{})  { logx.Info(sprintln(args...)) }
func (logger) Warnln(args ...interface{})  { logx.Info(sprintln(args...)) }
func (logger) Errorln(args ...interface{}) { logx.Error(sprintln(args...)) }
func (logger) Fatalln(args ...interface{}) { logx.Error(sprintln(args...)) }
func (logger) Panicln(args ...interface{}) { logx.Error(sprintln(args...)) }

func (logger) Printf(format string, args ...interface{}) { logx.Infof(format, args...) }
func (logger) Debugf(format string, args ...interface{}) { logx.Infof(format, args...) }
func (logger) Infof(format string, args ...interface{})  { logx.Infof(format, args...) }
func (logger) Warnf(format string, args ...interface{})  { logx.Infof(format, args...) }
func (logger) Errorf(format string, args ...interface{}) { logx.Errorf(format, args...) }
func (logger) Fatalf(format string, args ...interface{}) { logx.Errorf(format, args...) }
func (logger) Panicf(format string, args ...interface{}) { logx.Errorf(format, args...) }

func sprintln(args ...interface{}) string {
	s := fmt.Sprintln(args...)
	return s[:len(s)-1]
}

func Print(args ...interface{}) { defaultLogger.Print(args...) }
func Debug(args ...interface{}) { defaultLogger.Debug(args...) }
func Info(args ...interface{})  { defaultLogger.Info(args...) }
func Warn(args ...interface{})  { defaultLogger.Warn(args...) }
func Error(args ...interface{}) { defaultLogger.Error(args...) }
func Fatal(args ...interface{}) { defaultLogger.Fatal(args...) }
func Panic(args ...interface{}) { defaultLogger.Panic(args...) }

func Println(args ...interface{}) { defaultLogger.Println(args...) }
func Debugln(args ...interface{}) { defaultLogger.Debugln(args...) }
func Infoln(args ...interface{})  { defaultLogger.Infoln(args...) }
func Warnln(args ...interface{})  { defaultLogger.Warnln(args...) }
func Errorln(args ...interface{}) { defaultLogger.Errorln(args...) }
func Fatalln(args ...interface{}) { defaultLogger.Fatalln(args...) }
func Panicln(args ...interface{}) { defaultLogger.Panicln(args...) }

func Printf(format string, args ...interface{}) { defaultLogger.Printf(format, args...) }
func Debugf(format string, args ...interface{}) { defaultLogger.Debugf(format, args...) }
func Infof(format string, args ...interface{})  { defaultLogger.Infof(format, args...) }
func Warnf(format string, args ...interface{})  { defaultLogger.Warnf(format, args...) }
func Errorf(format string, args ...interface{}) { defaultLogger.Errorf(format, args...) }
func Fatalf(format string, args ...interface{}) { defaultLogger.Fatalf(format, args...) }
func Panicf(format string, args ...interface{}) { defaultLogger.Panicf(format, args...) }
