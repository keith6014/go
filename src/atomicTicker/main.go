// pollPeriodic project main.go
package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func dummyFunc() string {
	return "yahoo"
}

func httpMainHandler(w http.ResponseWriter, v atomic.Value) {
	w.Write([]byte(v.Load().(string)))
}

func main() {
	pollTicker := time.NewTicker(time.Second * time.Duration(5))

	var pollValue atomic.Value
	pollValue.Store(string("as of now. no data. wait 5 secs"))

	go func(v *atomic.Value, tck *time.Ticker) {
		tick := 0
		for t := range tck.C {
			fmt.Println(tick, t)
			pollValue.Store(dummyFunc())
			tick++
		}
	}(&pollValue, pollTicker)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		httpMainHandler(w, pollValue)
	})

	http.ListenAndServe(":8080", mux)
}
