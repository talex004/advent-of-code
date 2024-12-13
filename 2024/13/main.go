package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "math"
)

type coord struct {
  x, y int
}

type game struct {
  diffA, diffB coord
  prize coord
}

func main() {
  games := readInput("input")
  fmt.Println(solveSlow(games))
  adjustPart2(games)
  fmt.Println(solveMathing(games))
}

func readInput(fn string) (res []game) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	var crt game
	for _, ln := range lines {
	  switch {
	  case strings.HasPrefix(ln, "Button A:"):
	    fmt.Sscanf(ln, "Button A: X+%d, Y+%d", &crt.diffA.x, &crt.diffA.y)
	  case strings.HasPrefix(ln, "Button B:"):
	    fmt.Sscanf(ln, "Button B: X+%d, Y+%d", &crt.diffB.x, &crt.diffB.y)
	  case strings.HasPrefix(ln, "Prize:"):
	    fmt.Sscanf(ln, "Prize: X=%d, Y=%d", &crt.prize.x, &crt.prize.y)
	  case ln == "":
	    res = append(res, crt)
	    crt = game{}
    }
	}
	return res
}

func solveSlow(games []game) (res int) {
  for _, game := range games {
    res += bestSlow(game)
  }
  return res
}

func solveMathing(games []game) (res int) {
  for _, game := range games {
    res += bestMathing(game)
  }
  return res
}

func bestSlow(game game) (res int) {
  for na := 0; na < 100; na++ {
    xa := game.diffA.x * na
    xb := game.prize.x - xa
    if xb % game.diffB.x != 0 {
      continue
    }
    nb := xb / game.diffB.x
    ya := game.diffA.y * na
    yb := game.diffB.y * nb
    if ya+yb != game.prize.y {
      continue
    }
    cost := 3*na+nb
    if res == 0 || cost < res {
      res = cost
    }
  }
  return res
}

func bestMathing(g game) int {
  t1 := float64(g.diffB.y) - float64(g.diffA.y)*float64(g.diffB.x)/float64(g.diffA.x)
  t2 := float64(g.prize.y) - float64(g.diffA.y)*float64(g.prize.x)/float64(g.diffA.x)
  Fb := t2/t1
  Fa := (float64(g.prize.x) - float64(g.diffB.x)*Fb) / float64(g.diffA.x)
  if Fa < 0 || Fb < 0 {
    return 0
  }
  Na, Nb := int(math.Round(Fa)), int(math.Round(Fb))
  if !check(g, Na, Nb) {
    return 0
  }
  return 3*Na+Nb
}

func check(g game, Na, Nb int) bool {
  if g.diffA.x * Na + g.diffB.x * Nb != g.prize.x {
    return false
  }
  if g.diffA.y * Na + g.diffB.y * Nb != g.prize.y {
    return false
  }
  return true
}

func adjustPart2(games []game) {
  for i := range games {
    games[i].prize.x += 10000000000000
    games[i].prize.y += 10000000000000
  }
}

