package main

import (
	"fmt"
	"math/rand"
	"time"
)

const base = "0123456789_-.AZERTYUIOPMLKJHGFDSQWXCVBNazertyuiopmlkjhgfdsqwxcvbn"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGenerator(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = base[rand.Intn(len(base))]
	}
	return string(b)
}

func main() {
	fmt.Println(randomGenerator(16))
}
