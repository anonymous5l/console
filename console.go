package console

import (
	"fmt"
	"os"
	// "sync"
	"time"
)

// type log struct {
// 	format string
// 	prefix string
// 	argv   []interface{}
// }

// var logChan chan *log
// var mu sync.Mutex
// var cond *sync.Cond

type ConsoleLevel int

const (
	ALL   ConsoleLevel = 0
	DEBUG ConsoleLevel = 1
	WARN  ConsoleLevel = 1 << 1
	OK    ConsoleLevel = 1 << 2
	LOG   ConsoleLevel = 1 << 3
	ERR   ConsoleLevel = 1 << 4
	FATAL ConsoleLevel = 1 << 5
)

var mLevel ConsoleLevel
var mColor bool = true
var mFd = os.Stdout

func _log(format string, color int, prefix string, a ...interface{}) (n int, err error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	args := make([]interface{}, 2)
	args[0] = now

	if mColor {
		args[1] = fmt.Sprintf("\x1B[1;%dm%s\x1B[0m", color, prefix)
	} else {
		args[1] = prefix
	}
	args = append(args, a...)

	return fmt.Fprintln(mFd,fmt.Sprintf("[%s] %s "+format, args...))
}

// func init() {
// 	cond = sync.NewCond(&mu)
// 	mu.Lock()
// 	logChan = make(chan *log)
// 	go safeLoop()
// }

// func safeLoop() {
// 	for {
// 		log, ok := <-logChan
// 		if !ok {
// 			break
// 		}
// 		_log(log.format, log.prefix, log.argv...)
// 	}
// 	cond.Signal()
// }

func Ok(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;32m   [OK]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&OK == OK {
		_log(format, 32, "   [OK]", a...)
	}
}

func Log(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;34m [INFO]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&LOG == LOG {
		_log(format, 34, " [INFO]", a...)
	}
}

func Err(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;31m  [ERR]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&ERR == ERR {
		_log(format, 31, "  [ERR]", a...)
	}
}

func Fatal(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;35m[FATAT]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&FATAL == FATAL {
		_log(format, 35, "[FATAL]", a...)
	}
}

func Warn(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;33m [WARN]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&WARN == WARN {
		_log(format, 33, " [WARN]", a...)
	}
}

func Debug(format string, a ...interface{}) {
	// logChan <- &log{
	// 	format,
	// 	"\x1B[1;36m[DEBUG]\x1B[0m",
	// 	a,
	// }
	if mLevel == ALL || mLevel&DEBUG == DEBUG {
		_log(format, 36, "[DEBUG]", a...)
	}
}

func SetLevel(level ConsoleLevel) {
	mLevel = level
}

func SetColor(color bool) {
	mColor = color
}

func SetFD(f *os.File) {
	mFd = f
}

// func Abort() {
// 	close(logChan)
// 	cond.Wait()
// }
