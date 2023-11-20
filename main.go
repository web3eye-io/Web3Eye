package main

type xiaohu struct {
	A string
}

func (xh xiaohu) hhh() {
	println(xh.A)
	xh.hhhf()
}

func (xh xiaohu) hhhf() {
	println("hhhf ", xh.A)
}

type theshy struct {
	xiaohu
}

func (ts theshy) hehehe() {
	println("ssss ", ts.A)
}

func (ts theshy) hhhf() {
	println("sssass ", ts.A)
}

func main() {
	ts := theshy{xiaohu: xiaohu{A: "Sss"}}
	ts.hhh()
	ts.hehehe()
}
