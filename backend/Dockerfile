FROM golang:1.25.2

EXPOSE 3000

WORKDIR /app
COPY . /app

RUN go build .

CMD ["./randochess"]