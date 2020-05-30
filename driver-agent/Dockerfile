FROM golang:1.14.0-alpine

WORKDIR src/driver-agent

COPY . .

RUN go mod vendor

#RUN pwd
#RUN ls -als

RUN go build -o driver-agent cmd/main.go

#RUN ls -als 
 
ENTRYPOINT ["./driver-agent"]