package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":3000", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.ListenAndServe(*addr, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
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