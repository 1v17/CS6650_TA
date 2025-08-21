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

1. During docker build stage, the go.mod go version should be written as 1.23 (the major and minor version numbers).
2. The base image specified in Dockerfile must match the go.mod version of go.