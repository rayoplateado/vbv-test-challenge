package log

import (
	"log"
	"os"
)

func Info(message string) {
	writer := log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
	writer.Println(message)
}

func Error(message string) {
	writer := log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	writer.Println(message)
}
