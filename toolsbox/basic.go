package toolsbox

import (
	"fmt"
	"log"
	"os"
)

const (
	ROOTPATH = "/Users/oswaldo/dev/golang/octools/"
)

var processlog = loginit("process")

func loginit(logname string) *log.Logger {
	filepath := ROOTPATH + "logs/" + logname + ".log"
	fe, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	logger := log.New(fe, "["+logname+"]", log.LUTC|log.Lshortfile|log.LstdFlags)
	return logger
}
