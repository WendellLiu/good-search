gen-pb:
	protoc -I pkg/pb/ pkg/pb/*.proto --go_out=plugins=grpc:pkg/pb
