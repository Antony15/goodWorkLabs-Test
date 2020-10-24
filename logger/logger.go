package logger

import (
	"go/build"
	"log"
	"os"

	"github.com/Antony15/goodWorkLabs-Test/constants"
)

var (
	Log *log.Logger
)

// Initializes the log file to print logs
func init() {
	// set location of log file
	var logpath = build.Default.GOPATH + constants.LogFile

	var file, err = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
