package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {

	url := "http://localhost:3000/publish"
	topics := []string{"topic_1", "topic_2", "topic_3"}

	for i := 0; i < 100; i++ {
		topic := topics[rand.Intn(len(topics))]
		payload := []byte(fmt.Sprintf("foobarbaz_%d", i))

		res, err := http.Post(url+"/"+topic, "application/octet-stream", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			log.Fatal("status code is not 200")
		}
		fmt.Println(res)
	}
}
