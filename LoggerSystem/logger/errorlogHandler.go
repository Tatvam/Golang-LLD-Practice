package logger

type ErrorLogHandler struct {
	logType LogType
	next    ILoggerHandler
}

func InitErrorLogHandler() *ErrorLogHandler {
	return &ErrorLogHandler{
		logType: ERROR,
	}
}

func (e *ErrorLogHandler) SetNext(logHandler ILoggerHandler) {
	e.next = logHandler
}

func (e *ErrorLogHandler) LogEvent(msg string, logType LogType) {
	if logType != e.logType && e.next != nil {
		e.next.LogEvent(msg, logType)
	}
	if e.next == nil {
		// return error
	}
	e.FlushLog(msg)

}

func (e *ErrorLogHandler) FlushLog(msg string) {
	// Add Observer here
	println(msg)
}
