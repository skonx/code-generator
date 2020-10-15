package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	cg "github.com/trendev/go-pwdgen/generator"
)

const dir = "/tmp/secret-code/"
const file = dir + "secret.json"   // file with most recent secret
const logfile = dir + "secret.log" // log file

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

func main() {

	fmt.Println("Control command line arguments and set global settings :")
	codeSize := getFromEnvVar("CODE_SIZE", "code size", 64)
	delay := getFromEnvVar("DELAY", "delay (ms)", 200)

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
