package calc

// Add add two numbers
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

// Sub func to subtract two numbers
func Sub(a, b int) int {
	return a - b

}
