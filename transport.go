package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type Consumer interface {
	Start() error
}

type Producer interface {
	Start() error
}

type HTTPproducer struct {
	listenAddr string
	server     *Server
	producerch chan<- Message
}

func NewHTTPproducer(listenAddr string, producerch chan Message) *HTTPproducer {
	return &HTTPproducer{
		listenAddr: listenAddr,
		producerch: producerch,
	}
}

func (p *HTTPproducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)
	// commit
	if r.Method == "GET" {
	}

	if r.Method == "POST" {
		if len(parts) != 2 {
			fmt.Println("invalid action")
			return
		}
		p.producerch <- Message{
			Data:  []byte("we dont know yet"),
			Topic: parts[1],
		}
	}

	fmt.Println(parts)
}

func (p *HTTPproducer) Start() error {

	slog.Info("http transprot started at ", " port", p.listenAddr)
	return http.ListenAndServe(p.listenAddr, p)
}
