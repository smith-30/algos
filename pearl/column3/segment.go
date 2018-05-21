package column3

import (
	"bytes"
	"fmt"
	"sync"
)

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
)

func Set(b, flag Bits) Bits {
	return b | flag
}
func Clear(b, flag Bits) Bits {
	return b &^ flag
}
func Toggle(b, flag Bits) Bits {
	return b ^ flag
}
func Has(b, flag Bits) bool {
	return b&flag != 0
}

type (
	Segmenter struct {
		light []byte
	}
)

func showSegment(src [5]byte) {
	segChan := make(chan Segmenter, len(src))
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for s := range segChan {
			s.Flush()
		}
	}()

	for _, val := range src {
		var Seg Segmenter
		for _, flag := range []Bits{F0, F1, F2, F3, F4, F5, F6} {
			if Has(Bits(val), flag) {
				Seg.light = append(Seg.light, 1)
			} else {
				Seg.light = append(Seg.light, 0)
			}
		}
		segChan <- Seg
	}
	close(segChan)
	wg.Wait()
}

func (s Segmenter) Flush() {
	switch {
	case bytes.Equal(s.light, []byte{1, 0, 1, 1, 1, 1, 1}):
		fmt.Print(0)
	case bytes.Equal(s.light, []byte{0, 0, 0, 1, 0, 1, 0}):
		fmt.Print(1)
	case bytes.Equal(s.light, []byte{1, 1, 1, 0, 1, 1, 1}):
		fmt.Print(2)
	case bytes.Equal(s.light, []byte{1, 1, 1, 0, 1, 0, 1}):
		fmt.Print(3)
	case bytes.Equal(s.light, []byte{0, 1, 0, 1, 1, 0, 1}):
		fmt.Print(4)
	case bytes.Equal(s.light, []byte{1, 1, 1, 1, 0, 0, 1}):
		fmt.Print(5)
	case bytes.Equal(s.light, []byte{1, 1, 1, 1, 0, 1, 1}):
		fmt.Print(6)
	case bytes.Equal(s.light, []byte{0, 0, 1, 0, 1, 0, 1}):
		fmt.Print(7)
	case bytes.Equal(s.light, []byte{1, 1, 1, 1, 1, 1, 1}):
		fmt.Print(8)
	case bytes.Equal(s.light, []byte{0, 1, 1, 1, 1, 0, 1}):
		fmt.Print(9)
	default:
		fmt.Print(" ")
	}
}
