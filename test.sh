#!/bin/bash

go test -v ./... |tee test-result.txt
test_result=$?
cat test-result.txt |go2xunit |tee test-result.xml
exit $test_result
