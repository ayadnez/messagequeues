package main

import (
	"fmt"
	"testing"
)

func TestStorage(t *testing.T) {
	m := NewMemoryStorage()

	go func() {
		for i := range 10 {
			key := fmt.Sprintf("foobarbaz_%d", i)
			latestOffset, err := m.Push([]byte(key))
			if err != nil {
				t.Error(err)
			}

			data, err := m.Fetch(latestOffset)
			if err != nil {
				t.Error(err)
			}

			fmt.Println(string(data))
		}
	}()

}
