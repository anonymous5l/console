package console

import (
	"fmt"
	"sync"
	"time"
)

type log struct {
	format string
	prefix string
	argv   []interface{}
}

var logChan chan *log
var mu sync.Mutex
var cond *sync.Cond

func _log(format string, prefix string, a ...interface{}) (n int, err error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	args := make([]interface{}, 2)
	args[0] = now
	args[1] = prefix
	args = append(args, a...)

	return fmt.Printf("[%s] %s "+format+"\n", args...)
}

func init() {
	cond = sync.NewCond(&mu)
	mu.Lock()
	logChan = make(chan *log)
	go safeLoop()
}

func safeLoop() {
	for {
		log, ok := <-logChan
		if !ok {
			break
		}
		_log(log.format, log.prefix, log.argv...)
	}
	cond.Signal()
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

func Abort() {
	close(logChan)
	cond.Wait()
}
