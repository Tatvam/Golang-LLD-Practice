package logger

type ILoggerHandler interface {
	LogEvent(msg string, logType LogType)
	FlushLog(msg string)
}
