package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	cg "github.com/trendev/go-pwdgen/generator"
)

const dir = "/tmp/secret-code/"
const file = dir + "secret.json"   // file with most recent secret
const logfile = dir + "secret.log" // log file

type secret struct {
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"` // in nanoseconds
	Delay     int    `json:"delay"`     // in seconds, not really used... just for fun
}

func checkError(e error) {
	if e != nil {
		panic(e) // just panic ;)
	}
}

func getFromEnvVar(ev string, l string, dv int) int {
	var cst int
	if v := os.Getenv(ev); v != "" {
		s, err := strconv.Atoi(v)
		checkError(err)
		cst = s
	} else {
		cst = dv
	}

	fmt.Printf("- %s = %d\n", l, cst)
	return cst
}

func init() {
	fmt.Printf("\033[1;35m*** Starting %s ***\033[0m\n", os.Args[0])

	rand.Seed(time.Now().UnixNano())

	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("\033[1;33mDirectory \"%s\" already exists\033[0m\n", dir)
		} else {
			checkError(err)
		}
	} else {
		fmt.Printf("\033[1;32mDirectory \"%s\" did not exist and is now created 👍\033[0m\n", dir)
	}

	// creates the log file
	f, err := os.Create(logfile)
	checkError(err)
	defer f.Close()

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

func main() {

	fmt.Println("Control command line arguments and set global settings :")
	codeSize := getFromEnvVar("CODE_SIZE", "code size", 64)
	delay := getFromEnvVar("DELAY", "delay (ms)", 200)

	for range time.Tick(time.Millisecond * time.Duration(delay)) {
		s := secret{
			cg.Generate(codeSize),
			time.Now().UnixNano(),
			delay,
		}
		s.store(file, logfile)
	}
}
