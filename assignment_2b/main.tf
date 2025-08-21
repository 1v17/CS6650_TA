# You probably want to keep your ip address a secret as well
variable "ssh_cidr" {
  type        = string
  description = "Your home IP in CIDR notation"
}

# name of the existing AWS key pair
variable "ssh_key_name" {
  type        = string
  description = "Name of your existing AWS key pair"
}

# GitHub repository URL
variable "github_repo_url" {
  type        = string
  description = "GitHub repository URL to clone"
}

# The provider of your cloud service, in this case it is AWS. 
provider "aws" {
  region     = "us-west-2" # Which region you are working on
}

# Your ec2 instance
resource "aws_instance" "demo-instance" {
  ami                    = data.aws_ami.al2023.id
  instance_type          = "t2.micro"
  iam_instance_profile   = "LabInstanceProfile"
  vpc_security_group_ids = [aws_security_group.ssh.id]
  key_name               = var.ssh_key_name

  user_data = <<-EOF
    #!/bin/bash
    yum update -y
    
    # Install Git
    yum install -y git
    
    # Install Docker
    yum install -y docker
    
    # Start Docker service
    systemctl start docker
    systemctl enable docker
    
    # Add ec2-user to docker group
    usermod -a -G docker ec2-user
    
    # Wait for Docker to be ready
    sleep 10
    
    # Clone your repository
    cd /home/ec2-user
    git clone ${var.github_repo_url} app
    cd app/assignment_2b
    
    # Build and run Docker container
    docker build --tag gin_server .
    docker run gin_server
    
    # Change ownership to ec2-user
    chown -R ec2-user:ec2-user /home/ec2-user/app
  EOF

  tags = {
    Name = "terraform-go"
  }
}

# Your security that grants ssh access from 
# your ip address to your ec2 instance
resource "aws_security_group" "ssh" {
  name        = "allow_ssh_from_me"
  description = "SSH from a single IP"
  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.ssh_cidr]
  }
  ingress {
    description = "HTTP"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = [var.ssh_cidr]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# latest Amazon Linux 2023 AMI
data "aws_ami" "al2023" {
  most_recent = true
  owners      = ["amazon"]
  filter {
    name   = "name"
    values = ["al2023-ami-*-x86_64-ebs"]
  }
}

output "ec2_public_dns" {
  value = aws_instance.demo-instance.public_dns
}
