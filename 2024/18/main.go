package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "bytes"
)

type coord struct {
  x, y int
}

type maze struct {
  corrupt []coord
  size coord
  first int
}

func main() {
  mz := readInput(false)
  grid := mz.grid()
  fmt.Println(flow(mz, grid))
  for {
    mz.first++
    mz.updateGrid(grid)
    if flow(mz, grid) == 0 {
      fmt.Println("%v", mz.corrupt[mz.first-1])
      break
    }
  }
}

func readInput(sample bool) *maze {
  sz, fn, first := coord{71, 71}, "input", 1024
  if sample {
    sz, fn, first = coord{7, 7}, "input-sample", 12
  }
  data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	res := &maze{size: sz, first: first}
	for _, ln := range strings.Split(string(data), "\n") {
	  c := coord{}
	  if n, _ := fmt.Sscanf(ln, "%d,%d", &c.x, &c.y); n == 2 {
	    res.corrupt = append(res.corrupt, c)
	  }
	}
	return res
}

func (mz *maze) grid() (res [][]byte) {
  res = make([][]byte, mz.size.x)
  for i := range mz.size.x {
    res[i] = bytes.Repeat([]byte{'.'}, mz.size.y)
  }
  mz.updateGrid(res)
  return res
}

func (mz *maze) updateGrid(grid [][]byte) {
  for i := range mz.first {
    if i >= len(mz.corrupt) {
      break
    }
    c := mz.corrupt[i]
    grid[c.x][c.y] = '#'
  }
}

func (mz *maze) valid(c coord) bool {
  return c.x >= 0 && c.y >= 0 && c.x < mz.size.x && c.y < mz.size.y
}

func (c *coord) neighbors() []coord {
  return []coord{{c.x+1, c.y}, {c.x-1, c.y}, {c.x, c.y+1}, {c.x, c.y-1}}
}

func dumpGrid(grid [][]byte) {
  for _, row := range grid {
    fmt.Println(string(row))
  }
}

func flow(mz *maze, grid [][]byte) int {
  q := map[coord]struct{}{}
  cost := map[coord]int{}
  pos := coord{}
  q[pos] = struct{}{}
  cost[pos] = 0
  for len(q) > 0 {
    for c := range q {
      for _, c2 := range c.neighbors() {
        if !mz.valid(c2) || grid[c2.x][c2.y] == '#' {
          continue
        }
        if v, ok := cost[c2]; !ok || cost[c]+1 < v {
          cost[c2] = cost[c]+1
          q[c2] = struct{}{}
        }
      }
      delete(q, c)
    }
  }
  return cost[coord{mz.size.x-1, mz.size.y-1}]
}
