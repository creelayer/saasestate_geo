FROM golang:1.14-alpine3.12

WORKDIR /go/src/app
COPY . .

#RUN ["chmod", "+x", "./wait.sh"]

#RUN go get -d -v ./...
#RUN go install -v ./...

RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]