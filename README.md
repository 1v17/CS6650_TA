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

## Assignment 3b & 3c - Terraform for ECR/ECS

See the other [repo](https://github.com/1v17/CS6650_2b_demo).

## Assignmnet 4a

### Atomicity

Data race will more likely to appear when run with `-race` flag. This flag will make race detector add extra checks to every memory access to detect data races.

### Collections

1. Results of collection experiment:
   ```
   Mutex map test:
   len(m) = 50000, time = 8.8637ms
   RWMutex map test:
   len(m) = 50000, time = 12.5759ms
   sync.Map test:
   len(m) = 50000, time = 28.8603ms
   ```
2. In write-heavy sccenario, the `sync.Map` might be the most inefficient one. This struct is optimized for write-once, read-many operations and disjoint key access for multiple goroutines.

### File Access

1. Results of two file writing functions:
   ```
   Unbuffered write: 713.3351ms
   Buffered write:   10.9805m
   ```
2. Buffered file writing is much faster because it reduces the number of system calls made to the operating system, while every `f.Write` sends data directly to the OS, causing a system call for every line (100,000 times). System calls are slow and expensive.

### Context Switching

1. Result of context switching for goroutines:
   ```
   GOMAXPROCS=1: total=381.0931ms, avg switch=190ns
   GOMAXPROCS=8: total=506.1486ms, avg switch=253ns
   ```
2. When `GOMAXPROCS=1`, Go runs all goroutines on a single OS thread. Context switching between goroutines is handled entirely in user space by the Go scheduler, which is very fast because it doesn't involve the operating system. When `GOMAXPROCS=8` (or any value >1), Go can run goroutines on multiple OS threads. It is slower because the OS must manage thread scheduling and there may be additional overhead for synchronizing memory between threads (cache coherence, etc.).

## Assignment 4b

1. Post requests consistently take more time then Get requests. This pattern is expected since I use `sync.Map` to store the product data.
2. After switching from `HttpUser` to `FastHttpUser`, the throughput became wavey. It might hit certain threshold in my laptop?

## Assignment 5a

Note:

1. The week number should be changed to week 6, otherwise the number will become more and more confusing.
2. The students will probably use the product API implemented in assignment 4, which utilizes thread-safe map.
3. Remind them to get rid of `wait_time` (if added in last assignment) first, otherwise request per second can be fairly low.
4. The terraform template in [2b_demo](https://github.com/RuidiH/CS6650_2b_demo/tree/master/terraform) (actually in assignment 3b, the naming is really confusing) doesn't have a `.tfvars` file, perhaps we should provide them with an `example.tfvars` file?
5. Part 3 has some formatting problem, there are three 1s.
6. None of the resources hit 100% in part 3, which means my laptop was the bottleneck. For students, the bottleneck might be something else though, network speed for example.

### Part 3

#### Under `wait_time = between(1, 3)`

- Peak CPU utilization: 6.07%
- Peak memory utilization: 2.99%
- Average response time: 17ms
- P95 response time: 42ms
- Requests per second: 97.8
- Error rate: 0%

#### With no `wait_time`

- Peak CPU utilization: 55.2%
- Peak memory utilization: 11.98%
- Average response time: 40ms
- P95 response time: 110ms
- Requests per second: 2015.9
- Error rate: 0%

#### On loaned laptop from campus

- Peak CPU utilization: []%
- Peak memory utilization: []%
- Average response time: 68ms
- P95 response time: 110ms
- Requests per second: 2769
- Error rate: 0%

### Part 4

#### On my own laptop

- Peak CPU utilization: 26.58%
- Peak memory utilization: 2.27%
- Average response time: 32ms
- P95 response time:62ms
- Requests per second: 1855.8
- Error rate: 0%

#### On loaned laptop from campus

- Peak CPU utilization: 60.44%
- Peak memory utilization: 10.47%
- Average response time: 31ms
- P95 response time:57ms
- Requests per second: 5628.5
- Error rate: 0%

### Part 5

| Configuration | CPU  | Memory  | Relative Cost | Avg Response | RPS | Performance per $ |
| ------------- | ---- | ------- | ------------- | ------------ | --- | ----------------- |
| Baseline      | 256  | 512 MB  | $ (1x)        |              |     |                   |
| Scaled        | 1024 | 2048 MB | $ (4x)        |              |     |                   |
