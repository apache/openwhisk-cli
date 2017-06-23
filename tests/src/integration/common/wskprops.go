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
    "github.com/spf13/viper"
    "io/ioutil"
    "os"
)

type Wskprops struct {
    APIHost string
    APIVersion string
    AuthKey string
    ControllerHost string
    ControllerPort string
}

func GetWskprops() *Wskprops {
    var dep Wskprops
    dep.APIHost = ""
    dep.AuthKey = ""
    dep.APIVersion = "v1"

    viper.SetConfigName("whisk")
    viper.AddConfigPath(os.Getenv("OPENWHISK_HOME"))

    err := viper.ReadInConfig()
    if err == nil {
        authPath := viper.GetString("testing.auth")

        b, err := ioutil.ReadFile(authPath)
        if err == nil {
            dep.AuthKey = string(b)
        }
        dep.APIHost = viper.GetString("controller.hosts")
        dep.ControllerHost = viper.GetString("controller.hosts")
        dep.ControllerPort = viper.GetString("controller.host.basePort")
    }
    return &dep
}
