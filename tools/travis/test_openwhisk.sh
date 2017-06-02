#!/usr/bin/env bash

PIDS=()
ERRORS=0

function waitForAll() {
    for pid in ${PIDS[@]}; do
        wait $pid
        STATUS=$?
        echo "$pid finished with status $STATUS"
        if [ $STATUS -ne 0 ]
        then
            let ERRORS=ERRORS+1
        fi
    done
    PIDS=()
}

"$TRAVIS_BUILD_DIR/tools/travis/install_openwhisk.sh" &
PID=$!
PIDS+=($PID)

waitForAll

echo test openwhisk ERRORS = $ERRORS
exit $ERRORS
