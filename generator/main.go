package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	cg "github.com/trendev/pwdgen/generator"
)

const dir = "/tmp/secret-code/"
const file = dir + "secret.json"   // file with most recent secret
const logfile = dir + "secret.log" // log file

var codeSize int
var delay int // delay (in ms) between each secret creation

func checkError(e error) {
	if e != nil {
		panic(e) // just panic ;)
	}
}

func initConst(
	ev string, // environment var
	l string, // const label
	cst *int, // const to update
	dv int) { // default value
	if v := os.Getenv(ev); v != "" {
		s, err := strconv.Atoi(v)
		checkError(err)
		*cst = s
	} else {
		*cst = dv
	}

	fmt.Printf("- %s = %d\n", l, *cst)
}

func init() {
	fmt.Printf("\033[1;35m*** Starting %s ***\033[0m\n", os.Args[0])

	fmt.Println("Control command line arguments and set global settings :")
	initConst("CODE_SIZE", "code size", &codeSize, 64)
	initConst("DELAY", "delay (ms)", &delay, 200)

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

func main() {
	for {
		s := secret{
			cg.Generate(codeSize),
			time.Now().UnixNano(),
			delay,
		}
		s.store(file, logfile)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
