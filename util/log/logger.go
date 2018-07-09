package log

import (
	"fmt"
	"time"
)

var (
	debugging bool
	inf       = "[INFO]  |"
	err       = "[ERROR] |"
	deb       = "[DEBUG] |"
)

func Setup(debug bool) {
	debugging = debug
}

func Info(format string, args ...interface{}) {
	log(inf, format, args...)
}

func Err(format string, args ...interface{}) {
	log(err, format, args...)
}

func Debug(format string, args ...interface{}) {
	if debugging {
		log(deb, format, args...)
	}
}

func log(prefix, format string, args ...interface{}) {
	fmt.Print(prefix, time.Now().Format(time.Stamp), "| ")
	fmt.Printf(format, args...)
	fmt.Println()
}
