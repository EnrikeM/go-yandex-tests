// пакеты исполняемых приложений должны называться main
package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const shortenerlen = 7

var Urls = map[string]string{}

// функция shorten - создаёт рандомную ссылку.
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

	//Читаем содержимое запроса

	// Создаём мапу, в которой будем хранить пару: длинная ссылка: короткая ссылка

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

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		bodystr := string(body)
		shortURL := shorten()

		//присваиваем длинной ссылку короткую
		Urls[shortURL] = bodystr

		if err := req.ParseForm(); err != nil {
			res.Write([]byte(err.Error()))
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(shortURL))
		return
	}

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res.WriteHeader(http.StatusTemporaryRedirect)
	res.Write([]byte(Urls[p[len(p)-1]]))

}

//Что осталось - сохранить то, что передаём в POST в переменную, чтобы потом отдать
// Мб тут нужно создать мапу: при пост запросе сохраняем длинную ссылку как ключ и значение как короткую
// При гет запросе по значению маленькой ссылки (значения) выдать длинную (ключ)

//func encrypt - /: POST        curl -v -X POST 'https://practicum.yandex.ru/ '
//func decrypt - /{id}: GET
