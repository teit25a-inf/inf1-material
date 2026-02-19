package rekursion

import "fmt"

// Add berechnet x + y
func Add(x, y int) int {
	// x + 0 = x
	if y == 0 {
		return x
	}

	// x + y = x+(y-1) + 1
	return Add(x, y-1) + 1
}

// Add2 berechnet x + y
func Add2(x, y int) int {
	if y == 0 {
		return x
	}

	return Add2(x+1, y-1)
}

// Add3 berechnet x + y
func Add3(x, y int) int {
	for y != 0 {
		x++
		y--
	}

	return x
}

func ExampleAdd() {
	fmt.Println(Add(2, 3))

	// Output:
	// 5
}
