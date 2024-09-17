package logger

type InfoLogHandler struct {
	logType LogType
	next    ILoggerHandler
}

func InitInfoLogHandler() *InfoLogHandler {
	return &InfoLogHandler{
		logType: INFO,
	}
}

func (i *InfoLogHandler) SetNext(l ILoggerHandler) {
	i.next = l
}

func (i *InfoLogHandler) LogEvent(msg string, logType LogType) {
	if logType != i.logType && i.next != nil {
		i.next.LogEvent(msg, logType)
	}
	if i.next == nil {
		// return error
	}
	i.FlushLog(msg)

}

func (i *InfoLogHandler) FlushLog(msg string) {
	// Add Observer here
	println(msg)
}
