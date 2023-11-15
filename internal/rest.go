package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
	"sync"
)

// Server Configurations
type Server struct {
	Address     string
	APIToken    string
	RedisConn   redis.Conn
	credentials map[string]credentials
	mutex       sync.Mutex
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
	if !ctx.IsGet() && !ctx.IsPost() && !ctx.IsHead() && !ctx.IsPut() {
		s.respond(ctx, nil, fasthttp.StatusMethodNotAllowed)
		return
	}
	_, err := s.authenticate(ctx)
	if err != nil {
		s.respond(ctx, errorResult{Error: "Unauthorised"}, fasthttp.StatusUnauthorized)
		return
	}

	switch endpoint := string(ctx.Path()); strings.TrimPrefix(endpoint, "/") {
	case "":
		s.handleSingleExecute(ctx)
		return
	case "/pipeline":
		s.handlePipelineExecute(ctx)
		return
	default:
		segments := strings.Split(endpoint, "/")[1:]

		if len(ctx.PostBody()) > 0 {
			segments = append(segments, string(ctx.PostBody()))
		}

		if ctx.QueryArgs().String() != "" {
			qparts := strings.Split(ctx.QueryArgs().String(), "&")
			for _, qpart := range qparts {
				kv := strings.SplitN(qpart, "=", 2)
				if kv[0] == "_token" {
					continue
				}
				segments = append(segments, kv...)
			}
		}

		args := make([]interface{}, len(segments)-1)
		for i, data := range segments[1:] {
			args[i] = data
		}
		res, code := s.executeCommand(segments[0], args...)
		s.respond(ctx, res, code)
		return

	}
}

func (s *Server) handleSingleExecute(ctx *fasthttp.RequestCtx) {

	var args []interface{}

	if err := json.Unmarshal(ctx.PostBody(), &args); err != nil {
		s.respond(ctx, errorResult{Error: "ERR failed to parse command"}, fasthttp.StatusBadRequest)
		return
	}
	if len(args) == 0 {
		s.respond(ctx, errorResult{Error: "ERR empty command"}, fasthttp.StatusBadRequest)
		return
	}
	result, code := s.executeCommand(fmt.Sprint(args[0]), args[1:]...)
	s.respond(ctx, result, code)
	return
}

func (s *Server) handlePipelineExecute(ctx *fasthttp.RequestCtx) {
	var pipelineRequests [][]interface{}

	if err := json.Unmarshal(ctx.PostBody(), &pipelineRequests); err != nil {
		s.respond(ctx, errorResult{Error: "ERR failed to parse pipeline request"}, fasthttp.StatusBadRequest)
		return
	}
	if len(pipelineRequests) == 0 {
		s.respond(ctx, errorResult{Error: "ERR empty pipeline request"}, fasthttp.StatusBadRequest)
		return
	}

	var results []interface{}
	for _, request := range pipelineRequests {
		if len(request) == 0 {
			results = append(results, errorResult{Error: "ERR empty pipeline command"})
			continue
		}
		result, _ := s.executeCommand(fmt.Sprint(request[0]), request[1:]...)
		results = append(results, result)
	}
	s.respond(ctx, results, fasthttp.StatusOK)
	return
}

func (s *Server) executeCommand(commandName string, args ...interface{}) (interface{}, int) {
	// TODO: ACL
	res, err := s.RedisConn.Do(commandName, args...)
	if err != nil {
		return errorResult{Error: err.Error()}, fasthttp.StatusBadRequest
	}
	return successResult{Result: res}, fasthttp.StatusOK
}

func (s *Server) parseToken(ctx *fasthttp.RequestCtx) string {
	token := string(ctx.Request.Header.Peek("Authorization"))
	if token != "" {
		return strings.TrimPrefix(token, "Bearer ")
	}
	return ""
}

func (s *Server) authenticate(ctx *fasthttp.RequestCtx) (*credentials, error) {
	token := s.parseToken(ctx)
	if token == "" {
		return nil, errors.New("invalid token")
	}
	if token == s.APIToken {
		return &credentials{}, nil
	}
	s.mutex.Lock()
	credential, found := s.credentials[token]
	s.mutex.Unlock()
	if !found {
		return nil, errors.New("invalid token")
	}
	return &credential, nil
}

func (s *Server) respond(ctx *fasthttp.RequestCtx, data interface{}, status int) {
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(status)
	if data != nil {
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
}
