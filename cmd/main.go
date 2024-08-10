package main

import (
	"fmt"
	"httpwithlog/pkg/logging"
	"log/slog"
	"net/http"
)

var host string = "localhost:8080"

func main() {
	logging.LogInit()
	slog.SetDefault(logging.Logger)

	slog.Info(fmt.Sprintf("Server up with address: %v", host))

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerMain)

	if err := http.ListenAndServe(host, mux); err != nil {
		slog.Error(err.Error())
	}
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	slog.Info(fmt.Sprintf(r.RemoteAddr))
}
