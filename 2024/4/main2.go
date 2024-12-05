package main

import (
  "io/ioutil"
  "log"
  "fmt"
  "slices"
  "strings"
)

func main() {
  s := readInput()
  res := 0
  for i := range s {
    for j := range s[i] {
      if match(s, i, j) {
        res++
      }
    }
  } 
  fmt.Println(res)
}

func readInput() []string {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool {
	  return len(s) == 0
  })
  return lines
}

func match(s []string, i, j int) bool {
  if i < 1 || i >= len(s)-1 {
    return false
  }
  if j < 1 || j >= len(s[i])-1 {
    return false
  }
  if s[i][j] != 'A' {
    return false
  }
  if ! ((s[i-1][j-1] == 'M' && s[i+1][j+1] == 'S') || (s[i-1][j-1] == 'S' && s[i+1][j+1] == 'M')) {
    return false
  }
  if ! ((s[i-1][j+1] == 'M' && s[i+1][j-1] == 'S') || (s[i-1][j+1] == 'S' && s[i+1][j-1] == 'M')) {
    return false
  }
  return true
}
