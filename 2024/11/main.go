package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "strconv"
)

type call struct {
  val int
  blinks int
}

var cache = map[call]int{}

func main() {
  stones := readInput()
  fmt.Println(stones)
  // memoize
  for i:=1; i<=75; i++ {
    for _, s := range stones {
      count(s, i)
    }
  }
  
  // count
  sum := 0
  for _, s := range stones {
    sum += count(s, 75)
  }
  fmt.Println("sum", sum)
}

func readInput() (res []int) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	parts := strings.Split(strings.TrimSpace(string(data)), " ")
	for _, s := range parts {
	  i, err := strconv.Atoi(s)
	  if err != nil { 
	    log.Fatal(err)
	  }
	  res = append(res, i)
	}
	return res
}

func count(val, blinks int) (res int) {
  if blinks == 0 {
    return 1
  }
  
  c := call{val, blinks}
  if cached, ok := cache[c]; ok {
    return cached
  }
  defer func() {
    cache[c] = res
  }()
  
  switch {
  case val == 0:
    return count(1, blinks-1)
  case len(strconv.Itoa(val)) % 2 == 0:
    vv := split(val)
    return count(vv[0], blinks-1) + count(vv[1], blinks-1)
  default:
    return count(val*2024, blinks-1)
  }
}

func split(v int) []int {
  s := strconv.Itoa(v)
  i := len(s)/2
  a, _ := strconv.Atoi(s[:i])
  b, _ := strconv.Atoi(s[i:])
  return []int{a, b}
}

