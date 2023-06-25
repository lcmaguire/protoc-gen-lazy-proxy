
buf:
	go install . && buf generate

protoc-sample:
	go install . && \
	protoc -I=sample \
    --lazy-proxy_out=lp \
	--lazy-proxy_opt=paths=source_relative \
	sample/*.proto 

run-proxy:
	go run manual/connect-proxy/main.go

run-generated-proxy:
	go run lp/lazyproxy/main.go

run-grpc-server:
	go run manual/grpc/main.go

grpc-curl: # grpc curl for connect service
	grpcurl \
    -import-path ./sample -proto sample.proto -plaintext \
	-H 'Authorization:Bearer tuki' \
    -d '{"name" : "tuki"}' \
    localhost:8080 tutorial.SampleService/Sample
	
curl: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{"name" : "tuki"}' \
    http://localhost:8080/tutorial.SampleService/Sample
