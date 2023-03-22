package toolsbox

import (
	"fmt"
	"log"
	"os"
)

func LogInit(logname, filepath string) *log.Logger {
	fe, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	logger := log.New(fe, "["+logname+"]", log.LUTC|log.Lshortfile|log.LstdFlags)
	return logger
}
