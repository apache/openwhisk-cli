/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package common

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"unicode"
)

const (
	PropDisplayCert       = "client cert"
	PropDisplayKey        = "Client key"
	PropDisplayAuth       = "whisk auth"
	PropDisplayAPIHost    = "whisk API host"
	PropDisplayAPIVersion = "whisk API version"
	PropDisplayNamespace  = "whisk namespace"
	PropDisplayCLIVersion = "whisk CLI version"
	PropDisplayAPIBuild   = "whisk API build"
	PropDisplayAPIBuildNo = "whisk API build number"
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

func RemoveRedundentSpaces(in string) (out string) {
	white := false
	for _, c := range in {
		if unicode.IsSpace(c) {
			if !white {
				out = out + " "
			}
			white = true
		} else {
			out = out + string(c)
			white = false
		}
	}
	out = strings.TrimSpace(out)
	return
}

func GetTestActionFilename(fileName string) string {
	return GetRepoPath() + "/tests/src/dat/" + fileName
}

func GetRepoPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/apache/openwhisk-cli"
}

func GetBinPath() string {
	_, goFileName, _, _ := runtime.Caller(1)
	//  Yes, this assumes we're using the official build script.  I haven't
	//  figured out a better approach yet given the panoply of options.
	//  Maybe some sort of Go search path?
	return path.Join(path.Dir(goFileName), "../../../../build")
}

type InvalidArg struct {
	Cmd []string
	Err string
}
