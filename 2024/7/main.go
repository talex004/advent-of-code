package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "strconv"
  "log"
)

func main() {
  input := readInput()
  fmt.Println(part1(input))
}

type equation struct {
  result int
  terms []int
}

func readInput() (res []equation) {
  data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
  for _, ln := range lines {
    if ln == "" {
      continue
    }
    var eq equation
    pair := strings.Split(ln, ": ")
    eq.result, _ = strconv.Atoi(pair[0])
    terms := strings.Split(pair[1], " ")
    for _, term := range terms {
      ti, _ := strconv.Atoi(term) 
      eq.terms = append(eq.terms, ti)
    }
    res = append(res, eq)
  }	
  return res
}

func part1(eqs []equation) (sum int) {
  for _, eq := range eqs {
    if isValid1(eq) {
      sum += eq.result
    }
  } 
  return sum
}

func isValid1(eq equation) bool {
  ops := make([]byte, len(eq.terms)-1)
  var recurse func(int) bool
  recurse = func(i int) bool {
    if i == len(eq.terms)-1 {
      if calc(eq.terms, ops) == eq.result {
        return true
      }
      return false
    }
    for _, op := range []byte{'+', '*', '|'} {
      ops[i] = op
      if recurse(i+1) {
        return true
      }
    }
    return false
  }
  return recurse(0)
}

func calc(terms []int, ops []byte) int {
  res := terms[0]
  for i, t := range terms {
    if i == 0 {
      continue
    }
    if ops[i-1] == byte('+') {
      res = res + t
    } else if ops[i-1] == byte('*') {
      res = res * t
    } else if ops[i-1] == byte('|') {
      res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, t))
    }
  }
  return res
}
