package trace

import (
	"log"
	"time"
)

func traceLog(msg string)func(){
	start:=time.Now()
	log.Printf("enter %s",msg)
	return func(){
		log.Printf("exit %s (%s)",msg,time.Since(start))
	}
}

func BigSlowOperation(){
	defer traceLog("BigSlowOperation")()
	time.Sleep(3*time.Second)
}