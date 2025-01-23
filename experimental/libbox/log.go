package libbox

import (
	"github.com/metacubex/mihomo/log"
)

// LogLevel represents the logging level
type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

// Logger interface for mobile platforms
type Logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

// MobileLogger implements mobile logging
type MobileLogger struct {
	level LogLevel
}

func (l *MobileLogger) Trace(message string) {
	if l.level >= LogLevelDebug {
		log.Debugln(message)
	}
}

func (l *MobileLogger) Debug(message string) {
	if l.level >= LogLevelDebug {
		log.Debugln(message)
	}
}

func (l *MobileLogger) Info(message string) {
	if l.level >= LogLevelInfo {
		log.Infoln(message)
	}
}

func (l *MobileLogger) Warn(message string) {
	if l.level >= LogLevelWarning {
		log.Warnln(message)
	}
}

func (l *MobileLogger) Error(message string) {
	if l.level >= LogLevelError {
		log.Errorln(message)
	}
}

func (l *MobileLogger) Fatal(message string) {
	log.Fatalln(message)
}

func InitializeLog(level LogLevel) {
	// Set log level based on input
	switch level {
	case LogLevelError:
		log.SetLevel(log.ERROR)
	case LogLevelWarning:
		log.SetLevel(log.WARNING)
	case LogLevelInfo:
		log.SetLevel(log.INFO)
	case LogLevelDebug:
		log.SetLevel(log.DEBUG)
	}
}

func LogInfo(format string, args ...interface{}) {
	log.Infoln(format, args...)
}

func LogWarning(format string, args ...interface{}) {
	log.Warnln(format, args...)
}

func LogError(format string, args ...interface{}) {
	log.Errorln(format, args...)
}

func LogDebug(format string, args ...interface{}) {
	log.Debugln(format, args...)
}
