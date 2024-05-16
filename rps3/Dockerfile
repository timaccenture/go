FROM golang:1.22

WORKDIR /rps3

COPY go.mod ./
RUN go mod download && go mod verify

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /rps3

EXPOSE 8080


CMD ["/rps3"]

