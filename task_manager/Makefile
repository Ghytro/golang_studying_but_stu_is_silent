all: compile_proto
	go mod vendor
	go build

compile_proto:
	protoc --go_out="." pb/task.proto

clean: all
	del /f pb\task.pb.go
	rmdir /S /Q vendor
