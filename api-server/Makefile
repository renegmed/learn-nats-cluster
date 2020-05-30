 init-project:
	go mod init github.com/renegmed/riders-client
.PHONY: init-project


run-server:
	go run cmd/main.go --nats=nats://172.20.0.3:4222,nats://172.20.0.4:4223,nats://172.20.02:4224 --listen localhost:8090 
.PHONY: run-server

server-client:
	curl localhost:8090/
.PHONY: server-client

rider-request:
	curl -X POST localhost:8090/rides
.PHONY: rider-request