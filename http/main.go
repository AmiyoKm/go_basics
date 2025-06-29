package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var store []*Users

var mu sync.Mutex
var rw sync.RWMutex

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		content := map[string]any{
			"success": true,
			"content": "HELLO WORLD",
		}
		j, err := json.MarshalIndent(content, " ", "\t")
		if err != nil {
			log.Println("Invalid json")
		}

		j = append(j, '\n')
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})
	// Handle can only take func(w , r) if it is wrapped by http.HandlerFunc.
	// http.HandlerFunc() wraps the given function by calling it inside a ServeHTTP()
	// which satisfies the Handler interface
	mux.Handle("POST /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user Users

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request Error"))
			return
		}

		mu.Lock()

		prevID := 0
		if len(store) > 0 {
			prevID = store[len(store)-1].ID
		}
		user.ID = prevID + 1
		store = append(store, &user)

		mu.Unlock()

		j, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}))

	// http.HandlerFunc automatically wraps the given function to be a Handler by wrapping it by a ServerHTTP()
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		rw.RLock()
		users := store
		rw.RUnlock()

		j, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {

		pathID := r.PathValue("id")
		ID, err := strconv.Atoi(pathID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request Error"))
			return
		}
		rw.RLock()
		defer rw.RUnlock()
		for _, user := range store {
			if ID == user.ID {
				j, err := json.MarshalIndent(user, " ", "\t")
				if err != nil {
					log.Println("Invalid json")
				}

				j = append(j, '\n')
				w.Header().Set("Content-Type", "application/json")
				w.Write(j)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))

	})

	mux.HandleFunc("PATCH /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		pathID := r.PathValue("id")
		ID, err := strconv.Atoi(pathID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID cant be transformed to int"))
			return
		}
		var payload struct {
			Name  *string `json:"name,omitempty"`
			Email *string `json:"email,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request Error"))
			return
		}

		mu.Lock()
		defer mu.Unlock()
		for _, user := range store {
			if ID == user.ID {

				if payload.Email != nil {
					user.Email = *payload.Email
				}
				if payload.Name != nil {
					user.Name = *payload.Name
				}

				j, err := json.MarshalIndent(user, " ", "\t")
				if err != nil {
					log.Println("Invalid json")
				}

				j = append(j, '\n')
				w.Header().Set("Content-Type", "application/json")
				w.Write(j)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		pathID := r.PathValue("id")
		ID, err := strconv.Atoi(pathID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID cant be transformed to int"))
			return
		}

		mu.Lock()
		defer mu.Unlock()

		for i, user := range store {
			if ID == user.ID {
				store = append(store[:i], store[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	})

	http.ListenAndServe(":8080", mux)
}
