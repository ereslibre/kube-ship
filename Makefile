build:
	go build -o kube-ship cmd/kube-ship/main.go

install:
	go install ./...
