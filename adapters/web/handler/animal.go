package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cassioleodegario/zooapi/domain"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeAnimalHandlers(r *mux.Router, n *negroni.Negroni, service domain.AnimalServiceInterface) {
	r.Handle("/animal/{id}", n.With(
		negroni.Wrap(getAnimal(service)),
	)).Methods("GET", "OPTIONS")
}

func getAnimal(service domain.AnimalServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		animal, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(animal)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
