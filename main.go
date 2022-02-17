package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log := log.New(os.Stdout, "", log.Lmicroseconds|log.LUTC|log.Ldate|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("startup", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Something worked, don't ask how."))
	})

	err := http.ListenAndServe(":8080", r)
	return err
}
