package main

func swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func main() {
	a := 1
	b := 2
	a, b = swap(a, b)
}
