package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	UTC string `json:"utc,omitempty"`
	JST string `json:"jst,omitempty"`
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	endpoint, ok := os.LookupEnv("TEST_ENDPOINT")
	if !ok {
		endpoint = "http://localhost:8080/json"
	}
	var avg int64

	for i := 0; i < 9; i++ {
		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		if err != nil {
			log.Fatal(err)
		}
		t := time.Now()
		resp, err := http.DefaultClient.Do(req)
		result := Response{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		/*
			utcTime, _ := time.ParseInLocation(time.RFC3339Nano, result.UTC, time.UTC)
		*/
		jstTime, _ := time.ParseInLocation(time.RFC3339Nano, result.JST, jst)

		fmt.Printf("%d: %dms\n", i+1, jstTime.Sub(t).Milliseconds())
		avg += jstTime.Sub(t).Milliseconds()
	}

	fmt.Printf("avg: %dms", avg)
}
