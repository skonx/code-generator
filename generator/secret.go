package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type secret struct {
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"` // in nanoseconds
	Delay     int    `json:"delay"`     // in seconds, not really used... just for fun
}

func (s secret) store(file, logfile string) {
	// creates the secret file
	f, err := os.Create(file)
	checkError(err)
	defer f.Close()

	// open the logfile
	lf, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer lf.Close()

	jsonObject, err := json.MarshalIndent(&s, "", "\t") //indent with single tabs
	checkError(err)

	f.Write(jsonObject)
	f.Sync()

	logger := log.New(lf, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("- %d : %s", s.Timestamp, s.Code))

	fmt.Printf("%d : code \033[1;31m%s\033[0m saved in file \033[1;34m%s\033[0m\n", s.Timestamp, s.Code, file)
}
