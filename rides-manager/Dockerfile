FROM golang:1.14.0-alpine

WORKDIR src/riders-manager

COPY . .

RUN go mod vendor
#RUN pwd
#RUN ls -als

RUN go build -o riders-manager cmd/main.go

#RUN ls -als

#EXPOSE 8080
ENTRYPOINT ["./riders-manager"]