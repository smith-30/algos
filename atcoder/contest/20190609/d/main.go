package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, p int
	fmt.Scan(&n)
	fmt.Scan(&p)

	var max int

	reader := bufio.NewScanner(os.Stdin)

	dan := make([][]string, n, n)
	l, r, d, t := [2000][2000]int{}, [2000][2000]int{}, [2000][2000]int{}, [2000][2000]int{}
	for i := 0; i < n; i++ {
		reader.Scan()
		dan[i] = strings.Split(reader.Text(), "")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			if dan[i][j] == "#" {
				d[i][j] = 0
				l[i][j] = 0
			} else {
				if i == 0 {
					d[i][j] = 1
				} else {
					d[i][j] = d[i-1][j] + 1
				}
				if j == 0 {
					l[i][j] = 1
				} else {
					l[i][j] = l[i][j-1] + 1
				}
			}
		}
	}

	nn := n - 1
	pp := p - 1
	for i := nn; i >= 0; i-- {
		for j := pp; j >= 0; j-- {
			if dan[i][j] == "#" {
				r[i][j] = 0
				t[i][j] = 0
			} else {
				if j == pp {
					r[i][j] = 1
				} else {
					r[i][j] = r[i][j+1] + 1
				}
				if i == nn {
					t[i][j] = 1
				} else {
					t[i][j] = t[i+1][j] + 1
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			if dan[i][j] != "#" {
				v := l[i][j] + r[i][j] + t[i][j] + d[i][j] - 3
				if v > max {
					max = v
				}
			}
		}
	}

	fmt.Printf("%#v\n", max)
}
