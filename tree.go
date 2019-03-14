package playground

type Relation struct {
	parent int64
	child  int64
}

// i adalah beberapa input relation, relation terdiri dari parent dan child
// r adalah result, r adalah jumlah parent atau jumlah tree yang terbuat dari input beberapa i
// program ini akan memproses hanya maksimal satu child saja
// contoh :
//    0          3
//   /  \          \
//   1   2          4
// input {0, 1}, {0, 2}, {3, 4}
// r akan menghasilkan 2, karena terdapat 2 tree atau 2 parent saja yaitu 0 dan 3
func findTree(i ...Relation) (r int64) {
	var rmap = make(map[int64]int64)
	for _, s := range i {
		rmap[s.parent] = s.child
	}

	r = int64(len(rmap))

	return
}
