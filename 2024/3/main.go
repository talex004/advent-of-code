package main

import (
  "io/ioutil"
  "log"
  "regexp"
  "fmt"
  "strconv"
)

func main() {
  s := readInput()
  re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`)
  match := re.FindAllStringSubmatch(s, -1)
  //fmt.Println(match)
  fmt.Println(calc(match))
}

func readInput() string {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	return string(data)
}

func calc(match [][]string) (res int) {
  do := true
  for _, m := range match {
    if m[0] == "don't()" {
      do = false
    } else if m[0] == "do()" {
      do = true
    } else if do {
      res += mul(m[1], m[2])
    }
  }
  return res
}

func mul(s1, s2 string) int {
  i1, _ := strconv.Atoi(s1)
  i2, _ := strconv.Atoi(s2)
  return i1*i2 
}
