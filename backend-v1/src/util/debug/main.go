package debug

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

func logBuilder(context string, message string) string {
	return fmt.Sprintf("[%s] // %s", context, message)
}

func Info(context string, message string) {
	msg := fmt.Sprintf("%s - %s - Info - %s\n", time.Now().String(), context,
		message)

	log.Info(logBuilder(context, message))
	fmt.Printf(msg)
}

func Warning(context string, message string) {
	msg := fmt.Sprintf("%s - %s - Warning - %s\n", time.Now().String(), context,
		message)

	log.Warn(logBuilder(context, message))
	fmt.Printf(msg)
}

func Error(context string, message string) {
	msg := fmt.Sprintf("%s - %s - Error - %s\n", time.Now().String(), context,
		message)

	log.Error(logBuilder(context, message))
	fmt.Errorf(msg)
}