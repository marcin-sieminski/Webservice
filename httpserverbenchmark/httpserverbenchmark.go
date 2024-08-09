package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: 10 * time.Second,
	}
)

func api(w http.ResponseWriter, r *http.Request) {
	n := -1
	decoder := json.NewDecoder(r.Body)
	var payload Payload
	if decoder.Decode(&payload) == nil {
		n = payload.Value
	}
	w.Write([]byte(fmt.Sprintf("%d", n)))
}

func runServer(port int) {
	http.HandleFunc("/api-bench", api)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
}

func send(api string, value int, ch chan<- int) {
	ret := -1
	payload := Payload{Value: value}
	payloadBytes, _ := json.Marshal(payload)
	for {
		if resp, err := client.Post(api, "", bytes.NewReader(payloadBytes)); err == nil {
			if resp.StatusCode == 200 {
				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&ret)
				ch <- ret
				return
			}
		}
	}
}

type Payload struct {
	Value int `json:"value"`
}

func run() {
	sampleCount := 3000
	rand.Seed(time.Now().UTC().UnixNano())
	port := 20000 + rand.Intn(30000)
	go runServer(port)
	api := fmt.Sprintf("http://localhost:%d/api-bench", port)
	ch := make(chan int, sampleCount)
	for i := 1; i <= sampleCount; i++ {
		go send(api, i, ch)
	}
	sum := 0
	for i := 1; i <= sampleCount; i++ {
		sum += <-ch
	}
	fmt.Println(sum)
}

func main() {
	run()
}
