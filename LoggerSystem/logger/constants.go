package logger

type LogType int

const (
	INFO = iota
	WARN
	ERROR
)

func String(logType LogType) string {
	switch logType {
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return ""
	}
}
