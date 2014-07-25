package main

func init() {
	var winners int

	// This solution uses a recursive tree walking function that only explores valid paths.
	answers[192] = func() {
		var walker func(depth, l, a int)

		walker = func(depth, l, a int) {
			if depth > 30 {
				winners++
				return
			}

			walker(depth+1, l, 0) // represents an O
			if a < 2 {
				walker(depth+1, l, a+1) // represents an A
			}
			if l != 1 {
				walker(depth+1, 1, 0) // represents an L
			}
		}

		walker(1, 0, 0)
		fml(winners)
	}

}
