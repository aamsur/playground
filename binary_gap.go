package playground

import (
	"strconv"
	"strings"
)

// n adalah input number yang akan dijadikan bilangan binar dan akan di hitung banyak gapnya
func binGap(n int64) (r int) {
	splits := strings.Split(strconv.FormatInt(n, 2), "1")

	if len(splits) < 3 {
		return 0
	}

	for _, s := range splits {
		if len(s) > r {
			r = len(s)
		}
	}

	return
}
