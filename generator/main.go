package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/trendev/pwdgen/generator"
)

const path = "/tmp/secret-code/"
const filename = "secret.json"
const filepath = path + filename
const backupFilepath = path + "secret.log"

func init() {

	fmt.Printf("\033[1;35m*** Starting %s ***\033[0m\n", os.Args[0])

	controlEnvs()

	rand.Seed(time.Now().UnixNano())

	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("\033[1;33mDirectory \"%s\" already exists\033[0m\n", path)
		} else {
			check(err)
		}
	} else {
		fmt.Printf("\033[1;32mDirectory \"%s\" did not exist and is now created :)\033[0m\n", path)
	}

	f, err := os.Create(backupFilepath)
	check(err)
	defer f.Close()

}

func randomGenerator(size int) *Secret {
	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		sb.WriteByte(generator.Base[rand.Intn(len(generator.Base))])
	}

	secret := Secret{
		Code:      sb.String(),
		Timestamp: time.Now().UnixNano(),
		Delay:     delay,
	}

	return &secret
}

func saveInFile(secret *Secret) {
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()

	bkf, err := os.OpenFile(backupFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer bkf.Close()

	jsonObject, err := json.MarshalIndent(&secret, "", "    ") //indent with 4 spaces
	check(err)

	f.Write(jsonObject)
	f.Sync()

	logger := log.New(bkf, "", log.LstdFlags)
	logger.Println(fmt.Sprintf("- %d : %s", secret.Timestamp, secret.Code))

	fmt.Printf("%d : code \033[1;31m%s\033[0m saved in file \033[1;34m%s\033[0m\n", secret.Timestamp, secret.Code, filepath)
}

func main() {
	for {
		saveInFile(randomGenerator(codeSize))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
