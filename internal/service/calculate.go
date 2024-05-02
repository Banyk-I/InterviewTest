package service

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

type FactorialResult struct {
	Factorial int
	Name      string
}

// test
func (c *Calculator) CalculateFactorials(a, b int) (int, int) {
	ch := make(chan FactorialResult)
	go factorial(a, ch, "a")
	go factorial(b, ch, "b")

	resultA, resultB := <-ch, <-ch
	if resultA.Name != "a" {
		resultA, resultB = resultB, resultA
	}

	return resultA.Factorial, resultB.Factorial
}

func factorial(n int, ch chan<- FactorialResult, name string) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	ch <- FactorialResult{
		Factorial: fact,
		Name:      name,
	}
}
