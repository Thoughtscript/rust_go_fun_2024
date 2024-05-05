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

		results := e.GetExamples()
		err := json.NewEncoder(w).Encode(results)
		
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)

		url := r.URL.Query()
		id := url.Get("id")
		results := e.GetExample(id)

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UpdateExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPut {
		w.WriteHeader(http.StatusOK)

		url := r.URL.Query()
		id := url.Get("id")
		name := url.Get("name")
		val := url.Get("val")
		results := e.UpdateExample(id, name, val)

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DeleteExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {
		w.WriteHeader(http.StatusOK)

		url := r.URL.Query()
		id := url.Get("id")
		results := e.DeleteExample(id)

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func CreateExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)

		url := r.URL.Query()
		id := url.Get("id")
		name := url.Get("name")
		val := url.Get("val")
		results := e.CreateExample(id, name, val)

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
