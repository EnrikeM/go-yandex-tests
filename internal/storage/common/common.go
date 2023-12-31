package common

type Storage interface {
	Get(k string) (string, error)
	Set(k, v string)
}
