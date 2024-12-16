package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "bytes"
  "math"
)

const (
  Wall = '#'
  Space = '.'
  Start = 'S'
  End = 'E'
)

const (
  Right = 0
  Down = 1
  Left = 2
  Up = 3
)

type coord struct {
  i, j, dir int
}

var moveDescr = map[int]coord {
  Right: coord{0, 1, Right},
  Left: coord{0, -1, Left},
  Down: coord{1, 0, Down},
  Up: coord{-1, 0, Up},
}

type solver struct {
  // input
  world [][]byte
  start, end coord
  
  flow [][][4]int
  paths map[coord][]coord
}

func main() {
  s := readInput("input")
  s.init()
  fmt.Println(s.solveFlow())
  fmt.Println(s.solvePaths())
}

func readInput(fn string) (res solver) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	res.world = bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	for i, ln := range res.world {
	  if j := bytes.IndexByte(ln, Start); j != -1 {
	    res.start = coord{i, j, Right}
	  }
	  if j := bytes.IndexByte(ln, End); j != -1 {
	    res.end = coord{i, j, -1}
	  }
	}
	return res
}

func (c coord) withDir(d int) coord {
  return coord{c.i, c.j, d}
}

func (c coord) add(d coord) coord {
  return coord{c.i+d.i, c.j+d.j, c.dir}
}

func (c coord) next() map[coord]int {
  return map[coord]int{
    c.add(moveDescr[c.dir]): 1,
    c.withDir(next(c.dir)): 1000,
    c.withDir(prev(c.dir)): 1000,
    c.withDir(inv(c.dir)):  2000,
  }
}

func next(d int) int {
  return (d+1)%4
}

func prev(d int) int {
  return (d+3)%4
}

func inv(d int) int {
  return (d+2)%4
}

func (s *solver) solveFlow() int {
  q := map[coord]struct{}{}
  
  add := func(c coord) {
    q[c] = struct{}{}
  }
  add(s.end.withDir(Up))
  add(s.end.withDir(Down))
  add(s.end.withDir(Left))
  add(s.end.withDir(Right))
  
  s.flow[s.end.i][s.end.j] = [4]int{0, 0, 0, 0}
  
  for len(q) != 0 {
    for c := range q {
      delete(q, c)
      cf := s.getf(c)
      for n, cost := range c.next() {
        if s.get(n) == Wall {
          continue
        }
        nf := s.getf(n)
        if cf+cost < nf {
          s.setf(n, cf+cost)
          add(n)
          s.setPath(n, c)
          continue
        }
        if cf+cost == nf {
          s.addPath(n, c)
        }
      }
    }
  }
  
  return s.getf(s.start)
}

func (s *solver) solvePaths() int {
  spots := map[coord]struct{}{}
  var step func(coord)
  step = func(c coord) {
    spots[c] = struct{}{}
    for _, path := range s.paths[c] {
      if _, ok := spots[path]; ok {
        continue
      }
      step(path)
    }
  }
  step(s.start)
  
  spots0d := map[coord]struct{}{}
  for c := range spots {
    spots0d[c.withDir(-1)] = struct{}{}
  }
  //s.dumpWorld(spots0d)
  return len(spots0d)
}

func (s *solver) init() {
  s.flow = make([][][4]int, len(s.world))
  for i := range s.flow {
    s.flow[i] = make([][4]int, len(s.world[i]))
    for j := range s.flow[i] {
      s.flow[i][j] = [4]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}
    }
  }
  s.paths = map[coord][]coord{}
}

func (s *solver) get(c coord) byte {
  return s.world[c.i][c.j]
}

func (s *solver) getf(c coord) int {
  return s.flow[c.i][c.j][c.dir]
}

func (s *solver) setf(c coord, cost int) {
  s.flow[c.i][c.j][c.dir] = cost
}

func (s *solver) addPath(c, prev coord) {
  s.paths[c] = append(s.paths[c], prev)
}

func (s *solver) setPath(c, prev coord) {
  s.paths[c] = []coord{prev}
}

func (s *solver) dumpWorld(visited map[coord]struct{}) {
  fmt.Println("World map:")
  for i := range s.world {
    row := append([]byte{}, s.world[i]...)
    for j := range row {
      c := coord{i, j, -1}
      if _, ok := visited[c]; ok {
        row[j] = 'O'
      }
    }
    fmt.Println(string(row))
  }
}
