package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

var endpoint = os.Getenv("ENDPOINT")
var port = os.Getenv("PORT")

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/weather/{place}", func(w http.ResponseWriter, r *http.Request) {
		encPlace := chi.URLParam(r, "place")
		s := strings.NewReader(encPlace)

		rot := rot13Reader{s}
		decryptPlace := make([]byte, len(encPlace))
		if _, err := rot.Read(decryptPlace); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		resp, err := http.Get(fmt.Sprintf("%s/%s", endpoint, decryptPlace))
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		if resp.StatusCode != 200 {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
