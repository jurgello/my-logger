package mylogger

import "log"

func LogInfo(message string) {
	log.Printf("INFO -%v\n", message)
}

func LogWarning(message string) {
	log.Printf("WARN -%v\n", message)
}

func LogError(message string) {
	log.Printf("ERROR -%v\n", message)
}
