package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func configureTerraformOptions(t *testing.T, exampleFolder string) (*terraform.Options, *aws.Ec2Keypair) {
	// A unique ID we can use to namespace resources so we don't clash with anything already in the AWS account or
	// tests running in parallel
	uniqueID := random.UniqueId()

	// Give this EC2 Instance and other resources in the Terraform code a name with a unique ID so it doesn't clash
	// with anything else in the AWS account.
	instanceName := fmt.Sprintf("terratest-ssh-example-%s", uniqueID)

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	// awsRegion := aws.GetRandomRegion(t, []string{"eu-west-1", "us-east-1", "us-west-1", "us-west-2"}, nil)
	awsRegion := aws.GetRandomRegion(t, []string{"us-west-1"}, nil)

	// Create an EC2 KeyPair that we can use for SSH access
	keyPairName := fmt.Sprintf("terratest-ssh-example-%s", uniqueID)
	sshKeyPair, err := ssh.GenerateRSAKeyPairE(t, 2048)
	require.Nil(t, err, "Failed to generate ssh keypair")

	keyPair := &aws.Ec2Keypair{Name: keyPairName, Region: awsRegion, KeyPair: sshKeyPair}

	keyPathPublic := exampleFolder + "/.test-data/key.pub"
	ioErrPublic := ioutil.WriteFile(keyPathPublic, []byte(keyPair.PublicKey), 0644)
	require.Nil(t, ioErrPublic, "Failed to write public key file")

	keyPathPrivate := exampleFolder + "/.test-data/key"
	ioErrPrivate := ioutil.WriteFile(keyPathPrivate, []byte(keyPair.PrivateKey), 0600)
	require.Nil(t, ioErrPrivate, "Failed to write private key file")

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_region":       awsRegion,
			"instance_name":    instanceName,
			"key_name":         keyPairName,
			"public_key_path":  keyPathPublic,
			"private_key_path": keyPathPrivate,
		},
	}

	return terraformOptions, keyPair
}
