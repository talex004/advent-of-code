
package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"sort"
	"math/big"
)

func main() {
	a, b := readInput()
	sort.Ints(a)
	sort.Ints(b)
	println(score(a, b).String())
}

func readInput() (a []int, b []int) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	for _, ln := range lines {
		var v1, v2 int
		_, err := fmt.Sscanf(ln, "%d %d", &v1, &v2)
		if err != nil {
			log.Printf("Scanning line: %s", ln)
			continue
		}
		a = append(a, v1)
		b = append(b, v2)
	}
	return a, b
}

func sum(a, b []int) *big.Int {
	var res big.Int
	for i := range a {
		d := abs(a[i]-b[i])
		res.Add(&res, big.NewInt(d))
	}
	return &res
}

func score(a, b []int) *big.Int {
	var res big.Int
	for _, x := range a {
		mult := count(b, x)
		res.Add(&res, big.NewInt(int64(mult*x)))
	}
	return &res
}

func count(b []int, x int) int {
	i := sort.SearchInts(b, x)
	if i == len(b) || b[i] != x {
		return 0
	}
	res := 1
	for i+res != len(b) && b[i+res] == x {
		res++
	}
	return res
}

func abs(x int) int64 {
	if x < 0 {
		return int64(-x)
	}
	return int64(x)
}
