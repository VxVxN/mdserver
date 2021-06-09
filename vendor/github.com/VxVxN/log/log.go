package log

import (
	"fmt"
	"io/ioutil"
	slog "log"
	"os"
)

type LevelLog int

const (
	// CommonLog mode. Normal logging mode. uses info, warning, error, fatal logs.
	CommonLog LevelLog = iota
	// DebugLog mode. Usually only enabled when debugging. Very verbose logging.
	DebugLog
	// TraceLog mode. Designates finer-grained informational events than the Debug.
	TraceLog
)

var (
	// Trace level. Designates finer-grained informational events than the Debug.
	Trace *slog.Logger
	// Debug level. Usually only enabled when debugging.
	Debug *slog.Logger
	// Info level. Records that inform about what is happening in the program.
	Info *slog.Logger
	// Warning level. Non-critical entries that deserve attention.
	Warning *slog.Logger
	// Error level. Logs. Used for errors that should definitely be noted.
	Error *slog.Logger
	// Fatal level, highest level of severity after it, you need to crash the program.
	Fatal *slog.Logger
)

// Init Initializes the logs. Creates a log file with the specified logging level.
func Init(pathLog string, lvlLog LevelLog, isTest bool) error {
	var err error
	var file *os.File

	writer := ioutil.Discard

	if !isTest {
		file, err = os.OpenFile(pathLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("can't open log file: %v", err)
		}
		writer = file
	}

	Info = slog.New(writer,
		"INFO:    ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Warning = slog.New(writer,
		"WARNING: ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Error = slog.New(writer,
		"ERROR:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Fatal = slog.New(writer,
		"FATAL:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Debug = slog.New(ioutil.Discard,
		"DEBUG:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	Trace = slog.New(ioutil.Discard,
		"TRACE:   ",
		slog.Ldate|slog.Ltime|slog.Lshortfile)

	switch lvlLog {
	case DebugLog:
		Debug = slog.New(writer,
			"DEBUG:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
	case TraceLog:
		Debug = slog.New(writer,
			"DEBUG:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
		Trace = slog.New(writer,
			"TRACE:   ",
			slog.Ldate|slog.Ltime|slog.Lshortfile)
	default:
	}
	return nil
}
