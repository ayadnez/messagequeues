package main

import (
	"log"
)

func main() {

	cfg := &Config{
		ListenAddr: ":3000",
		// StoreProducerFunc: func() Storer { return NewMemoryStorage() },
		StoreProducerFunc: func() Storer { return NewMemoryStorage() },
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(s)

	// offset, _ := s.StoreProducerFunc().Push([]byte("message1"))

	// data, err := s.StoreProducerFunc().Fetch(offset)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))
	s.Start()

}
