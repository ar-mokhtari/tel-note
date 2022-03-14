package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func menu(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is Server port 1212, Welcome")
	fmt.Fprintf(w, "\n --- List of service(s) ---\n")
	io.WriteString(w, ShowMenu())
}

func Serv() {
	http.HandleFunc("/", menu)
	log.Fatal(http.ListenAndServe(":1212", nil))
}
