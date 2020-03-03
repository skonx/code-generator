package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const base = "0123456789_-.AZERTYUIOPMLKJHGFDSQWXCVBNazertyuiopmlkjhgfdsqwxcvbn"

const path = "/tmp/secret-code/"
const filename = "password.key"
const filepath = path + filename

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {

	//os.Getenv("CODE_LENGTH")

	rand.Seed(time.Now().UnixNano())

	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("\033[1;33mDirectory \"%s\" already exists\033[0m\n", path)
		} else {
			check(err)
		}
	} else {
		fmt.Printf("\033[1;32mDirectory \"%s\" did not exist and will be created :)\033[0m\n", path)
	}
}

func randomGenerator(size int) string {
	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		sb.WriteByte(base[rand.Intn(len(base))])
	}
	return sb.String()
}

func saveInFile(code string) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	f.WriteString(code)
	f.Sync()
	fmt.Printf("%s : code \033[1;31m%s\033[0m saved in file \033[1;34m%s\033[0m\n", time.Now().String(), code, filepath)
}

func main() {
	for {
		saveInFile(randomGenerator(64))
		time.Sleep(2 * time.Second)
	}
}
