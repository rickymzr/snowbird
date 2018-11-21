package ca

type CA struct {
	name string
}

func New(name string) CA {

	return CA{name}
}

func (ca *CA) Name() string {
	return "ca"
}

func (ca *CA) Start() bool {
	return true
}
