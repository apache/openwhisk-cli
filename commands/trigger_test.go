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
	"errors"
	"github.com/apache/openwhisk-client-go/whisk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var Triggers = make(map[string]*whisk.Trigger)

type MockedTriggerService struct {
}

func (t MockedTriggerService) List(options *whisk.TriggerListOptions) ([]whisk.Trigger, *http.Response, error) {
	return []whisk.Trigger{}, &http.Response{}, nil
}
func (t MockedTriggerService) Insert(trigger *whisk.Trigger, overwrite bool) (*whisk.Trigger, *http.Response, error) {
	Triggers[trigger.Name] = trigger
	return trigger, &http.Response{}, nil
}
func (t MockedTriggerService) Get(triggerName string) (*whisk.Trigger, *http.Response, error) {
	var trigger *whisk.Trigger
	var ok bool
	var err error = nil
	var httpResponse http.Response
	if trigger, ok = Triggers[triggerName]; !ok {
		err = errors.New("Unable to get trigger")
		httpResponse = http.Response{StatusCode: 404}
	}
	return trigger, &httpResponse, err
}
func (t MockedTriggerService) Delete(triggerName string) (*whisk.Trigger, *http.Response, error) {
	return &whisk.Trigger{}, &http.Response{}, nil
}
func (t MockedTriggerService) Fire(triggerName string, payload interface{}) (*whisk.Trigger, *http.Response, error) {
	return &whisk.Trigger{}, &http.Response{}, nil
}

var _ = Describe("Trigger Command", func() {
	t := Trigger{}
	name := "awesomeTrigger"
	client := whisk.Client{Triggers: &MockedTriggerService{}, Config: &whisk.Config{}}
	args := []string{name}

	BeforeEach(func() {
		Triggers = make(map[string]*whisk.Trigger)
	})

	It("should update an existing trigger", func() {
		Triggers[name] = &whisk.Trigger{}
		Expect(len(Triggers)).To(Equal(1))
		err := t.Update(&client, args)
		Expect(err).To(BeNil())
		Expect(len(Triggers)).To(Equal(1))
	})

	It("should create a trigger on update when it does not exist yet", func() {
		Expect(len(Triggers)).To(Equal(0))
		err := t.Update(&client, args)
		Expect(err).To(BeNil())
		Expect(len(Triggers)).To(Equal(1))
	})
})
