FROM golang:alpine3.18
WORKDIR /app
COPY . .
RUN go build -o bin/front main.go
CMD [ "./bin/front" ]