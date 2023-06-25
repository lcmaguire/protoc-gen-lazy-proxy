
buf:
	go install . && buf generate

run-generated-proxy:
	go run lp/lazyproxy/main.go

run-grpc-server:
	go run manual/grpc/main.go

grpc-curl: # grpc curl for connect service
	grpcurl \
    -import-path ./proto/sample -proto sample.proto -plaintext \
	-H 'Authorization:Bearer tuki' \
    -d '{"name" : "tuki"}' \
    localhost:8080 sample.SampleService/Sample
	
curl: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{"name" : "tuki"}' \
    http://localhost:8080/sample.SampleService/Sample
