package main

import "fmt"

func main() {
	hanoi(3, "A", "B", "C")
}

func hanoi(disk int, src, tmp, dst string) {
	if disk > 0 {
		hanoi((disk - 1), src, dst, tmp)
		fmt.Println(src, " -> ", dst)
		hanoi((disk - 1), tmp, src, dst)
	}
}
