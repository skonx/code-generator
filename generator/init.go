package main

import (
	"fmt"
	"os"
	"strconv"
)

// default and base values
const (
	Base            = "0123456789_-.AZERTYUIOPMLKJHGFDSQWXCVBNazertyuiopmlkjhgfdsqwxcvbn"
	DefaultCodeSize = 64
	DefaultDelay    = 10
)

var codeSize int
var delay int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func controlIntEnvVar(env string, msg string, value *int, defaultValue int) {
	if size := os.Getenv(env); size != "" {
		s, err := strconv.Atoi(size)
		check(err)
		*value = s
	} else {
		*value = defaultValue
	}

	fmt.Printf("- %s = %d\n", msg, *value)
}

func controlEnvs() {
	fmt.Println("Control command line arguments and set global settings :")
	controlIntEnvVar("CODE_SIZE", "code size", &codeSize, DefaultCodeSize)
	controlIntEnvVar("DELAY", "delay (seconds)", &delay, DefaultDelay)
}
