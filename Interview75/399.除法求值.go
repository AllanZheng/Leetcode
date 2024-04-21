package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, cur := range equations {
		i, j := cur[0], cur[1]
		if _, has := id[i]; !has {
			id[i] = len(id)
		}
		if _, has := id[j]; !has {
			id[j] = len(id)
		}
	}
	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	for i, eq := range equations {
		fa, w = merge(id[eq[0]], id[eq[1]], values[i], w, fa)
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start, fa, w) == find(end, fa, w) {
			res[i] = w[start] / w[end]
		} else {
			res[i] = -2
		}
	}
	return res

}
func find(x int, fa []int, w []float64) int {
	if fa[x] != x {
		f := find(fa[x], fa, w)
		w[x] *= w[fa[x]]
		fa[x] = f
	}
	return fa[x]
}
func merge(from, to int, val float64, w []float64, fa []int) (fa1 []int, w1 []float64) {
	fFrom, fTo := find(from, fa, w), find(to, fa, w)
	fa1 = fa
	w1 = w
	w1[fFrom] = val * w[to] / w[from]
	fa[fFrom] = fTo
	return fa1, w1
}
