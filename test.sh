#!/bin/bash

go test -v ./... > ./test-results/test-result.txt
test_result=$?
cat ./test-results/test-result.txt |go2xunit |tee ./test-results/test-result.xml
exit $test_result
