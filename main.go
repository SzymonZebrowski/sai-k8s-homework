package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	stressTime, err := strconv.ParseInt(r.URL.Query().Get("time"), 0, 64)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	n := 4
	quit := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Duration(stressTime) * time.Second)
	for i := 0; i < n; i++ {
		quit <- true
	}

	fmt.Fprintf(w, "Hello world from %s", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/stress", Handler)
	http.ListenAndServe(":8000", nil)
}
