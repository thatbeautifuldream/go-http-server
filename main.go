package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var addr = flag.String("addr", getEnv("PORT", ":3000"), "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/events", events)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(*addr, nil)
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"this", "is", "an", "event", "stream", "from", "milind"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// intentional delay
		time.Sleep(420 * time.Millisecond)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return ":" + value
	}
	return fallback
}