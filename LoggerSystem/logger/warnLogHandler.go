package logger

type WarnLogHandler struct {
	logType LogType
	next    ILoggerHandler
}

func InitWarnLogHandler() *WarnLogHandler {
	return &WarnLogHandler{
		logType: WARN,
	}
}

func (w *WarnLogHandler) SetNext(logHandler ILoggerHandler) {
	w.next = logHandler
}

func (w *WarnLogHandler) LogEvent(msg string, logType LogType) {
	if logType != w.logType && w.next != nil {
		w.next.LogEvent(msg, logType)
	}
	if w.next == nil {
		// return error
	}
	w.FlushLog(msg)

}

func (w *WarnLogHandler) FlushLog(msg string) {
	// Add Observer here
	println(msg)
}
