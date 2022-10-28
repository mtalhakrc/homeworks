package main

func main() {
	sd := asd{xd: 10}
	sd.a()
}

type test interface {
	a() test
}
type asd struct {
	xd int
}

func (asd) a() test {
	xd := asd{xd: 1}
	return xd
}
