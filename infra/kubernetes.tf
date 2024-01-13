locals {
    cluster_name = "${var.environment}-cluster"
}

data "aws_eks_cluster" "cluster" {
    name = local.cluster_name
}

data "aws_eks_cluster_auth" "cluster" {
    name = local.cluster_name
}

provider "kubernetes" {
    host                   = data.aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
    token                  = data.aws_eks_cluster_auth.cluster.token
}

module "eks-kubeconfig" {
    source     = "hyperbadger/eks-kubeconfig/aws"
    version    = "2.0.0"

    depends_on   = [module.eks]
    cluster_name =  local.cluster_name
}

resource "local_file" "kubeconfig" {
    content  = module.eks-kubeconfig.kubeconfig
    filename = "kubeconfig_${local.cluster_name}"
}

module "eks" {
    source  = "terraform-aws-modules/eks/aws"
    version = "19.21.0"

    cluster_name    = "${local.cluster_name}"
    cluster_version = "1.27"
    subnet_ids      = module.vpc.private_subnets

    cluster_endpoint_public_access  = true

    cluster_addons = {
        coredns = {
            most_recent = true
        }
        kube-proxy = {
            most_recent = true
        }
        vpc-cni = {
            most_recent = true
        }
    }

    vpc_id = module.vpc.vpc_id

    eks_managed_node_groups = {
        first = {
            desired_capacity = 1
            max_capacity     = 3
            min_capacity     = 1

            instance_type = "t2.medium"
        }
    }

    tags = {
        Terraform = "true"
        Environment = var.environment
    }
}