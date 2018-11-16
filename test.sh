#!/bin/bash

go test -v ./... |tee test-result.txt
cat test-result.txt |go2xunit |tee test-result.xml