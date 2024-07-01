FROM golang:1.18.0-alpine3.15

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o mnemosyne .

EXPOSE 8080

CMD [ "./mnemosyne" ]
