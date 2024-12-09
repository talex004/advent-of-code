package main

import (
  "io/ioutil"
  "log"
  "strings"
  "fmt"
  "math"
)

type coord struct {
  x, y int
}

var board coord

func main() {
  antennas := readInput()
  fmt.Println(part1(antennas))
}

func readInput() map[rune][]coord {
  res := map[rune][]coord{}
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	for i, ln := range strings.Split(string(data), "\n") {
	  for j, fq := range ln {
	    if fq == '.' || fq == '#' {
	      continue
      }
      res[fq] = append(res[fq], coord{i, j})
	  }
	  if len(ln) > 0 {
	    board.y = len(ln)
	    board.x = i+1
    }
	}
	return res
}

func part1(antennas map[rune][]coord) int {
  all := map[coord]struct{}{}
  for _, coords := range antennas {
     all = merge(all, allAntinodes(coords))
  }
  return len(all)
}

func allAntinodes(in []coord) map[coord]struct{} {
  res := map[coord]struct{}{}
  for i, a := range in {
    for j, b := range in {
      if j <= i {
        continue
      }
      for _, n1 := range antinodes2(a, b) {
        if isValid(n1) {
          res[n1] = struct{}{}
        }
      }
    }
  }
  return res
}

func merge(a, b map[coord]struct{}) map[coord]struct{} {
  res := a
  for x := range b {
    res[x] = struct{}{}
  }
  return res
}

func isValid(a coord) bool {
  if a.x < 0 || a.y < 0 {
    return false
  }
  if a.x >= board.x || a.y >= board.y {
    return false
  }
  return true
}

func antinodes1(a, b coord) []coord {
  if a.x > b.x {
    a, b = b, a
  }
  dx := abs(a.x-b.x)
  dy := abs(a.y-b.y)
  minx, maxx := minmax(a.x, b.x)
  miny, maxy := minmax(a.y, b.y)
  if a.y < b.y {
    return []coord{{minx-dx, miny-dy}, {maxx+dx, maxy+dy}}
  }
  return []coord{{minx-dx, maxy+dy}, {maxx+dx, miny-dy}}
}

func antinodes2(a, b coord) (res []coord) {
  if a.x > b.x {
    a, b = b, a
  }
  if a.x == b.x {
    for i:=0; i<board.y; i++ {
      res = append(res, coord{a.x, i})
    }
    return res
  }
  if a.y == b.y {
    for i:=0; i<board.x; i++ {
      res = append(res, coord{i, a.y})
    }
    return res
  }
  for y:=0; y<board.y; y++ {
    x := float64(b.x) - float64(b.y-y)*float64(b.x-a.x)/float64(b.y-a.y)
    if math.Abs(x - math.Round(x)) < 0.00001 {
      res = append(res, coord{int(math.Round(x)), y})
    }
  }
  return res
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func minmax(a, b int) (int, int) {
  if a < b {
    return a, b
  }
  return b, a
}
