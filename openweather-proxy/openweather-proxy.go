package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var apiKey = os.Getenv("OWM_API_KEY")
var port = os.Getenv("PORT")

type Response struct {
	Temp float64 `json:"temp"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/weather/{place}", func(w http.ResponseWriter, r *http.Request) {
		owm, err := owm.NewCurrent("C", "en", apiKey)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		place := chi.URLParam(r, "place")

		if err = owm.CurrentByName(place); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		resp, err := json.Marshal(Response{Temp: owm.Main.Temp})
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
		if _, err := w.Write(resp); err != nil {
			log.Println(err)
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
