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

type robot struct {
  p, v coord
}

var robots []robot
var grid coord

func main() {
  //robots, grid = readInput("input-sample"), coord{11, 7}
  robots, grid = readInput("input"), coord{101, 103}
  //part1()
  part2()
}

func readInput(fn string) (res []robot) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	for _, ln := range lines {
	  if ln == "" {
	    continue
	  }
	  var p, v coord
	  fmt.Sscanf(ln, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)
	  res = append(res, robot{p, v})
	}
	return res
}

func part1() {
  moveAll(100)
  fmt.Println(safetyFactor())
}

func part2() {
  for i:=0; i<50+10403; i++{
    moveAll(1)
    //b1, b2 := boundingBox()
    //sz := boxSize(b1, b2)
    if i % 103 == 32 {
      printState(i+1)
    }
    //if sz.x < 70 && sz.y < 70 {
    //  fmt.Println(i, sz)
    //} 
  }
}

func moveAll(moves int) {
  for i := range robots {
    move(&robots[i], moves)
  }
}

func move(r *robot, moves int) {
  vx, vy := r.v.x, r.v.y
  if vx < 0 {
    vx += grid.x
  }
  if vy < 0 {
    vy += grid.y
  }
  r.p.x = (r.p.x + moves * vx) % grid.x
  r.p.y = (r.p.y + moves * vy) % grid.y
}

func safetyFactor() int {
  var q1, q2, q3, q4 int
  hx, hy := grid.x/2, grid.y/2
  for _, r := range robots {
    switch {
      case r.p.x < hx && r.p.y < hy:
        q1++
      case r.p.x > hx && r.p.y < hy:
        q2++
      case r.p.x < hx && r.p.y > hy:
        q3++
      case r.p.x > hx && r.p.y > hy:
        q4++
    }
  }
  fmt.Println(q1, q2, q3, q4)
  return q1*q2*q3*q4
}

func boundingBox() (pmin, pmax coord){
  pmin = grid
  for _, r := range robots {
    px, py := r.p.x, r.p.y 
    if px < pmin.x {
      pmin.x = px
    }
    if py < pmin.y {
      pmin.y = py
    }
    if px > pmax.x {
      pmax.x = px
    }
    if py > pmax.y {
      pmax.y = py
    }
  }
  return pmin, pmax
}

func boxSize(p1, p2 coord) coord {
  return coord{p2.x-p1.x, p2.y-p1.y}
}

func printState(iter int) {
  pic := make([][]byte, grid.y)
  for j := range grid.y {
    pic[j] = bytes.Repeat([]byte{' '}, grid.x)
  }
  for i := range grid.y {
    for j := range grid.x {
      c := countRobots(i, j)
      if c != 0 {
        pic[i][j] = byte('0' + c)
      }
    }
  }
  fmt.Printf("==============iter %d==================\n", iter)
  for _, ln := range pic {
    fmt.Println(string(ln))
  }
}

func countRobots(i, j int) (res int) {
  for _, r := range robots {
    if r.p.x == i && r.p.y == j {
      res++
    }
  }
  return res
}
