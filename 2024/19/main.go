package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
)

var maxTowelLen int
var mapPossible = map[string]int{} 

func main() {
  towels, designs := readInput(false)
  fmt.Println(part2(towels, designs))
}

func readInput(sample bool) (map[string]struct{}, []string) {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
  blocks := strings.Split(string(data), "\n\n")
  towels := strings.Split(blocks[0], ", ")
  designs := strings.Split(blocks[1], "\n")
  
  towelsMap := map[string]struct{}{}
  for _, t := range towels {
    towelsMap[t] = struct{}{}
    if len(t) > maxTowelLen {
      maxTowelLen = len(t)
    }
  }
  return towelsMap, designs
}

func part2(towels map[string]struct{}, designs []string) (res int) {
  for i, design := range designs {
    if design == "" {
      continue
    }
    res += numPossible(towels, design)
    fmt.Println(i+1, "out of", len(designs))
  }
  return res
}

func numPossible(towels map[string]struct{}, design string) (res int) {
  if len(design) == 0 {
    return 1
  }
  if v, ok := mapPossible[design]; ok {
    return v
  }

  for i := 1; i<=maxTowelLen; i++ {
    if i > len(design) {
      break
    }
    cdt := design[:i]
    if _, ok := towels[cdt]; !ok {
      continue
    }
    res += numPossible(towels, design[i:])
  }
  mapPossible[design] = res
  return res
}


