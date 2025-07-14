package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	*Config
	Topics map[string]Storer
}

func NewServer(cfg *Config) (*Server, error) {

	return &Server{
		Config: cfg,
		Topics: make(map[string]Storer),
	}, nil
}

func (s *Server) Start() {

	http.ListenAndServe(s.ListenAddr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func (s *Server) createTopic(name string) bool {
	if _, ok := s.Topics[name]; !ok {
		s.Topics[name] = s.StoreProducerFunc()
		return true

	}
	return false
}
