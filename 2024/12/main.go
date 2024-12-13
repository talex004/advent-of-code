package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "maps"
  "slices"
)

func main() {
  grid = readInput("input")
  fmt.Println(part1())
  fmt.Println(part2())
}

type coord struct {
  x, y int
}

type region struct {
  plots []coord
  crop byte
}

type uniqueRegion struct {
  first coord
  crop byte
}

type edge struct {
  p1, p2 coord
}

func (e *edge) dir() rune {
  if e.p1.x == e.p2.x {
    return 'h'
  }
  if e.p1.y == e.p2.y {
    return 'v'
  }
  log.Fatalf("invalid edge %v", *e)
  return ' '
}

func (r *region) area() int {
  return len(r.plots)
}

func (r *region) perim() (sum int) {
  plots := map[coord]struct{}{}
  for _, plot := range r.plots {
    plots[plot] = struct{}{}
  }
  for _, p := range r.plots {
    crt := 4
    if isNeighbor(plots, coord{p.x+1, p.y}) { crt-- }
    if isNeighbor(plots, coord{p.x-1, p.y}) { crt-- }
    if isNeighbor(plots, coord{p.x, p.y+1}) { crt-- }
    if isNeighbor(plots, coord{p.x, p.y-1}) { crt-- }
    sum += crt
  }
  return sum
}

func (r *region) sides() int {
  var dbg = func(string, ...any) (int, error) { return 0, nil }
  if r.crop == 'J' {
    //dbg = fmt.Printf
  }
  plots := map[coord]struct{}{}
  for _, plot := range r.plots {
    plots[plot] = struct{}{}
  }
  edges := map[edge]struct{}{} 
  addEdge := func(x1, y1, x2, y2 int, dbgs string) {
    e := edge{coord{x1, y1}, coord{x2, y2}}
    edges[e] = struct{}{}
    dbg("add edge %s %v %v\n", dbgs, e, e.dir())
  }
  for _, p := range r.plots {
    dbg("for plot %v\n", p)
    x, y := p.x, p.y
    if !isNeighbor(plots, coord{x-1, y}) { addEdge(x, y, x, y+1, "up") }
    if !isNeighbor(plots, coord{x+1, y}) { addEdge(x+1, y, x+1, y+1, "down") } 
    if !isNeighbor(plots, coord{x, y+1}) { addEdge(x, y+1, x+1, y+1, "right") } 
    if !isNeighbor(plots, coord{x, y-1}) { addEdge(x, y, x+1, y, "left") }
  }
  
  numEdges := map[coord]int{}
  for ed := range edges {
    numEdges[ed.p1]++
    numEdges[ed.p2]++
  }

LOOPING:
  for {
    for edi := range edges {
      for edj := range edges {
        if edi == edj {
          continue
        }
        if adjacent(edi, edj) && numEdges[common(edi, edj)] == 2 {
          edges[merge(edi, edj)] = struct{}{}
          delete(edges, edi)
          delete(edges, edj)
          continue LOOPING
        }
      }
    }
    break
  }
  return len(edges)
}

func (r *region) String() string {
  return fmt.Sprintf("{%c:area %d, perim %d}", r.crop, r.area(), r.perim())
}

func (r uniqueRegion) String() string {
  return fmt.Sprintf("{%c:%v}", r.crop, r.first)
}

var grid []string

func readInput(fn string) (res []string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func part1() (cost int) {
  regions := map[coord]*region{}
  for i, row := range grid {
    for j, _ := range row {
      fillRegion(regions, coord{i, j}, nil)
    }
  }
  regionList := unique(regions)
  for _, r := range regionList {
    rg := regions[r.first]
    cost += rg.area() * rg.perim()
  }
  return cost
}

func part2() (cost int) {
  regions := map[coord]*region{}
  for i, row := range grid {
    for j, _ := range row {
      fillRegion(regions, coord{i, j}, nil)
    }
  }
  regionList := unique(regions)
  for _, r := range regionList {
    rg := regions[r.first]
    sd := rg.sides()
    cost += rg.area() * sd
    //fmt.Printf("Region %v, area: %v, sides: %v\n", rg, rg.area(), sd)
  }
  return cost
}

func fillRegion(regions map[coord]*region, c coord, r *region) {
  if !isValidCoord(c) {
    return
  }
  if _, ok := regions[c]; ok {
    return
  }
  if r == nil {
    r = &region{
      plots:[]coord{c},
      crop: grid[c.x][c.y],
    }
  } else {
    if grid[c.x][c.y] != r.crop {
      return
    }
    r.plots = append(r.plots, c)
  }
  regions[c] = r
  fillRegion(regions, coord{c.x-1, c.y}, r) 
  fillRegion(regions, coord{c.x+1, c.y}, r)
  fillRegion(regions, coord{c.x, c.y-1}, r)
  fillRegion(regions, coord{c.x, c.y+1}, r)
}

func isValidCoord(c coord) bool {
  return c.x >= 0 && c.y >= 0 && c.x < len(grid) && c.y < len(grid[c.x])
}

func isNeighbor(plots map[coord]struct{}, c coord) bool {
  if !isValidCoord(c) {
    return false
  }
  _, ok := plots[c]
  return ok
}

func unique(in map[coord]*region) (res []uniqueRegion) {
  set := map[uniqueRegion]struct{}{}
  for _, r := range in {
    unq := uniqueRegion{r.plots[0], r.crop}
    set[unq] = struct{}{}
  }
  return slices.Collect(maps.Keys(set))
}

func adjacent(ed1, ed2 edge) bool {
  if ed1.p1 != ed2.p2 && ed1.p2 != ed2.p1 {
    return false
  }
  if ed1.dir() != ed2.dir() {
    return false
  }
  return true // maybe?
}

func common(ed1, ed2 edge) coord {
  if ed1.p1 == ed2.p2 {
    return ed1.p1
  }
  if ed2.p1 == ed1.p2 {
    return ed2.p1
  }
  return coord{-1, -1} // invalid
}

func merge(ed1, ed2 edge) edge {
  if ed1.dir() == 'h' {
    // horizontal merge
    x := ed1.p1.x
    ymin := min(ed1.p1.y, ed1.p2.y, ed2.p1.y, ed2.p2.y)
    ymax := max(ed1.p1.y, ed1.p2.y, ed2.p1.y, ed2.p2.y)
    return edge{coord{x, ymin}, coord{x, ymax}}
  }
  if ed1.dir() == 'v' {
    // vertical merge
    y := ed1.p1.y
    xmin := min(ed1.p1.x, ed1.p2.x, ed2.p1.x, ed2.p2.x)
    xmax := max(ed1.p1.x, ed1.p2.x, ed2.p1.x, ed2.p2.x)
    return edge{coord{xmin, y}, coord{xmax, y}}
  }
  log.Fatalf("invalid merge")
  return edge{}
}

