#gRPC
for ssl
choco install openssl

cd ssl 
powershell -ExecutionPolicy unrestricted .\ssl.ps1

go install update tools 