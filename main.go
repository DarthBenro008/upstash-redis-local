package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"upstash-redis-local/internal"
)

func main() {

	server := internal.Server{Address: ":8000", APIToken: "gg", RedisConn: connectToRedis()}
	defer server.Serve()
}

func connectToRedis() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return conn
}
