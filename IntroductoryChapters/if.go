package main

// Basic if
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func more(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

// If with precondition.
func maxAreaCheck(length, width, limit int) bool {
	if area := length * width; area < limit {
		return true
	} else {
		return false
	}
}
func main() {

}
