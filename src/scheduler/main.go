package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync/atomic"
	"time"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

type Tasks struct {
	Task []Task
	Id   int
}

type Task struct {
	Ticker *time.Ticker
	cmd    string
	output atomic.Value
}

func (t *Tasks) ExecuteAll() {
	for i := range t.Task {
		log.Printf("Starting task %d %s\n", i, t.Task[i].cmd)
		go t.Task[i].Run()
	}
}

func (t *Task) Run() {
	for range t.Ticker.C {
		log.Println("Executing", t.cmd)
		out, err := exec.Command(t.cmd).Output()
		if err != nil {
			log.Println(err)
		}
		t.output.Store(string(out[:]))
	}
}

func (t *Task) hello(w http.ResponseWriter, r *http.Request) {
	val := t.output.Load().(string)
	fmt.Fprintf(w, "hello. The uptime is %s\n", val)
}

func CreateTask(duration time.Duration, Command string) Task {
	t := Task{
		Ticker: time.NewTicker(duration),
		cmd:    Command,
	}
	t.output.Store("No output yet")
	return t
}

func main() {

	A := Tasks{Task: []Task{
		CreateTask(time.Second*7, "ps"),
		CreateTask(time.Second*3, "uptime"),
		CreateTask(time.Second*4, "hostname"),
	}}

	A.ExecuteAll()

	log.Println("Starting webserver")
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", A.Task[0].hello)

	fs := http.FileServer(http.Dir("swagger-ui/dist"))
	mux.Handle("/api/", http.StripPrefix("/api/",
		fs))

	err := http.ListenAndServe(":3000", Log(mux))
	log.Println(err)
}
