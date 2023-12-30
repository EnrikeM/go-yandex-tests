package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	storage "github.com/EnrikeM/go-yandex-tests/internal/storage"
)

const (
	ShortenPath  = "/"
	shortenerlen = 7
)

var letters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) > 2 {
		http.Error(w, "", http.StatusNotFound)
	}

	switch {
	case len(path[len(path)-1]) == 0:
		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		long, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read request body "+err.Error(), http.StatusBadRequest)
			return
		}
		short := getShort()
		storage.Used.Set(short, string(long))
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", short)))
	default:
		if r.Method != http.MethodGet {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		short := path[len(path)-1]
		long, err := storage.Used.Get(short)
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		w.Header().Set("Location", long)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func getShort() string {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	b := make([]rune, shortenerlen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
