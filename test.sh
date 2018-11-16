#!/bin/bash

go test -v ./... |tee test-result.txt
result=$?
cat test-result.txt |go2xunit |tee test-result.xml
exit $result