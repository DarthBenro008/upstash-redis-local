package main

import "upstash-redis-local/internal"

func main() {

	server := internal.Server{Address: ":8000"}
	defer server.Serve()
}
