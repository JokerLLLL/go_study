package adder

func Adder() func() int {
	sum := 0
	return func() int {
		 sum += 1
		 return sum
	}
}

type iAdder func(value int) (int, iAdder)

func Adder2(base int) iAdder {
	return func(value int) (i int, adder iAdder) {
		return base + value, Adder2(base + value)
	}
}
