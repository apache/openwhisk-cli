<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

# Changelog

## v1.0.0
  * Allow log stripping to tolerate a missing stream identifier. (#444)
  * Add --logs options on activation get to return stripped logs as a convenience. (#445)
  * RestAssured fixes (#441)
  * Remove namespace property from wskprops (#434)
  * "wsk property get" can now return raw output for specific properties  (#430)
  * Add dynamic column sizing to wsk activation list command (#427)

## v0.10.0

* Integrate wskdeploy via `project` subcommand
* Enhanced columnar output in `activation list`
* CLI support for OpenWhisk enhancements including:
  * Support for specifying intra-container concurrency
  * New supported action languages: Ballerina, .Net, and Go

## v0.9.0

* Initial release as an Apache Incubator project.
