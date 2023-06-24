run-proxy:
	go run manual/connect-proxy/main.go

run-grpc-server:
	go run manual/grpc/main.go

grpc-curl: # grpc curl for connect service
	grpcurl \
    -import-path ./sample -proto sample.proto -plaintext \
    -d '{"name" : "tuki"}' \
    localhost:8081 sample.SampleService
	
curl: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{"name" : "tuki"}' \
    http://localhost:8080/sample.SampleService