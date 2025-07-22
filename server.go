package main

import (
	"fmt"
	"log/slog"
)

type Message struct {
	Topic string
	Data  []byte
}

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	*Config
	Topics     map[string]Storer
	Consumers  []Consumer
	producers  []Producer
	Producerch chan Message
	quitch     chan struct{}
}

func NewServer(cfg *Config) (*Server, error) {
	producerch := make(chan Message)
	return &Server{
		Config:     cfg,
		Topics:     make(map[string]Storer),
		quitch:     make(chan struct{}),
		Producerch: producerch,
		producers: []Producer{
			NewHTTPproducer(cfg.ListenAddr, producerch),
		},
	}, nil
}

func (s *Server) Start() {
	// for _, consumer := range s.consumers {
	// 	if err := consumer.Start(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	for _, producer := range s.producers {
		go func(p Producer) {
			if err := p.Start(); err != nil {
				fmt.Println(err)
			}
		}(producer)
	}
	s.loop()
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL.Path)
// }

func (s *Server) loop() {
	fmt.Println("looping .. ")
	for {
		select {
		case <-s.quitch:
			return
		case msg := <-s.Producerch:
			fmt.Println("produced message -> ", msg)
			fmt.Printf("message data is %v and topic is %v", string(msg.Data), msg.Topic)

			offset, err := s.Publish(msg)
			if err != nil {
				slog.Error("failed to publish the message : ", err)
			} else {
				slog.Info("produced message offset :", "offset", offset)
			}
		}
	}

}

func (s *Server) Publish(msg Message) (int, error) {
	store := s.getStoreForTopic(msg.Topic)
	return store.Push(msg.Data)
}
func (s *Server) getStoreForTopic(topic string) Storer {
	if _, ok := s.Topics[topic]; !ok {
		s.Topics[topic] = s.StoreProducerFunc()
		slog.Info("created new topic : ", "topic", topic)
	}

	return s.Topics[topic]

}
