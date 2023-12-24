// пакеты исполняемых приложений должны называться main
package main

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	if err := run(); err != nil {
		panic(err)
	}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const shortenerlen = 7

// функция shortenet - создаёт рандомную ссылку.
func shorten() string {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	b := make([]rune, shortenerlen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	http.HandleFunc(`/`, mainPage)
	return http.ListenAndServe(`:8080`, http.HandlerFunc(mainPage))
}

func mainPage(res http.ResponseWriter, req *http.Request) {

	p := strings.Split(req.URL.Path, "/") // ["", ""]

	if len(p) > 2 {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	if len(p[len(p)-1]) == 0 {
		if req.Method != http.MethodPost {
			res.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		shortURL := shorten()
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(shortURL))
		return
	}
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res.WriteHeader(http.StatusTemporaryRedirect)
	res.Write([]byte("https://practicum.yandex.ru/"))
}

//func encrypt - /: POST        curl -v -X POST 'https://practicum.yandex.ru/ '
//func decrypt - /{id}: GET
