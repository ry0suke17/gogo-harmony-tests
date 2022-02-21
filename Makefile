dev/dep/install:
	git clone https://github.com/gogo/protobuf.git -b v1.3.2 ./third_party/github.com/gogo/protobuf
	rm -rf ./third_party/github.com/gogo/protobuf/.git
	git clone https://github.com/grpc-ecosystem/grpc-gateway.git -b v1.16.0 ./third_party/github.com/grpc-ecosystem/grpc-gateway
	rm -rf ./third_party/github.com/grpc-ecosystem/grpc-gateway/.git
	cd ./third_party/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && go install ./... && cd -
	cd ./third_party/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && go install ./... && cd -
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip
	unzip protoc-3.19.4-osx-x86_64.zip -d third_party/protoc-3.19.4-osx-x86_64/
	rm -rf protoc-3.19.4-osx-x86_64.zip
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.0
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@v1.3.2

dev/protoc/gen/go:
	./third_party/protoc-3.19.4-osx-x86_64/bin/protoc \
		--proto_path=./third_party \
		-I=./proto \
		--go_out=./proto/go \
		./proto/test.proto

dev/protoc/gen/gogofaster:
	./third_party/protoc-3.19.4-osx-x86_64/bin/protoc \
		--proto_path=./third_party \
		-I=./proto \
		--gogofaster_out=./proto/gogofaster \
		./proto/test.proto
