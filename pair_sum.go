package playground

import "strconv"

// `inputNumbers` adalah array number yang akan dicari pairnya
// `find` adalah angka yang akan dicari pairnya (yang jika disumkan akan menghasilkan sejumlah angka `find`) di dalam `inputNumbers`
func pairOfSum(inputNumbers []int64, find int64) (r string) {

	maxElement := len(inputNumbers)
	for k, n := range inputNumbers {
		if k == maxElement {
			break
		}

		if n + inputNumbers[k+1] == find {
			return "first pair is " + strconv.FormatInt(n, 10) + " and " + strconv.FormatInt(inputNumbers[k+1], 10)
		}
	}

	return "not found"
}
