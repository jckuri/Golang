package my_package

func Square(x float64) float64 {
	return x * x
}

func Add64(x float64) float64 {
	return x + 64
}

type Result struct {
	number float64
}

func HundredDividedBy(x float64) *Result {
	if x == 0. {return nil}
	return &Result {100. / x}
}
