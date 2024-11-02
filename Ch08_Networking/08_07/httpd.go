// Limit value size to 1k
// key value db
/*
Key/Value DB
$ curl -d 'hello' http://localhost:8000/key

$ curl http://localhost:8000/k1
hello

$ curl http://localhost:8000/k2
404 Not Found
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	db DB
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Path[1:]
	value := s.db.Get(key)
	if value == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Write(value)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	defer r.Body.Close()

	rdr := io.LimitReader(r.Body, 1<<10)
	value, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}
	s.db.Set(key, value)
	fmt.Fprintf(w, "%s set \n", key)
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.postHandler(w, r)
		return
	case http.MethodGet:
		s.getHandler(w, r)
		return
	}

	http.Error(w, "Bad Method", http.StatusMethodNotAllowed)
}

func main() {
	s := Server{}
	http.HandleFunc("/", s.Handler)
	addr := ":8080"

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
