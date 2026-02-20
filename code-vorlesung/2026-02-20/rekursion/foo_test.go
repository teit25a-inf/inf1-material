package rekursion

import "fmt"

// Foo berechnet 2^n.
func Foo(n int) int {
	result := 1

	// 2^0 == 1
	if n == 0 {
		return result
	}

	// 2^n == 2^(n-1) * 2
	foo_result := Foo(n - 1)
	result = foo_result * 2
	return result
}

func ExampleFoo() {
	fmt.Println(Foo(5))

	// Output:
}
