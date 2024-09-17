package logger

import "sync"

var mu sync.Mutex
var logger *LoggerInstance
var once sync.Once

type LoggerInstance struct {
	logHandler ILoggerHandler
	// Add Handler
}

// 1st way Aggressive Locking
func GetInstanceAgg() *LoggerInstance {
	mu.Lock()
	defer mu.Unlock()
	if logger == nil {
		return &LoggerInstance{
			// Initialize Handler
		}
	}
	return logger
}

// 2nd Way Check-Lock-Check (Better approach)

func GetInstanceClc() *LoggerInstance {
	if logger == nil {
		mu.Lock()
		defer mu.Unlock()
		if logger == nil {
			return &LoggerInstance{
				// Initialize Handler
			}
		}
	}
	return logger
}

// 3rd Way using sync.Once (Best Way)

func GetInstance() *LoggerInstance {
	once.Do(func() {
		logger = &LoggerInstance{
			logHandler: initLogHandler(),
		}
	})
	return logger
}

func (l *LoggerInstance) Warn(msg string) {
	l.createLog(WARN, msg)
}

func (l *LoggerInstance) Error(msg string) {
	l.createLog(ERROR, msg)
}

func (l *LoggerInstance) Info(msg string) {
	l.createLog(INFO, msg)
}

func (l *LoggerInstance) createLog(logType LogType, msg string) {
	l.logHandler.LogEvent(msg, logType)
}

func initLogHandler() *InfoLogHandler {
	infoLogHandler := InitInfoLogHandler()
	warnLogHandler := InitWarnLogHandler()
	errorLogHandler := InitErrorLogHandler()

	infoLogHandler.SetNext(warnLogHandler)
	warnLogHandler.SetNext(errorLogHandler)

	return infoLogHandler

}
