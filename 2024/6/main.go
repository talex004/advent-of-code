package main

import (
  "io/ioutil"
  "log"
  "bytes"
  "fmt"
)

func main() {
  maze := readInput()
  println(solve2(maze))
}

func readInput() [][]byte {
  data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := bytes.Split(data, []byte{'\n'})
	return lines
}

func solve1(maze [][]byte) {
  x, y := findStart(maze)
  if x == -1 || y == -1 {
    return
  }
  println(x, y)
  dir := 0
  
  for {
      maze[x][y] = 'x'
      x2, y2 := x, y
      switch(dir) {
      case 0: // up
        x2--
      case 1: // right
        y2++
      case 2: // down
        x2++
      case 3: // left
        y2--
      }
      
      if x2 < 0 || x2 >= len(maze) || y2 < 0 || y2 >= len(maze[x2]) {
        break // done
      }
      
      next := maze[x2][y2]
      if next == '.' || next == 'x' {
        // keep moving
        x, y = x2, y2
      } else if next == '#' {
        // turn, will move there next round
        dir = (dir+1)%4
      } else {
        log.Fatalf("unexpected %c", next)
      }
  }
}

func findStart(maze [][]byte) (x, y int) {
  x, y = -1, -1
  for i, ln := range maze {
    if j := bytes.Index(ln, []byte{'^'}); j != -1 {
      x, y = i, j
      break
    }
  }
  return x, y
}

type step struct {
  x, y int
  dir int // 0 up, 1, right, 2 down, 3 left 
}

func solve2(maze [][]byte) (res int) {
  x, y := findStart(maze)
  for i, ln := range maze {
    for j, old := range ln {
      if i == x && j == y || old == byte('#') {
        continue
      }
      maze[i][j] = byte('#')
      if hasLoop(maze, x, y) {
        res++
      }
      maze[i][j] = old
    }
  }
  return res
}

func hasLoop(maze [][]byte, x, y int) bool {
  path := []step{{x, y, 0}}
  
  for {
    next, ok := next(maze, path[len(path)-1])
    if !ok {
      return false
    }
    if already(path, next) {
      return true
    }
    path = append(path, next)
  }
}

func next(maze [][]byte, s step) (step, bool) {
  dir := s.dir
  for {
    n := s
    n.dir = dir
    switch n.dir {
    case 0:
      n.x--
    case 1:
      n.y++
    case 2:
      n.x++
    case 3:
      n.y--
    }
    
    if n.x < 0 || n.x >= len(maze) || n.y < 0 || n.y >= len(maze[n.x]) {
      return n, false
    }
    if maze[n.x][n.y] != byte('#') {
      return n, true
    }
    dir = (dir+1) % 4
  }
}

func already(path []step, next step) bool {
  for _, s := range path {
    if s.x == next.x && s.y == next.y && s.dir == next.dir {
      return true
    }
  }
  return false
}
