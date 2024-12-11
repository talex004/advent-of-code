package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "slices"
)

type coord struct {
  x, y int
}

func main() {
  grid := readInput()
  fmt.Println(part1(grid))
  fmt.Println(part2(grid))
}

func readInput() (res [][]byte) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	for _, ln := range strings.Split(string(data), "\n") {
	  if len(ln) == 0 {
	    continue
	  }
	  lni := make([]byte, len(ln))
	  for i, ch := range ln {
	    lni[i] = byte(ch-'0')
	  }
    res = append(res, lni)
	}
	return res
}

func part1(grid [][]byte) (score int) {
  dests := findDests(grid)
  heads := map[coord][]coord{}
  for _, d := range dests {
    floodFill(grid, heads, d)
  }
  for _, head := range heads {
    score += len(head)
  }
  return score
}

func part2(grid [][]byte) (rating int) {
  dests := findDests(grid)
  heads := map[coord]int{}
  for _, d := range dests {
    floodFill2(grid, heads, d)
  }
  for _, v := range heads {
    rating += v
  }
  return rating
}

func findDests(grid [][]byte) (res []coord) {
  for i, row := range grid {
    for j, h := range row {
      if h == 9 {
        res = append(res, coord{i, j})
      }
    }
  }
  return res
}

func floodFill(grid [][]byte, heads map[coord][]coord, dest coord) {
  seen := map[coord]struct{}{}
  
  var visit func(coord)
  visit = func(p coord) {
    if _, ok := seen[p]; ok {
      return
    }
    seen[p] = struct{}{}
    for _, n := range next(grid, p) {
      if grid[n.x][n.y] == 0 {
        recordPath(heads, n, dest)
        continue
      }
      visit(n)
    }
  }
  visit(dest)
}

func floodFill2(grid [][]byte, heads map[coord]int, dest coord) {
  seen := [10]map[coord]int{}
  for i := 0; i<=9; i++ {
    seen[i] = map[coord]int{}
  }
  seen[9][dest] = 1
  
  var visit func(coord)
  visit = func(p coord) {
    height := grid[p.x][p.y]
    paths := seen[height][p]
    
    for _, n := range next(grid, p) {
      seen[height-1][n] += paths 
    }
  }
  
  for i:=9; i>=0; i-- {
    for n := range seen[i] {
      visit(n)
    }
  }
  
  for coord, rating := range seen[0] {
    heads[coord] += rating
  }
}

func next(grid [][]byte, p coord) (res []coord) {
  v := grid[p.x][p.y]
  if n := (coord{p.x-1, p.y}); isValid(grid, n) && grid[n.x][n.y] == v-1 {
    res = append(res, n)
  }
  if n := (coord{p.x+1, p.y}); isValid(grid, n) && grid[n.x][n.y] == v-1 {
    res = append(res, n)
  }
  if n := (coord{p.x, p.y-1}); isValid(grid, n) && grid[n.x][n.y] == v-1 {
    res = append(res, n)
  }
  if n := (coord{p.x, p.y+1}); isValid(grid, n) && grid[n.x][n.y] == v-1 {
    res = append(res, n)
  }
  return res
}

func recordPath(heads map[coord][]coord, head, dest coord) {
  old := heads[head]
  if slices.Contains(old, dest) {
    return
  }
  heads[head] = append(old, dest)
}

func isValid(grid [][]byte, p coord) bool {
  return p.x >= 0 && p.y >= 0 && p.x < len(grid) && p.y < len(grid[p.x])
}
