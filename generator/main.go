package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	cg "github.com/trendev/pwdgen/generator"
)

const path = "/tmp/secret-code/"
const filename = "secret.json"
const filepath = path + filename
const backupFilepath = path + "secret.log"

var codeSize int
var delay int

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getEnvVar(env string, msg string, value *int, dv int) {
	if v := os.Getenv(env); v != "" {
		s, err := strconv.Atoi(v)
		checkError(err)
		*value = s
	} else {
		*value = dv
	}

	fmt.Printf("- %s = %d\n", msg, *value)
}

func init() {
	fmt.Printf("\033[1;35m*** Starting %s ***\033[0m\n", os.Args[0])

	fmt.Println("Control command line arguments and set global settings :")
	getEnvVar("CODE_SIZE", "code size", &codeSize, 64)
	getEnvVar("DELAY", "delay (ms)", &delay, 200)

	rand.Seed(time.Now().UnixNano())

	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("\033[1;33mDirectory \"%s\" already exists\033[0m\n", path)
		} else {
			checkError(err)
		}
	} else {
		fmt.Printf("\033[1;32mDirectory \"%s\" did not exist and is now created 👍\033[0m\n", path)
	}

	f, err := os.Create(backupFilepath)
	checkError(err)
	defer f.Close()

}

type secret struct {
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"` // nanosecond
	Delay     int    `json:"delay"`     // second, just for fun
}

func (s secret) store() {
	f, err := os.Create(filepath)
	checkError(err)
	defer f.Close()

	bkf, err := os.OpenFile(backupFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer bkf.Close()

	jsonObject, err := json.MarshalIndent(&s, "", "\t") //indent with single tab
	checkError(err)

	f.Write(jsonObject)
	f.Sync()

	logger := log.New(bkf, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("- %d : %s", s.Timestamp, s.Code))

	fmt.Printf("%d : code \033[1;31m%s\033[0m saved in file \033[1;34m%s\033[0m\n", s.Timestamp, s.Code, filepath)
}

func main() {
	for {
		s := secret{
			cg.Generate(codeSize),
			time.Now().UnixNano(),
			delay,
		}
		s.store()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
