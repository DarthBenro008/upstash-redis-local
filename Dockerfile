FROM golang:1.21

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8000
ENTRYPOINT ["./bin/upstash-redis-local"]
