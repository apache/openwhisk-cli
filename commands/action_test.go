// +build unit

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

package commands

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBallerinaBinaryFile(t *testing.T) {
	file := "file.balx"
	args := []string{"name", file}
	parms := ActionFlags{}

	f, error := os.Create(file)
	if error != nil {
		t.Fatalf("could not create file: %s", error.Error())
	}
	d := []byte("balx")
	f.Write(d)
	f.Close()

	exec, error := getExec(args, parms)
	os.Remove(file)

	if error != nil {
		t.Errorf("unexpected exec error: %s", error.Error())
	} else {
		assert.Equal(t, exec.Kind, "ballerina:default")
		assert.Equal(t, base64.StdEncoding.EncodeToString(d), *exec.Code)
	}
}
