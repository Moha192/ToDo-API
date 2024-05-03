FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod download

WORKDIR /app/cmd
RUN go build -o todoApp .
EXPOSE 8080

CMD ["./todoApp"]