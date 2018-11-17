#!/bin/bash

go test -v ./... $@ |tee ./test-results/test-result.txt
test_result=${PIPESTATUS[0]}
cat ./test-results/test-result.txt |go2xunit |tee ./test-results/test-result.xml
exit $test_result
