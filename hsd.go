// hsdは文字列の距離を計算する関数を提供します
package hsd

// 文字列lとrの距離を返却する
func StringDistance(l, r string) int {
	return Distance([]rune(l), []rune(r))
}

func Distance(a, b []rune) int {
	if len(a) != len(b) {
		return -1
	}

	dist := 0
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}
