package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
)

type puzzle struct {
  keys [][5]int8
  locks [][5]int8
}

func main() {
  pzl := parseKeys(readInput(false))
  fmt.Println(pzl.part1())
}

func readInput(sample bool) (res [][]string) {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
  if err != nil {
	  log.Fatalf("Reading input: %v", err)
  }
  blocks := strings.Split(string(data), "\n\n")
  for _, block := range blocks {
    if len(block) < 3 {
      continue
    }
    res = append(res, strings.Split(string(block), "\n"))
  }
  return res
}

func parseKeys(blocks [][]string) (res puzzle) {
  for _, block := range blocks {
    k := [5]int8{}
    for _, row := range block[1:len(block)-1] {
      for j, ch := range row {
        if ch == '#' {
          k[j]++
        }
      }
    }
    if block[0] == "#####" {
      res.locks = append(res.locks, k)
    } else {
      res.keys = append(res.keys, k)
    }
  }
  return res
}

func (p *puzzle) part1() (res int) {
  for _, l := range p.locks {
    for _, k := range p.keys {
      if !overlap(l, k) {
        res++
      }
    }
  }
  return res
}

func overlap(lock, key [5]int8) bool {
  for i:=0; i<5; i++ {
    if lock[i]+key[i] > 5 {
      return true
    }
  }
  return false
}
