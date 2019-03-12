package playground

import (
	"strconv"
	"strings"
)

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
