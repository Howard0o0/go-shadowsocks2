package tinylog

import (
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// LogDebug output debug level msg to terminal
func LogDebug(format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}
