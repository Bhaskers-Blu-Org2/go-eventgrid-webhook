package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Microsoft/go-eventgrid-handler/logb"
)

// port to listen on (can be changed with -port option)
var port = 8080

// main - application startup
func main() {
	// read the flags
	logPath := flag.String("logpath", "./", "path to write log files")
	p := flag.Int("port", port, "TCP port to listen on")
	flag.Parse()

	// validate the port flag
	// use default port if bad port
	if *p >= 0 && *p <= 65535 {
		// set the listen port
		port = *p
	}

	// setup the log multi writer
	if err := setupLogs(*logPath); err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on port: ", port)

	// run the server
	// this blocks until the http server shuts down
	if err := runServer(port); err != nil {
		log.Println("ERROR", err)
	}

	log.Println("Server Exit")
}

// setupLogs - sets up the multi writer for the log files
func setupLogs(logPath string) error {
	// make the log directory if it doesn't exist
	if err := os.MkdirAll(logPath, 0666); err != nil {
		return err
	}

	// prepend date and time to log entries
	log.SetFlags(log.Ldate | log.Ltime)

	fileName := logPath + "app"

	// use instance ID to differentiate log files between instances in App Services
	if iid := os.Getenv("WEBSITE_ROLE_INSTANCE_ID"); iid != "" {
		fileName += "_" + strings.TrimSpace(iid)
	}

	fileName += ".log"

	// open the log file

	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	// setup a multiwriter to log to file and stdout
	wrt := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(wrt)
	logb.Logger.SetOutput(wrt)

	return nil
}
