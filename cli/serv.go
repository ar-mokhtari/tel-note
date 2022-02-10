package cli

import (
	"fmt"
	"log"
	"net/http"
)

func Serv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is Server port 1212, Welcome")
	})
	log.Fatal(http.ListenAndServe(":1212", nil))
}
