package summation

func Add(a, b interface{}) interface{} {

	switch x := a.(type) {
	case int:
		switch y := b.(type) {
		case int:
			return x + y
		case float64:
			return float64(x) + y
		default:
			panic("unsupported type for b")
		}
	case float64:
		switch y := b.(type) {
		case int:
			return x + float64(y)
		case float64:
			return x + y
		default:
			panic("unsupported type for b")
		}
	case string:
		switch y := b.(type) {
		case string:
			return x + y
		default:
			panic("unsupported type for b")
		}
	default:
		panic("unsupported type for a")
	}
	return nil
}
