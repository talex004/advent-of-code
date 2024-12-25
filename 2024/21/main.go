package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "strconv"
)

func main() {
  sum := 0
  for _, code := range readInput(false) {
    if code == "" {
      continue
    }
    sum += numCode(code)*codeLengthNum(code)
  }
  fmt.Println(sum)
}

func readInput(sample bool) []string {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
  if err != nil {
	  log.Fatalf("Reading input: %v", err)
  }
  return strings.Split(string(data), "\n")
}

type coord struct {
  x, y int
}

var numKeys = []string{"789", "456", "123", "#0A"}
var dirKeys = []string{"#^A", "<v>"}

var numKeyLookup = map[rune]coord{
  '7': {0, 0},
  '8': {0, 1},
  '9': {0, 2},
  '4': {1, 0},
  '5': {1, 1},
  '6': {1, 2},
  '1': {2, 0},
  '2': {2, 1},
  '3': {2, 2},
  '#': {3, 0},
  '0': {3, 1},
  'A': {3, 2},
}

var dirKeyLookup = map[rune]coord{
  '#': {0, 0},
  '^': {0, 1},
  'A': {0, 2},
  '<': {1, 0},
  'v': {1, 1},
  '>': {1, 2},
}

type codeKey struct {
  code string
  level int
}

var codeLenMap = map[codeKey]int{}

func codeLength(code string, level int) (res int) {
  if level == 0 {
    return len(code)
  }
  ck := codeKey{code, level}
  if v, ok := codeLenMap[ck]; ok {
    return v
  }
  defer func() {
    codeLenMap[ck] = res
  }()
  
  lastPos := dirKeyLookup['A']
  
  for _, press := range code {
    pos := dirKeyLookup[press]
    diff := coord{pos.x-lastPos.x, pos.y-lastPos.y}
    count := 0

    //winCode := ""
    for _, nextCode := range keyPresses(diff) {
      if !isValid(lastPos, nextCode, dirKeyLookup['#']) {
        continue
      }
      if v := codeLength(nextCode, level-1); count == 0 || v < count {
        count = v
        //winCode = nextCode
      }
    }
    //fmt.Println(string([]rune{press}), "\t", winCode, "\t", count)
    if count == 0 {
      fmt.Println("Nothing valid for diff", diff, "from", lastPos, "to", pos, "for", string(press))
    }
    res += count
    lastPos = pos
  }
  return res
}

func codeLengthNum(code string) (res int) {  
  lastPos := numKeyLookup['A']
  
  for _, press := range code {
    pos := numKeyLookup[press]
    diff := coord{pos.x-lastPos.x, pos.y-lastPos.y}
    count := 0
    //winCode := ""
    for _, nextCode := range keyPresses(diff) {
      if !isValid(lastPos, nextCode, numKeyLookup['#']) {
        continue
      }
      if v := codeLength(nextCode, 25 /*level here*/); count == 0 || v < count {
        count = v
        //winCode = nextCode
      }
    }
    //fmt.Println(string([]rune{press}), "\t", winCode, "\t", count)
    if count == 0 {
      fmt.Println("Nothing valid for diff", diff, "from", lastPos, "to", pos, "for", string(press))
      fmt.Println(keyPresses(diff))
    }
    res += count
    lastPos = pos
  }
  return res
}

func keyPresses(diff coord) (res []string) {
  var step func(string, coord)
  step = func(code string, diff coord) {
    if diff == (coord{}) {
      res = append(res, code+"A")
      return
    }
    
    // try vertical
    if diff.x > 0 {
      step(code+"v", coord{diff.x-1, diff.y})
    } else if diff.x < 0 {
      step(code+"^", coord{diff.x+1, diff.y})
    }
    // try horizontal
    if diff.y > 0 {
      step(code+">", coord{diff.x, diff.y-1})
    } else if diff.y < 0 {
      step(code+"<", coord{diff.x, diff.y+1})
    }
  }
  
  step("", diff)
  return res
}

func isValid(pos coord, code string, bad coord) bool {
  for _, ch := range code {
    switch ch {
      case '^':
        pos.x--
      case 'v':
        pos.x++
      case '>':
        pos.y++
      case '<':
        pos.y--
    }
    if pos == bad {
      return false
    }
  }
  return true
}

func numCode(c string) int {
  a, _ := strconv.Atoi(c[:len(c)-1])
  return a
} 

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

