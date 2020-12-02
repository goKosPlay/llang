build:
	go build  -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -o ./bin/tran_llang ./cmd/engine/main.go
run:
	./bin/tran_llang