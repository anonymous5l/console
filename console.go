package console

import (
	"fmt"
	"time"
)

type log struct {
	format string
	prefix string
	argv   []interface{}
}

var logChan chan *log

func _log(format string, prefix string, a ...interface{}) (n int, err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s ", now, prefix)
	return fmt.Printf(format+"\n", a...)
}

func init() {
	logChan = make(chan *log)
	go safeLoop()
}

func safeLoop() {
	for {
		log := <-logChan
		_log(log.format, log.prefix, log.argv...)
	}
}

func Ok(format string, a ...interface{}) {
	logChan <- &log{
		format,
		"\x1B[1;32m   [OK]\x1B[0m",
		a,
	}
}

func Log(format string, a ...interface{}) {
	logChan <- &log{
		format,
		"\x1B[1;34m  [LOG]\x1B[0m",
		a,
	}
}

func Err(format string, a ...interface{}) {
	logChan <- &log{
		format,
		"\x1B[1;31m  [ERR]\x1B[0m",
		a,
	}
}

func Warn(format string, a ...interface{}) {
	logChan <- &log{
		format,
		"\x1B[1;33m [WARN]\x1B[0m",
		a,
	}
}

func Debug(format string, a ...interface{}) {
	logChan <- &log{
		format,
		"\x1B[1;33m[DEBUG]\x1B[0m",
		a,
	}
}
