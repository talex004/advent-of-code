package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "bytes"
)

const (
  wall = byte('#')
  box = byte('O')
  robot = byte('@')
  empty = byte('.')
  
  up = byte('^')
  down = byte('v')
  left = byte('<')
  right = byte('>')
)

type coord struct {
  i, j int
}

var directions = map[byte]coord{
  up: coord{-1, 0},
  down: coord{1, 0},
  left: coord{0, -1},
  right: coord{0, 1},
}

var world [][]byte

func main() {
  moves, pos := readInput("input")
  fmt.Println(pos)
  for _, mv := range moves {
    pos = makeMove(pos, mv)
    fmt.Println(string(mv), pos)
  }
  dumpWorld()
  fmt.Println(gps())
}

func readInput(fn string) (moves []byte, pos coord) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	blocks := bytes.Split(data, []byte("\n\n"))
	for i, ln := range bytes.Split(blocks[0], []byte("\n")) {
	  if len(ln) == 0 {
	    continue
	  }
	  world = append(world, ln)
	  if j := bytes.IndexByte(ln, robot); j != -1 {
	    pos = coord{i, j}
	  }
	}
	moves = bytes.ReplaceAll(blocks[1], []byte("\n"), []byte{})
	return moves, pos
}

func makeMove(p0 coord, mv byte) coord {
  space, found := findEmptySpace(p0, mv)
  if !found {
    return p0
  }
  return pushBoxes(space, p0, mv)
}

func findEmptySpace(p0 coord, mv byte) (coord, bool) {
  for {
    p0 = addCoord(p0, directions[mv])
    switch world[p0.i][p0.j] {
      case empty:
        return p0, true
      case wall:
        return coord{}, false
    }
  }
}

func pushBoxes(dest, src coord, mv byte) coord {
  p0 := dest
  dir := directions[mv]
  for p0 != src {
    p1 := subCoord(p0, dir)
    world[p0.i][p0.j] = world[p1.i][p1.j]
    p0 = p1
  }
  world[p0.i][p0.j] = empty
  return addCoord(p0, dir)
}

func addCoord(c0, c1 coord) (res coord) {
  res.i = c0.i + c1.i
  res.j = c0.j + c1.j
  return res
}

func subCoord(c0, c1 coord) (res coord) {
  res.i = c0.i - c1.i
  res.j = c0.j - c1.j
  return res
}

func dumpWorld() {
  for _, ln := range world {
    fmt.Println(string(ln))
  }
}

func gps() (res int) {
  for i, ln := range world {
    for j, ch := range ln {
      if ch == box {
        res += 100*i + j
      }
    }
  }
  return res
}
