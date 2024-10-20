package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"url-shortener/storage"

	"github.com/gorilla/mux"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var mu sync.Mutex

func generateShortURL() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	json.NewDecoder(r.Body).Decode(&request)

	originalURL := request["url"]
	shortURL := generateShortURL()

	mu.Lock()
	storage.Store(shortURL, originalURL)
	mu.Unlock()

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]

	originalURL := storage.Retrieve(shortURL)
	if originalURL != "" {
		http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
