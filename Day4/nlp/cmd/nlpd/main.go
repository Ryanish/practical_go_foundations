package main

import (
	"encoding/json"
	"fmt"
	"gofoundation/day4/nlp"
	"io"
	"log"
	"net/http"

	"github.com/353solutions/nlp/stemmer"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(log.Writer(), "nlp ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger,
	}
	//routing
	// http.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/tokenize", tokenizeHandler)
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// run server
	addr := ":8080"
	s.logger.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)

}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	//before gorilla muxif
	// r.Method != http.MethodPost {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	rdr := io.LimitReader(r.Body, 1_000_000)
	data, err := io.ReadAll(rdr)
	if err != nil {
		s.logger.Printf("error: can't rerad - %s", err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	text := string(data)

	tokens := nlp.Tokenize(text)

	resp := map[string]any{
		"tokens": tokens,
	}

	data, err = json.Marshal(resp)
	if data != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)

}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
	fmt.Println(w, "OK")
}
