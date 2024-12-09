gen:
	protoc \
	--proto_path=proto \
	--go_out=pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=pb \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto
clean:
	rm pb/*.go
runsv:
	go run server/main.go --part=8000
runcl:
	go run client/main.go --address 0.0.0.0:8000

run:clean gen runsv runcl
