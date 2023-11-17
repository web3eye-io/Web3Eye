package main

type xiaohu struct {
	A string
}

func (xh xiaohu) hhh() {
	println(xh.A)
}

type theshy struct {
	xiaohu
}

func (ts theshy) hehehe() {
	println("ssss ", ts.A)
}

func main() {
	ts := theshy{xiaohu: xiaohu{A: "Sss"}}
	ts.hhh()
	ts.hehehe()
}
