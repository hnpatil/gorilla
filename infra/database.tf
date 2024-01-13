locals {
    postgres_version = "13.10"
    db_user = "postgres"
}

resource "aws_security_group" "db" {
    description = "Security group to allow inbound/outbound from db"
    vpc_id      = module.vpc.vpc_id
    depends_on  = [module.vpc]

    ingress {
        from_port   = "5432"
        to_port     = "5432"
        protocol    = "tcp"
        cidr_blocks = [module.vpc.vpc_cidr_block]
        self        = true
    }

    tags = {
        Name = "${var.environment}-db-sg"
    }   
}

resource "random_password" "db" {
    length  = 16
    special = false
}


resource "aws_db_subnet_group" "db" {
  subnet_ids = module.vpc.private_subnets

  tags = {
    Name = "${var.environment}-db-subnet-group"
  }
}

module "db" {
    source = "terraform-aws-modules/rds/aws"

    identifier = "${var.environment}db"

    engine            = "postgres"
    engine_version    = local.postgres_version
    instance_class    = "db.t3.micro"
    allocated_storage = 10

    db_name  = "${var.environment}db"
    username = local.db_user
    password = random_password.db.result
    port     = "5432"

    iam_database_authentication_enabled = true

    vpc_security_group_ids = [aws_security_group.db.id]

    maintenance_window = "Mon:00:00-Mon:03:00"
    backup_window      = "03:00-06:00"

    db_subnet_group_name = aws_db_subnet_group.db.name

    family = "postgres${element(split(".", local.postgres_version), 0)}"

    tags = {
        Name = "${var.environment}-db-instance"
    }   
}

resource "aws_secretsmanager_secret" "db" {
    name = "${var.environment}-db-url"
    
    tags = {
        Name = "${var.environment}-db-url"
    }
}

resource "aws_secretsmanager_secret_version" "db" {
    secret_id =  aws_secretsmanager_secret.db.id
    secret_string = "postgresql://${local.db_user}:${random_password.db.result}@${module.db.db_instance_address}:5432/${var.environment}db"
}