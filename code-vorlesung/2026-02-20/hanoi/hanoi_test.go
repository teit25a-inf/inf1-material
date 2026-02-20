package hanoi

func ExampleMove() {
	Move("A", "C")
	Move("A", "B")
	Move("C", "B")
	Move("A", "C")
	Move("B", "A")
	Move("B", "C")
	Move("A", "C")

	// Output:
	// A -> C
	// A -> B
	// C -> B
	// A -> C
	// B -> A
	// B -> C
	// A -> C
}

func ExampleHanoi3() {
	Hanoi3("A", "B", "C")

	// Output:
	// A -> C
	// A -> B
	// C -> B
	// A -> C
	// B -> A
	// B -> C
	// A -> C
}
