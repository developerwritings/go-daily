package main

type Parent struct {
	Child Child
}

type Child struct {
	name string
}

func main() {

	slics := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// add 1000 to slice

	for i := 0; i <= 1000; i++ {
		slics = append(slics, i)
	}

	// removed first 10 elements

	for i := 0; i <= 10; i++ {
		slics = append(slics, slics[:i])
	}

	// even numbers
	s2 := make([]int{})
	for i := 0; i <= len(slics); i++ {
		if slics[i]/2 == 0 {
			s2 = append(slics, slics[i])
		}
	}

}
