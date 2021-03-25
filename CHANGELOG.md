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

## v1.2.0

- Update Whisk Deploy (openwhisk-wskdeploy) dependency to v1.2.0 (#510)
- Prep. for potential 1.2.0 release (#498)
- Update for travis migration (#492)
- Bump openwhisk-client-go dependency (#493)
- Remove trailing slash on apihost #481 (#485)
- Recognize .rs extension as a Rust action kind (#495)
- Remove last Godeps, update Gogradle for gomod and Ansible setup (#496)
- Update Gradle/Wrapper to latest version (#497)

## v1.1.0

- Upgrade all Go dependencies to latest (#490)
- Migrated to using go mod to manage dependencies (#489)
- Upgrade travis to go 1.15
- Support passing del annotation (#488)
- Add an overwrite flag to "package bind" (#474)
- Trigger parameter issue (#479)
- remove test for download of iOS SDK (#478)
- build binary test artifacts (#477)
- Update test file (#463)
- Fix regex for log stripping. (#462)
- Ensure that the pollSince is greater than Activation start time (#461)

## v1.0.0

- Allow log stripping to tolerate a missing stream identifier. (#444)
- Add --logs options on activation get to return stripped logs as a convenience. (#445)
- RestAssured fixes (#441)
- Remove namespace property from wskprops (#434)
- "wsk property get" can now return raw output for specific properties  (#430)
- Add dynamic column sizing to wsk activation list command (#427)

## v0.10.0

- Integrate wskdeploy via `project` subcommand
- Enhanced columnar output in `activation list`
- CLI support for OpenWhisk enhancements including:
- Support for specifying intra-container concurrency
- New supported action languages: Ballerina, .Net, and Go

## v0.9.0

- Initial release as an Apache Incubator project.
