package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test http to ELB
func testHTTPToELB(t *testing.T, terraformOptions *terraform.Options) {
	t.Parallel()

	// Run `terraform output` to get the value of an output variable
	instanceURL := terraform.Output(t, terraformOptions, "address")

	// It can take a minute or so for the Instance to boot up, so retry a few times
	maxRetries := 30
	timeBetweenRetries := 5 * time.Second

	instanceText := "TBD"

	// Verify that we get back a 200 OK with the expected instanceText
	http_helper.HttpGetWithRetry(t, instanceURL, 200, instanceText, maxRetries, timeBetweenRetries)
}
