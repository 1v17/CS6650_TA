# Simple Echo Service Test Script
param(
    [string]$BaseUrl = "http://localhost:8080/echo",
    [string]$Message = "Hello World!"
)

Write-Host "Testing Echo Service at: $BaseUrl" -ForegroundColor Green

# Test GET endpoint
Write-Host "`nTesting GET endpoint..." -ForegroundColor Yellow
$getUrl = "$BaseUrl" + "?message=" + [System.Web.HttpUtility]::UrlEncode($Message)
try {
    $getResult = Invoke-RestMethod -Uri $getUrl -Method Get
    Write-Host "GET Response: $getResult" -ForegroundColor Cyan
} catch {
    Write-Host "GET Failed: $($_.Exception.Message)" -ForegroundColor Red
}

# Test POST endpoint
Write-Host "`nTesting POST endpoint..." -ForegroundColor Yellow
try {
    $postResult = Invoke-RestMethod -Uri $BaseUrl -Method Post -Body $Message -ContentType "text/plain"
    Write-Host "POST Response: $postResult" -ForegroundColor Cyan
} catch {
    Write-Host "POST Failed: $($_.Exception.Message)" -ForegroundColor Red
}
