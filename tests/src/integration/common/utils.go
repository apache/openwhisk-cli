package common

import (
	"fmt"
	"os"
	"regexp"
	"io"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func CreateFile(filePath string) {
	var _, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		var file, err = os.Create(filePath)
		checkError(err)
		defer file.Close()
	}
	return
}

func ReadFile(filePath string) string {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	return string(text)
}

func WriteFile(filePath string, lines []string) {
	var file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	for _, each := range lines {
		_, err = file.WriteString(each + "\n")
		checkError(err)
	}

	err = file.Sync()
	checkError(err)
}

func DeleteFile(filePath string) {
	var err = os.Remove(filePath)
	checkError(err)
}

func RemoveRedundentSpaces(str string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(str, "")
	return re_inside_whtsp.ReplaceAllString(final, " ")
}

func GetTestActionFilename(fileName string) string {
	return os.Getenv("GOPATH") + "/src/github.com/openwhisk/openwhisk-cli/tests/src/dat/" + fileName
}

type InvalidArg struct {
	Cmd []string
	Err string
}