package utils

import (
	"net/http"
	"time"
)

func TimedLoop(delay time.Duration, interval time.Duration, times int, body func()) {
	time.Sleep(delay)
	for i := 0; i < times; i++ {
		body()
		time.Sleep(interval)
	}
}

func HttpGet(url string) string {
	r, _ := http.Get(url)
	if r == nil {
		return ""
	}
	body := make([]byte, 1024)
	r.Body.Read(body)
	return string(body)
}
