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
	"github.com/apache/openwhisk-client-go/whisk"
	"os"
	"os/exec"
)

const cmd = "wsk"
const arg = "-i"

type Wsk struct {
	Path     string
	Arg      []string
	Dir      string
	Wskprops *whisk.Wskprops
}

func NewWsk() *Wsk {
	return NewWskWithPath(GetBinPath())
}

func NewWskWithPath(path string) *Wsk {
	var dep Wsk
	dep.Path = cmd
	dep.Arg = []string{arg}
	dep.Dir = path
	pi := whisk.PropertiesImp{
		OsPackage: whisk.OSPackageImp{},
	}
	dep.Wskprops, _ = whisk.GetDefaultWskProp(pi)
	return &dep
}

func (wsk *Wsk) Exists() bool {
	_, err := os.Stat(wsk.Dir + "/" + wsk.Path)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (wsk *Wsk) RunCommand(s ...string) ([]byte, error) {
	cs := wsk.Arg
	cs = append(cs, s...)
	command := exec.Command(wsk.Dir+"/"+wsk.Path, cs...)
	command.Dir = wsk.Dir
	return command.CombinedOutput()
}

func (wsk *Wsk) ListNamespaces() ([]byte, error) {
	return wsk.RunCommand("namespace", "list", "--apihost", wsk.Wskprops.APIHost,
		"--auth", wsk.Wskprops.AuthKey)
}
