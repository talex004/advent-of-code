package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "bytes"
)

const (
  Wall = '#'
  Space = '.'
  Start = 'S'
  End = 'E'
)

type coord struct {
  x, y int
}

type solver struct {
  world [][]byte
  start, end coord
  flows map[coord]int
}

func main() {
  sv := readInput(false)
  
  fmt.Println(sv.flow())
  fmt.Println(sv.cheats2(20, 100))
}

func readInput(sample bool) (res solver) {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	res.world = bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	for i, ln := range res.world {
	  if j := bytes.IndexByte(ln, Start); j != -1 {
	    res.start = coord{i, j}
	  }
	  if j := bytes.IndexByte(ln, End); j != -1 {
	    res.end = coord{i, j}
	  }
	}
	return res
}

func (c *coord) neighbors() []coord {
  return []coord{{c.x+1, c.y}, {c.x-1, c.y}, {c.x, c.y+1}, {c.x, c.y-1}}
}

func (c *coord) neighbors2() []coord {
  return []coord{
    {c.x+2, c.y},
    {c.x-2, c.y},
    {c.x, c.y+2},
    {c.x, c.y-2},
    {c.x+1, c.y-1},
    {c.x+1, c.y+1},
    {c.x-1, c.y-1},
    {c.x-1, c.y+1},
  }
}

func dumpGrid(grid [][]byte) {
  for _, row := range grid {
    fmt.Println(string(row))
  }
}

func (sv *solver) valid(c coord) bool {
  return c.x >= 0 && c.y >= 0 && c.x < len(sv.world) && c.y < len(sv.world[c.x])
}

func (sv *solver) flow() int {
  q := map[coord]struct{}{}
  cost := map[coord]int{}
  pos := sv.start
  q[pos] = struct{}{}
  cost[pos] = 0
  for len(q) > 0 {
    for c := range q {
      for _, c2 := range c.neighbors() {
        if sv.world[c2.x][c2.y] == '#' {
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
  sv.flows = cost
  return cost[sv.end]
}

func (sv *solver) cheats() (res int) {
  all := map[int]int{}
  for p1, cost1 := range sv.flows {
    for _, p2 := range p1.neighbors2() {
      cost2, ok := sv.flows[p2]
      if !ok {
        continue
      }
      saved := cost2-cost1-2
      if saved <= 0 {
        continue
      }
      all[saved] += 1
      if saved >= 100 {
        res++
      }
    }
  }
  //fmt.Println(all)
  return res
}

func (sv *solver) cheats2(max, limit int) (res int) {
  all := map[int]int{}
  for p1, cost1 := range sv.flows {
    for p2, cost2 := range sv.flows {
      path := abs(p2.x-p1.x)+abs(p2.y-p1.y)
      if path > max {
        continue
      }
      saved := cost2-cost1-path
      if saved < limit {
        continue
      }
      all[saved] += 1
      res++
    }
  }
  //fmt.Println(all)
  return res
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}
