
package main

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {
	data := readInput()
  println(countSafe(data))	
}

func readInput() (res [][]int) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	for _, ln := range lines {
	  if ln == "" {
	    continue
    }
		words := strings.Split(ln, " ")
		var values []int
		for _, wd := range words {
		  x, err := strconv.Atoi(wd)
		  if err != nil {
		    log.Fatal(err)
		  }
		  values = append(values, x)
		}
		res = append(res, values)
	}
	return res
}

func countSafe(data [][]int) (res int) {
  for _, ln := range data {
    if isSafe(ln, -1) {
      res++
      continue
    }
    for i := 0; i<len(ln); i++ {
      if isSafe(ln, i) {
        res++
        break
      }
    }
  }
  return res
}

func isSafe(values []int, exclude int) bool {
  lastSign := 0
  last := values[0]
  if exclude == 0 {
    last = values[1]
  }
  for i, val := range values {
    if i == exclude {
      continue
    }
    if i == 0 || (i == 1 && exclude == 0) {
      continue
    }
    d := abs(val - last)
    if d > 3 || d < 1 {
      return false
    }
    sign := (val - last)/d
    if lastSign != 0 {
      if sign != lastSign {
        return false
      }
    }
    lastSign = sign
    last = val
  }
  return true
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}


