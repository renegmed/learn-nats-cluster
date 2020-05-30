FROM golang:1.14.0-alpine

WORKDIR src/api-server

COPY . .

#RUN go mod init
RUN go mod vendor
#RUN pwd
#RUN ls -als 

RUN go build -o api-server cmd/main.go

#RUN ls -als

EXPOSE 8080
ENTRYPOINT ["./api-server"]