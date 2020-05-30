 init-project:
	go mod init github.com/renegmed/driveragent
.PHONY: init-project


run-server:
	go run cmd/main.go --nats "nats://localhost:4221,nats://localhost:4222,nats://localhost:4223" --listen localhost:8090 
.PHONY: run-server

server-client:
	curl localhost:8090/
.PHONY: server-client

rider-request:
	curl -X POST localhost:8080/rides
.PHONY: rider-request