package logging

//go:generate gostub Logger

type Logger interface {
	Printf(msg string, args ...interface{})
}

type NopLogger struct{}

func (l NopLogger) Printf(msg string, args ...interface{}) {
}

var DefaultLogger Logger = NopLogger{}

func Printf(msg string, args ...interface{}) {
	DefaultLogger.Printf(msg, args...)
}
