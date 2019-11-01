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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStripTimestamp(t *testing.T) {
	logs := map[string]string{
		"2018-05-02T19:33:32.829992819Z stdout: this is stdout stderr: this is still stdout":  "this is stdout stderr: this is still stdout",
		"2018-05-02T19:33:32.829992819Z stderr: this is stderr stdout: this is still stderr":  "this is stderr stdout: this is still stderr",
		"2018-05-02T19:33:32.829992819Z  stdout: this is stdout stderr: this is still stdout": "this is stdout stderr: this is still stdout",
		"2018-05-02T19:33:32.829992819Z  stderr: this is stderr stdout: this is still stderr": "this is stderr stdout: this is still stderr",
		"2018-05-02T19:33:32.89Z stdout: this is stdout":                                      "this is stdout",
		"2018-05-02T19:33:32.89Z this is a msg":                                               "this is a msg",
		"2018-05-02T19:33:32.89Z  this is a msg":                                              " this is a msg",
		"anything stdout: this is stdout":                                                     "anything stdout: this is stdout",
		"anything stderr: this is stderr":                                                     "anything stderr: this is stderr",
		"stdout: this is stdout":                                                              "stdout: this is stdout",
		"stderr: this is stderr":                                                              "stderr: this is stderr",
		"this is stdout":                                                                      "this is stdout",
		"this is stderr":                                                                      "this is stderr",
		"something":                                                                           "something",
		"":                                                                                    ""}
	assert := assert.New(t)

	for log, expected := range logs {
		assert.Equal(stripTimestamp(log), expected)
	}
}
