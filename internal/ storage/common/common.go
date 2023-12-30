package common

// Здесь создадим интерфейс Storage общий для всех наших хранилищ
type Storage interface {
	Get(k string) (string, error)
	Set(k, v string)
}
