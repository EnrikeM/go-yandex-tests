package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/EnrikeM/go-yandex-tests/internal/app/handlers"
	storage "github.com/EnrikeM/go-yandex-tests/internal/storage"
	"github.com/EnrikeM/go-yandex-tests/internal/storage/common"
	"github.com/EnrikeM/go-yandex-tests/internal/storage/inmem"
)

func main() {
	im := common.Storage(inmem.New())
	storage.SetUsed(&im)

	http.HandleFunc(handlers.ShortenPath, handlers.HandlerFunc)
	if err := http.ListenAndServe(":8080", http.HandlerFunc(handlers.HandlerFunc)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
