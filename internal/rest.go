package internal

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

// Server Configurations
type Server struct {
	Address     string
	APIToken    string
	Credentials map[string]credentials
}

type credentials struct {
	Username string
	Password string
}

type errorResult struct {
	Error string `json:"error"`
}

type successResult struct {
	Result interface{} `json:"result"`
}

// Serve exposed function to start the server
func (s *Server) Serve() {

	if err := fasthttp.ListenAndServe(s.Address, s.requestHandler); err != nil {
		log.Fatalf("Error in serving: %v", err)
	}
}

// Handle requests for each query
func (s *Server) requestHandler(ctx *fasthttp.RequestCtx) {

	s.respond(ctx, successResult{Result: "hello-world"}, fasthttp.StatusOK)
}

func (s *Server) respond(ctx *fasthttp.RequestCtx, data interface{}, status int) {
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(status)
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("something went wrong %v\n", err)
		s.respond(ctx, errorResult{Error: fmt.Sprintf("something went wrong: %v", err)}, fasthttp.StatusInternalServerError)
	}
	_, err = ctx.Write(b)
	if err != nil {
		log.Printf("something went wrong %v\n", err)
	}
}
