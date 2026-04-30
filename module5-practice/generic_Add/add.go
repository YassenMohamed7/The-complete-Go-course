package genericAdd

func GenericAdd[T int | float64 | string](a, b T) T {
	return a + b
}
