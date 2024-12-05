package main

import (
  "io/ioutil"
  "log"
  "fmt"
  "slices"
  "strings"
)

var target1, target2 = "XMAS", "SAMX"

func main() {
  s := readInput()
  res := 0
  res += countHorizontal(s)
  res += countVertical(s)
  res += countDiagonalUp(s)
  res += countDiagonalDown(s)
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

func countHorizontal(d []string) (res int) {
  for _, s := range d {
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  return res
}

func countVertical(d []string) (res int) {
  var tmp strings.Builder
  tmp.Grow(len(d))
  for i := range d[0] {
    tmp.Reset()
    for _, line := range d {
      tmp.WriteByte(line[i])
    }
    s := tmp.String()
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  return res
}

func countDiagonalUp(d []string) (res int) {
  var tmp strings.Builder
  tmp.Grow(len(d))
  
  toStr := func(i, j int) string {
    tmp.Reset()
    for k:=0; i-k>=0 && j+k < len(d[0]); k++ {
      tmp.WriteByte(d[i-k][j+k])
    }
    return tmp.String()
  }
  for i:=0; i<len(d); i++ {
    s := toStr(i, 0)
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  for j:=1; j<len(d[0]); j++ {
    s := toStr(len(d)-1, j)
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  return res
}

func countDiagonalDown(d []string) (res int) {
  var tmp strings.Builder
  tmp.Grow(len(d))
  
  toStr := func(i, j int) string {
    tmp.Reset()
    for k:=0; i+k<len(d) && j+k < len(d[0]); k++ {
      tmp.WriteByte(d[i+k][j+k])
    }
    return tmp.String()
  }
  for i:=0; i<len(d); i++ {
    s := toStr(i, 0)
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  for j:=1; j<len(d[0]); j++ {
    s := toStr(0, j)
    res += strings.Count(s, target1)
    res += strings.Count(s, target2)
  }
  return res
}
