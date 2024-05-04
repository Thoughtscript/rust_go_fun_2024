package handlers

import (
	"encoding/json"
	"net/http"
	e "goserver/domain"
)

func GetExamples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		results :=  e.GetExamples()
		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
