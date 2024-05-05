package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	j "goserver/jobs"
	"time"
)

// CreateJob - POST - create a worker and it to the queue
func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		url := r.URL.Query()
		// These are case-sensitive parsings
		user := url.Get("user")
		password := url.Get("password")
		cmd := url.Get("cmd")
		scheduled, err := time.Parse(time.RFC3339, url.Get("scheduled"))
		if err != nil {
			log.Println("Exception encountered - please supply a valid scheduled time!")
		}

		if j.VerifyPassword(user, password) {
			worker := j.NewWorker(scheduled, cmd)
			j.AddWorker(worker)
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(worker)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryPool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		url := r.URL.Query()
		user := url.Get("user")
		password := url.Get("password")

		if j.VerifyPassword(user, password) {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetAllJobs())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		url := r.URL.Query()
		user := url.Get("user")
		password := url.Get("password")

		if j.VerifyPassword(user, password) {
			uid := url.Get("uuid")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetJob(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func StopJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		url := r.URL.Query()
		user := url.Get("user")
		password := url.Get("password")

		if j.VerifyPassword(user, password) {
			uid := url.Get("uuid")
			w.WriteHeader(http.StatusCreated)
			j.StopWorker(uid)
			err := json.NewEncoder(w).Encode(j.GetStatus(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		url := r.URL.Query()
		user := url.Get("user")
		password := url.Get("password")

		if j.VerifyPassword(user, password) {
			uid := url.Get("uuid")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetStatus(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
