package playground

func reverse(n int64) (r int64) {
	for n > 0 {
		r = r*10 + n%10

		n = n / 10
	}

	return
}