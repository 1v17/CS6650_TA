# CS6650_TA

## Assignment 1
### Local Test
Build Docker image:

```sh
docker build -t echo-service .
```

Run Docker image
```sh
docker run -p 8080:8080 echo-service
```

### Upload to AWS
Install AWS Tool ECR model on Windows:
1. Use administrator role to open PowerShell, run:
    ```sh
    Install-Module -Name AWS.Tools.ECR -Force
    ```
2. Failed to uplaod the actually Dockerfile

## Assignment 1b - AWS EC2
Note: 
1. On Windows, the ssh key must be specified as `C:\Users\UserName\.ssh\key_name.pem` in the ssh login command.
2. Change security groups inbound rules.

## Assignment 2a - Terraform

No problem, though additional files should be added to .gitignore.

## Assignment 2b - Dockerfile with GO

1. Remember to update aws credentials with each session info.
2. During docker build stage, the go.mod go version should be written as 1.23 (the major and minor version numbers).
3. The base image specified in Dockerfile must match the go.mod version of go.

## Assignment 2c - ECR/ECS Workflow

1. After setting up the aws credentials, use the following command to authenticate Docker to ECR:
    ```pwsh
    $password = aws ecr get-login-password --region us-west-2
    docker login -u AWS -p $password $ECR_BASE
    ```
    You will get a warning, but it will succeed. Otherwise there will be a `400 bad request` error.

2. The docker tag command on Powershell should be:
    ```pwsh
    docker tag gin_server:latest ${ECR_URL}:latest
    ```
