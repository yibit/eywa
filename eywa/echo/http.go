package echo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HttpD struct {
	Port string
}

func (s *HttpD) Start(port string) {
	server := http.Server{
		Addr: ":" + port,
	}
	s.Port = port

	http.HandleFunc("/", s.Echo)

	log.Fatal(server.ListenAndServe())
}

func writeAck(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintln(w, message)
}

func (s *HttpD) Echo(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		writeAck(w, http.StatusBadRequest, "")
		return
	}

	setHeaderIfNotExist(w, r, "Content-Type", "text/plain")

	writeAck(w, http.StatusOK, string(data))
}

func setHeaderIfNotExist(w http.ResponseWriter, r *http.Request, header, value string) {
	if v := r.Header.Get(header); v != "" {
		w.Header().Set(header, r.Header.Get(header))
	} else {
		w.Header().Set(header, value)
	}
}
