variable "environment" {
    description = "Name of the environment"
    type = string
    default = "staging"
}

variable "region" {
    description = "AWS Region"
    type = string
    default = "us-east-1"
}

variable "availability_zones" {
    description = "Availability zones in AWS Region"
    type = list(string)
    default = ["us-east-1a", "us-east-1b"]
}

variable "public_subnets_cidr" {
    description = "CIDR blocks for public subnets"
    type = list(string)
    default = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_subnets_cidr" {
    description = "CIDR blocks for private subnets"
    type = list(string)
    default = ["10.0.3.0/24", "10.0.4.0/24"]
}