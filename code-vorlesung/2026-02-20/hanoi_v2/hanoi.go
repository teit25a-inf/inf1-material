package hanoi_v2

import "fmt"

// Bewegt eine einzelne Platte von s nach z.
func Move(s, z string) {
	fmt.Printf("%s -> %s\n", s, z)
}

// Löst die Türme von Hanoi.
func Hanoi(h int, s, m, z string) {
	if h == 0 {
		return
	}
	Hanoi(h-1, s, z, m)
	Move(s, z)
	Hanoi(h-1, m, s, z)
}

// Bewegt einen Turm der Höhe 0 von s über m nach z.
func Hanoi0(s, m, z string) {
}

// Bewegt einen Turm der Höhe 1 von s über m nach z.
func Hanoi1(s, m, z string) {
	Hanoi0(s, z, m)
	Move(s, z)
	Hanoi0(m, s, z)
}

// Bewegt einen Turm der Höhe 2 von s über m nach z.
func Hanoi2(s, m, z string) {
	Hanoi1(s, z, m) // Parke die oberste Platte auf m.
	Move(s, z)      // Bewege die unterste Platte.
	Hanoi1(m, s, z) // Bewege die geparkte Platte nach z.
}

// Bewegt einen Turm der Höhe 3 von s über m nach z.
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m) // Parke den Turm der Höhe 2 von s auf m.
	Move(s, z)      // Bewege die unterste Platte.
	Hanoi2(m, s, z) // Bewege den geparkten Turm nach z.
}

// Bewegt einen Turm der Höhe 4 von s über m nach z.
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m) // Parke den Turm der Höhe 3 von s auf m.
	Move(s, z)      // Bewege die unterste Platte.
	Hanoi3(m, s, z) // Bewege den geparkten Turm nach z.
}

// Bewegt einen Turm der Höhe 5 von s über m nach z.
func Hanoi5(s, m, z string) {
	Hanoi4(s, z, m) // Parke den Teilturm von s auf m.
	Move(s, z)      // Bewege die unterste Platte.
	Hanoi4(m, s, z) // Bewege den geparkten Turm nach z.
}
