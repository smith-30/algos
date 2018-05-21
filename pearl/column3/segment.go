package column3

import (
	"fmt"
	"math/bits"
)

type (
	Segmenter struct {
	}
)

func showSegment(src [5]byte) {
	for _, val := range src {
		fmt.Printf("%#v\n", bits.OnesCount8(val))
	}
}
