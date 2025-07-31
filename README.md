# CS6650_TA

## Assignment 1 (Geoff's)
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
   
## Assignment 2b - Terraform