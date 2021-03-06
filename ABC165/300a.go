package main

import "fmt"

type four struct {
	a int
	b int
	c int
	d int
}

func makeArray(n int, m int, p []four) int {
	var temp, max int
	an := make([]int, n)
	for i := 0; i < n; i++ {
		an[i]++
	}

	for {
		flag := true
		for i := n - 1; i >= 0; i-- {
			if an[i] < m {
				an[i]++
				for j := i + 1; j < n; j++ {
					an[j] = an[i]
				}
				flag = false
				break
			}
		}
		temp = calc(an, p)
		if max < temp {
			max = temp
		}
		if flag {
			break
		}
	}
	return max
}

func calc(an []int, p []four) int {
	n := len(p)
	temp := 0
	for i := 0; i < n; i++ {
		if an[p[i].b-1]-an[p[i].a-1] == p[i].c {
			temp += p[i].d
		}
	}
	return temp
}

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &q)
	p := make([]four, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&q[i].a, &p[i].b, &p[i].c, &p[i].d)
	}
	fmt.Println(makeArray(n, m, p))

}
