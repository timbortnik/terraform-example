Terraform Provider + Terratest example
==================

This code is based on Terraform provider and Terratest examples
- Website: https://www.terraform.io
- Website: https://blog.gruntwork.io/open-sourcing-terratest-a-swiss-army-knife-for-testing-infrastructure-code-5d883336fcd5

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.11 (to run tests)

Notes
-----

Two-tier terrform template example was modified to
- accept SSH private key path as a parameter
- output backend ip address
- tag backend with meaningful Name
- aws elb name hardcode changed to avoid conflicts in parallel execution

CI
--

CircleCI runs tests on branches 
[![CircleCI](https://circleci.com/gh/timbortnik/terraform-example.svg?style=svg)](https://circleci.com/gh/timbortnik/terraform-example)

Tests
-----

This repo should be clonet to go path

Dependencies can be installed by ./test-setup.sh

Tests can be executed by ./test.sh

test-results folder contains test-result.xml, junit-compatible

```
<?xml version="1.0" encoding="UTF-8"?>

  <testsuite name="terraform-example/test" tests="8" errors="0" failures="0" skip="0">
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier" time="18.14">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/HTTP_to_ELB" time="0.50">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/SCP_to_public_host" time="4.83">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/SSH_to_public_host" time="2.33">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/Check_Nginx_access_log_for_external_ip" time="3.03">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/Check_Nginx_access_log_for_URL_path" time="2.62">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/Check_Nginx_listens_on_localhost" time="2.32">

    </testcase>
    <testcase classname="terraform-example/test" name="TestTerraformTwoTier/Check_Nginx_worker_processes_auto_setting" time="2.50">

    </testcase>
  </testsuite>
```