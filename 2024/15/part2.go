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
  
  box1 = byte('[')
  box2 = byte(']')
  
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

var isMovable = map[byte]bool{
  box1: true,
  box2: true,
  empty: true,
}

var world [][]byte

func main() {
  moves, pos := readInput("input")
  for _, mv := range moves {
    pos = makeMove(pos, mv)
  }
  dumpGrid(world)
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
	  var newln []byte
	  for _, ch := range ln {
	    var newch []byte
	    switch ch {
      case robot:
        newch = []byte{robot, empty}
      case box:
        newch = []byte{box1, box2}
      case wall, empty:
        newch = bytes.Repeat([]byte{ch}, 2)
	    }
	    newln = append(newln, newch...)
	  } 
	  world = append(world, newln)
	  if j := bytes.IndexByte(newln, robot); j != -1 {
	    pos = coord{i, j}
	  }
	}
	moves = bytes.ReplaceAll(blocks[1], []byte("\n"), []byte{})
	return moves, pos
}

func makeMove(p0 coord, mv byte) coord {
  if mv == left || mv == right {
    return makeHorizontalMove(p0, mv)
  }
  return makeVerticalMove(p0, mv)
}

func makeHorizontalMove(p0 coord, mv byte) coord {
  space, found := findEmptySpace(p0, mv)
  if !found {
    return p0
  }
  return pushBoxesH(space, p0, mv)
}

func makeVerticalMove(p0 coord, mv byte) coord {
  mask, p1, ok := pushMask(p0, mv)
  if !ok {
    return p0
  }
  pushBoxesV(mask, mv)
  return p1
}

func pushMask(p0 coord, mv byte) ([][]byte, coord, bool) {
  mask := blankMask()
  set(mask, p0, 'x')
  
  dir := directions[mv]
  p1 := p0
  for {
    found := false
    for j, m := range mask[p1.i] {
      if m != 'x' {
        continue
      }
      ppush := addCoord(coord{p1.i, j}, dir)
      pushobj := get(world, ppush)
      if !isMovable[pushobj] {
        return nil, coord{}, false
      }
      if pushobj == box1 || pushobj == box2 {
        set(mask, ppush, 'x')
        set(mask, complement(ppush, pushobj), 'x')
        found = true
      }
    }
    if !found {
      break
    }
    p1 = addCoord(p1, dir)
  }
  return mask, addCoord(p0, dir), true
}

func pushBoxesV(mask [][]byte, ch byte) {
  if ch == up {
    // copy from top
    for i := 0; i<len(world)-1; i++ {
      ln := world[i]
      for j := range ln {
        if mask[i+1][j] == 'x' {
          world[i][j] = world[i+1][j]
          world[i+1][j] = empty
        }
      } 
    }
    return
  }
  if ch == down {
    // copy from bottom
    for i := len(world)-1; i>0; i-- {
      ln := world[i]
      for j := range ln {
        if mask[i-1][j] == 'x' {
          world[i][j] = world[i-1][j]
          world[i-1][j] = empty
        }
      } 
    }
    return
  }
  log.Fatalf("unknown move %v", ch)
}

func set(what [][]byte, c coord, ch byte) {
  what[c.i][c.j] = ch
}

func get(what [][]byte, c coord) byte {
  return what[c.i][c.j]
}

func complement(p1 coord, ch byte) coord {
  if ch == box1 {
    return addCoord(p1, directions[right])
  } else if ch == box2 {
    return addCoord(p1, directions[left])
  }
  log.Fatalf("ch %v invalid", ch)
  return coord{}
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

func pushBoxesH(dest, src coord, mv byte) coord {
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

func blankMask() (res [][]byte) {
  res = make([][]byte, len(world))
  for i := range world {
    res[i] = bytes.Repeat([]byte{empty}, len(world[i]))
  }
  return res
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

func dumpGrid(grid [][]byte) {
  for _, ln := range grid {
    fmt.Println(string(ln))
  }
}

func gps() (res int) {
  for i, ln := range world {
    for j, ch := range ln {
      if ch == box1 {
        res += 100*i + j
      }
    }
  }
  return res
}
