package helper

import (
	"time"

	"github.com/najeira/measure"
)

func Sum() int {
	defer measure.Start("sum").Stop()
	time.Sleep(20 * time.Millisecond)
	return 0
}

func Sub() int {
	defer measure.Start("sub").Stop()
	time.Sleep(40 * time.Millisecond)
	return 0
}
