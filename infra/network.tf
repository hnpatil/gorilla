module "vpc" {
    source  = "terraform-aws-modules/vpc/aws"
    version = "5.5.0"

    name                 = "${var.environment}-vpc"
    cidr                 = "10.0.0.0/16"
    azs                  = var.availability_zones
    private_subnets      = var.private_subnets_cidr
    public_subnets       = var.public_subnets_cidr
    enable_nat_gateway   = true
    single_nat_gateway   = true
    enable_dns_hostnames = true

    public_subnet_tags = {
        "kubernetes.io/cluster/${local.cluster_name}" = "shared"
        "kubernetes.io/role/elb"                      = "1"
    }

    private_subnet_tags = {
        "kubernetes.io/cluster/${local.cluster_name}" = "shared"
        "kubernetes.io/role/internal-elb"             = "1"
    }

    tags = {
        Terraform = "true"
        Environment = var.environment
    }
}