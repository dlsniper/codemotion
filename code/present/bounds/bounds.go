package bounds

func bound() {
	a := []int{1,2,3,4}
	for i:=0; i < len(a); i++ {
		_ = a[i]
	}
}

func bound2() {
	a := []int{1,2,3,4}
	for i := range a {
		_ = a[i]
	}
}

func main() {
	bound()
	bound2()
}