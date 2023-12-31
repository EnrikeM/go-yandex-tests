package storage

import "github.com/EnrikeM/go-yandex-tests/internal/storage/common"

var Used common.Storage

func SetUsed(s *common.Storage) {
	Used = *s
}
