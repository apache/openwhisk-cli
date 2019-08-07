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

package main

import (
	"fmt"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"os"

	"github.com/apache/openwhisk-cli/commands"
	"github.com/apache/openwhisk-cli/wski18n"
	"github.com/apache/openwhisk-client-go/whisk"
)

// CLI_BUILD_TIME holds the time of the CLI build.  During gradle builds,
// this value will be overwritten via the command:
//     go build -ldflags "-X main.CLI_BUILD_TIME=nnnnn"   // nnnnn is the new timestamp
var CLI_BUILD_TIME string = "not set"

var cliDebug = os.Getenv("WSK_CLI_DEBUG") // Useful for tracing init() code

var T goi18n.TranslateFunc

func init() {
	if len(cliDebug) > 0 {
		whisk.SetDebug(true)
	}

	T = wski18n.T

	// Rest of CLI uses the Properties struct, so set the build time there
	commands.Properties.CLIVersion = CLI_BUILD_TIME
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println(T("Application exited unexpectedly"))
		}
	}()

	if err := commands.Execute(); err != nil {
		commands.ExitOnError(err)
	}
	return
}
