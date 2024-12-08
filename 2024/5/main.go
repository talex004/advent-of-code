package main

import (
  "io/ioutil"
  "log"
  "strings"
  "fmt"
  "strconv"
  "slices"
)

func main() {
  order, updates := readInput()
  
  sum, sum2 := 0, 0
  for _, upd := range updates {
    if isValid(upd, order) {
      sum += upd[len(upd)/2]
    } else {
      sum2 += reorder(upd, order)[len(upd)/2]
    }
  } 
  fmt.Println(sum, sum2)

}

func readInput() (map[[2]int]struct{}, [][]int) {
  o := map[[2]int]struct{}{}
  p := [][]int{}
  
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	stage := 0
	for _, ln := range lines {
	  if ln == "" {
	    stage++
	    continue
    }
	  if stage == 0 {
		  var v1, v2 int
		  _, err := fmt.Sscanf(ln, "%d|%d", &v1, &v2)
		  if err != nil {
			  log.Fatalf("Scanning line: %s", ln)
		  }
		  o[[2]int{v1, v2}] = struct{}{}
	  } else if stage == 1 {
	    words := strings.Split(ln, ",")
	    var lni []int
	    for _, w := range words {
	      i, err := strconv.Atoi(w)
	      if err != nil {
	        log.Fatalf("Converting %s: %v", w, err)
	      }
	      lni = append(lni, i)
	    }
	    p = append(p, lni)
	  }
	}
	return o, p
}

func isValid(upd []int, order map[[2]int]struct{}) bool {
  for i, p1 := range upd {
    for _, p2 := range upd[i+1:] {
      if _, ok := order[[2]int{p2, p1}]; ok {
        return false
      }
    }
  }
  return true
}

func reorder(upd []int, order map[[2]int]struct{}) []int {
  slices.SortFunc(upd, func(p1, p2 int) int {
    if _, ok := order[[2]int{p1, p2}]; ok {
      return -1
    }
    if _, ok := order[[2]int{p2, p1}]; ok {
      return 1
    }
    return 0
  })
  return upd
}
