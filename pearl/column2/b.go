package column2

// なるべく少ないメモリと速さになるようにしたが
// lとcountが近ければ近いほどメモリ割り当てが大きくなるのが辛い
// count が l / 2より大きければ逆回しとかやったほうが
// 割当量が減るなぁ。。多分。
func RotateTaskB(src []string, count int) {
	l := len(src)
	if l == 0 || count == 0 {
		return
	}

	var (
		c int
	)

	switch {
	case l == count:
		return
	case l < count:
		c = count % l
	case l > count:
		c = count
	}

	// 後ろに回す分のデータを新規に確保 (メモリ割り当てが起こる)
	rewind := make([]string, len(src[0:c]))
	copy(rewind, src[0:c])

	// 回転させないデータ領域を先頭にもってくる
	for idx, val := range src[c:l] {
		src[idx] = val
	}

	// 後ろに回す分のデータを整列させたデータの後ろにくっつける
	copy(src[l-c:l], rewind[0:])
}

func NewRotateTaskB(src []string, count int) {
	l := len(src)
	if l == 0 || count == 0 {
		return
	}

	var (
		c int
	)

	switch {
	case l == count:
		return
	case l < count:
		c = count % l
	case l > count:
		c = count
	}

	reverse(src, 0, c-1)
	reverse(src, c, l-1)
	reverse(src, 0, l-1)
}

func reverse(src []string, s, e int) {
	var i int
	for {
		s = s + i
		e = e - i
		if s < e {
			t := src[s]
			src[s] = src[e]
			src[e] = t
		} else {
			break
		}
		i += 1
	}
}
