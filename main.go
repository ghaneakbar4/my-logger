package mylogger

import "log"

//LogInfo برای تست
func LogInfo(msg string) {
	log.Printf("log is-%v",msg)
}