package service

type Service interface {
	Name() string
	Start() bool
}
