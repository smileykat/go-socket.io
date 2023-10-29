package logger

import (
	"fmt"
	"log"
	"strings"
)

func Error(msg string, err error) {
	//"golang.org/x/exp/slog"
	//Log.Error(msg, "err", err.Error())
	log.Printf("[ERROR] %v\n", err)
}

func Info(msg string, args ...interface{}) {
	//Log.Info(msg, args...)
	sb := strings.Builder{}
	sb.WriteString(msg)
	if len(args) > 0 {
		sb.WriteString(": ")
		for i := 0; i < len(args); i += 2 {
			if i > 0 {
				sb.WriteString(", ")
			}
			key, val := args[i], interface{}("")
			if i+1 < len(args) {
				val = args[i+1]
			}
			sb.WriteString(fmt.Sprintf("%v=%v", key, val))
		}
	}
	log.Printf("[INFO] %s\n", sb.String())
}
