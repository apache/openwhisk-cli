#!/usr/bin/env bash
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

set -e

ROOT_DIR="$(git rev-parse --show-toplevel)"

set +e
FILE_EXT=".go"
STAGED_FILES=$(git diff --name-only --no-color --diff-filter=d --exit-code -- "${ROOT_DIR}/*$FILE_EXT")
STAGED_FILES_DETECTED=$?
set -e

if [ "${STAGED_FILES_DETECTED}" -eq 1 ]; then
    # Re-format and re-add all staged files
    for FILE in ${STAGED_FILES}
    do
      gofmt -s -w "${ROOT_DIR}/${FILE}"
      git add -- "${ROOT_DIR}/${FILE}"
    done
fi

exit 0
