package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	bind := ""
  message := ""

	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":8080", "Socket to bind.")
	flagset.StringVar(&message, "message", "Hello simple app", "Text for display.")
	flagset.Parse(os.Args[1:])

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})
	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(bind, nil))
}
