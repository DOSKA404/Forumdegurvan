FROM golang:alpine3.18
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o bin/api main.go
CMD [ "./bin/api" ]