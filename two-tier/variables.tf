variable "public_key_path" {
  description = <<DESCRIPTION
Path to the SSH public key to be used for authentication.
Example: ~/.ssh/terraform.pub
DESCRIPTION
}

variable "private_key_path" {
  description = <<DESCRIPTION
Path to the SSH private key to be used for authentication.
Example: ~/.ssh/id_rsa
DESCRIPTION
}

variable "key_name" {
  description = "Desired name of AWS key pair"
}

variable "instance_name" {
  description = "Desired name of EC2 instance"
  default     = "Two-tier test backend"
}

variable "aws_region" {
  description = "AWS region to launch servers."
  default     = "us-west-2"
}

# Ubuntu Precise 12.04 LTS (x64)
variable "aws_amis" {
  default = {
    eu-west-1 = "ami-674cbc1e"
    us-east-1 = "ami-1d4e7a66"
    us-west-1 = "ami-969ab1f6"
    us-west-2 = "ami-8803e0f0"
  }
}
