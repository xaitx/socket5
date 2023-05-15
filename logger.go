package socket5

import (
	"log"
)

// Logger 日志接口
type Logger interface {
	// Info 打印信息级别日志
	Info(args ...interface{})

	// Infof 格式化打印信息级别日志
	Infof(format string, args ...interface{})

	// Error 打印错误级别日志
	Error(args ...interface{})

	// Errorf 格式化打印错误级别日志
	Errorf(format string, args ...interface{})
}

// DefaultLogger 默认日志
type DefaultLogger struct{}

// Info 打印信息级别日志
func (l *DefaultLogger) Info(args ...interface{}) {
	log.Print(args...)
}

// Infof 格式化打印信息级别日志
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Error 打印错误级别日志
func (l *DefaultLogger) Error(args ...interface{}) {
	log.Print(args...)
}

// Errorf 格式化打印错误级别日志
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// CustomLogger 自定义日志记录器
type CustomLogger struct {
	logger Logger
}

// NewCustomLogger 使用自定义日志记录器创建一个新的 CustomLogger
func NewCustomLogger(logger Logger) *CustomLogger {
	return &CustomLogger{logger: logger}
}

// Info 打印信息级别日志
func (s *CustomLogger) Info(args ...interface{}) {
	s.logger.Info(args...)
}

// Infof 格式化打印信息级别日志
func (s *CustomLogger) Infof(format string, args ...interface{}) {
	s.logger.Infof(format, args...)
}

// Error 打印错误级别日志
func (s *CustomLogger) Error(args ...interface{}) {
	s.logger.Error(args...)
}

// Errorf 格式化打印错误级别日志
func (s *CustomLogger) Errorf(format string, args ...interface{}) {
	s.logger.Errorf(format, args...)
}

// DefaultLoggerInstance 默认日志实例
var DefaultLoggerInstance = NewCustomLogger(&DefaultLogger{})

// SetLogger 设置默认日志记录器
func SetLogger(logger Logger) {
	DefaultLoggerInstance = NewCustomLogger(logger)
}

// Info 打印信息级别日志
func Info(args ...interface{}) {
	DefaultLoggerInstance.Info(args...)
}

// Infof 格式化打印信息级别日志
func Infof(format string, args ...interface{}) {
	DefaultLoggerInstance.Infof(format, args...)
}

// Error 打印错误级别日志
func Error(args ...interface{}) {
	DefaultLoggerInstance.Error(args...)
}

// Errorf 格式化打印错误级别日志
func Errorf(format string, args ...interface{}) {
	DefaultLoggerInstance.Errorf(format, args...)
}
