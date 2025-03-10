package logger

import (
	"errors"
	"github.com/saichler/shared/go/share/queues"
	"github.com/saichler/types/go/common"
	"testing"
	"time"
)

type ILogMethod interface {
	Log(level common.LogLevel, msg string)
}

type LoggerImpl struct {
	queue      *queues.Queue
	logMethods []ILogMethod
	logLevel   common.LogLevel
}

type LoggerEntry struct {
	anys []interface{}
	t    int64
	l    common.LogLevel
}

func NewLoggerImpl(logMethods ...ILogMethod) *LoggerImpl {
	logImpl := &LoggerImpl{}
	logImpl.logMethods = logMethods
	logImpl.queue = queues.NewQueue("Logger Queue", 50000)
	go logImpl.processQueue()
	return logImpl
}

func (loggerImpl *LoggerImpl) processQueue() {
	for {
		entry := loggerImpl.queue.Next().(*LoggerEntry)
		str := FormatLog(entry.l, entry.t, entry.anys...)
		if len(loggerImpl.logMethods) == 1 {
			loggerImpl.logMethods[0].Log(entry.l, str)
		} else if len(loggerImpl.logMethods) == 2 {
			loggerImpl.logMethods[0].Log(entry.l, str)
			loggerImpl.logMethods[1].Log(entry.l, str)
		} else if len(loggerImpl.logMethods) == 3 {
			loggerImpl.logMethods[0].Log(entry.l, str)
			loggerImpl.logMethods[1].Log(entry.l, str)
			loggerImpl.logMethods[2].Log(entry.l, str)
		}
	}
}

func newEntry(l common.LogLevel, anys ...interface{}) *LoggerEntry {
	return &LoggerEntry{
		t:    time.Now().Unix(),
		l:    l,
		anys: anys,
	}
}

func (loggerImpl *LoggerImpl) Empty() bool {
	return loggerImpl.queue.Size() == 0
}

func (loggerImpl *LoggerImpl) Trace(anys ...interface{}) {
	if loggerImpl.logLevel > common.Trace_Level {
		return
	}
	loggerImpl.queue.Add(newEntry(common.Trace_Level, anys...))
}

func (loggerImpl *LoggerImpl) Debug(anys ...interface{}) {
	if loggerImpl.logLevel > common.Debug_Level {
		return
	}
	loggerImpl.queue.Add(newEntry(common.Debug_Level, anys...))
}

func (loggerImpl *LoggerImpl) Info(anys ...interface{}) {
	if loggerImpl.logLevel > common.Info_Level {
		return
	}
	loggerImpl.queue.Add(newEntry(common.Info_Level, anys...))
}

func (loggerImpl *LoggerImpl) Warning(anys ...interface{}) {
	if loggerImpl.logLevel > common.Warning_Level {
		return
	}
	loggerImpl.queue.Add(newEntry(common.Warning_Level, anys...))
}

func (loggerImpl *LoggerImpl) Error(anys ...interface{}) error {
	anys = append(anys, FileAndLine(".go", false))
	loggerImpl.queue.Add(newEntry(common.Error_Level, anys...))
	err := FormatLog(common.Error_Level, time.Now().Unix(), anys...)
	return errors.New(err)
}

func (loggerImpl *LoggerImpl) Fail(t interface{}, args ...interface{}) {
	args = append(args, FileAndLine("tests", true))
	loggerImpl.Error(args...)
	ts, ok := t.(*testing.T)
	if ok {
		ts.Fail()
	}
}

func (loggerImpl *LoggerImpl) SetLogLevel(level common.LogLevel) {
	loggerImpl.logLevel = level
}
