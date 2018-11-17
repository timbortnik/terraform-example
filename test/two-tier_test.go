package test

import (
	"io/ioutil"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/require"
)

// An example of how to test the Terraform module in examples/terraform-ssh-example using Terratest. The test also
// shows an example of how to break a test down into "stages" so you can skip stages by setting environment variables
// (e.g., skip stage "teardown" by setting the environment variable "SKIP_teardown=true"), which speeds up iteration
// when running this test over and over again locally.
func TestTerraformTwoTier(t *testing.T) {

	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../two-tier", ".")

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)
		terraform.Destroy(t, terraformOptions)

		keyPair := test_structure.LoadEc2KeyPair(t, exampleFolder)
		aws.DeleteEC2KeyPair(t, keyPair)
	})

	// Deploy the example
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions, keyPair, keyPathPublic, keyPathPrivate := configureTerraformOptions(t, exampleFolder)

		// Save the options and key pair so later test stages can use them
		test_structure.SaveTerraformOptions(t, exampleFolder, terraformOptions)
		test_structure.SaveEc2KeyPair(t, exampleFolder, keyPair)

		ioErrPublic := ioutil.WriteFile(keyPathPublic, []byte(keyPair.PublicKey), 0644)
		require.Nil(t, ioErrPublic, "Failed to write public key file")

		ioErrPrivate := ioutil.WriteFile(keyPathPrivate, []byte(keyPair.PrivateKey), 0600)
		require.Nil(t, ioErrPrivate, "Failed to write private key file")

		// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Make sure we can SSH to the public Instance directly from the public Internet and the private Instance by using
	// the public Instance as a jump host
	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, exampleFolder)
		keyPair := test_structure.LoadEc2KeyPair(t, exampleFolder)

		t.Run("SSH to public host and adjust nginx log settings", func(t *testing.T) {
			testSSHToPublicHost(t, terraformOptions, keyPair, "echo 123")
		})

		t.Run("SCP to public host", func(t *testing.T) {
			testSCPToPublicHost(t, terraformOptions, keyPair)
		})

		t.Run("HTTP to ELB", func(t *testing.T) {
			testHTTPToELB(t, terraformOptions)
		})
	})

}
