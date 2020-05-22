package interface6

type A interface {
	S()
}

type D struct {
}

func (d D) S() {

}

type C struct {
}

func (c *C) S() {

}

func main() {
	var a1 A = D{}
	var a2 A = &D{}
	var a3 A = C{}
	var a4 A = &C{}

	refect
}
