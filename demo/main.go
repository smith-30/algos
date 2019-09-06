package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeira/measure"
	"github.com/smith-30/algos/demo/helper"
)

const header = "Key,Count,Sum,Min,Max,Avg,Rate,P95"

func main() {
	defer measure.Start("sum").Stop()
	file, err := os.Create("file.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "%s\n", header)

	go func() {
		helper.Sum()
		helper.Sum()
	}()
	go func() {
		helper.Sum()
		helper.Sum()
	}()

	helper.Sub()
	helper.Sub()

	// api の口作ってもいいなぁ
	stats := measure.GetStats()
	stats.SortDesc("sum")

	// print stats in CSV format
	for _, s := range stats {
		fmt.Fprintf(file, "%s,%d,%f,%f,%f,%f,%.9f,%f\n",
			s.Key, s.Count, s.Sum, s.Min, s.Max, s.Avg, s.Rate, s.P95)
	}

	// api の口作ってもいいなぁ
	measure.Reset()
}
