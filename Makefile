up:
	docker-compose up --build -d 
.PHONY: up

logs-nats-1:
	docker logs nats_1 -f 

logs-nats-2:
	docker logs nats_2 -f 

logs-nats-3:
	docker logs nats_3 -f 


# To get connect urls:
connect-urls:
	curl http://127.0.0.1:8222/varz | more
# get the three connect urls and pass to api-server main.go 
# e.g   > go run cmd/main.go --nats "nats://172.20.0.3:4222,nats://172.20.0.4:4223,nats://172.20.02:4224"


logs-api-server:
	docker logs api-server -f 

logs-rides-manager: 
	docker logs rides-manager -f 

logs-driver-agent-regular:
	docker logs driver-agent-regular -f 

logs-driver-agent-suv:
	docker logs driver-agent-suv -f 

.PHONY: logs-nats-1 logs-nats-2 logs-nats-3 logs-api-server logs-rides-manager logs-driver-agent-regular logs-driver-agent-suv
 


request-suv:
	curl -X POST http://localhost:8080/rides  -d '{"type":"suv", "location": {"lat": 12342243.344, "lng": 345534555.458}}'
.PHONY: request-suv

request-regular:
	curl -X POST http://localhost:8080/rides  -d '{"type":"regular", "location": {"lat": 12342243.344, "lng": 345534555.458}}'
.PHONY: request-regular

request-bus:
	curl -X POST http://localhost:8080/rides  -d '{"type":"bus", "location": {"lat": 12342243.344, "lng": 345534555.458}}'
.PHONY: request-bus