package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "math/rand"
  "slices"
)

type gate struct {
  gateType string
  in1, in2 string
  out string
}

type logic struct {
  states map[string]int
  gates []gate
}

func main() {
  l := readInput(false)
  orig := l.check()
  
  swaps := stage1(l, orig)
  fmt.Println(stage2(l, swaps))
}

func stage1(l logic, orig int) (swaps [][2]string) {
  var allGates []string
  for _, g := range l.gates {
    allGates = append(allGates, g.out)
  }
  
  for _, g1 := range allGates {
    for _, g2 := range allGates {
      if g2 <= g1 {
        continue
      }
      l.swap(g1, g2)
      if l.check() < orig {
        swaps = append(swaps, [2]string{g1, g2})
      }
      l.swap(g1, g2)
    }
  }
  return swaps
}

func stage2(l logic, swaps [][2]string) string {
  for i1, s1 := range swaps {
    for i2, s2 := range swaps {
      if i2 >= i1 {
        break
      }
      if conflicts(s1, s2) {
        continue
      }
      
      for i3, s3 := range swaps {
        if i3 >= i2 {
          break
        }
        if conflicts(s1, s2, s3) {
          continue
        }
        for i4, s4 := range swaps {
          if i4 >= i3 {
            break
          }
          if conflicts(s1, s2, s3, s4) {
            continue
          }
          // here
          l.swap(s1[0], s1[1])
          l.swap(s2[0], s2[1])
          l.swap(s3[0], s3[1])
          l.swap(s4[0], s4[1])
          if l.check() == 0 {
            if l.checkRandom() {
              res := []string{s1[0], s1[1], s2[0], s2[1], s3[0], s3[1], s4[0], s4[1]}
              slices.Sort(res)
              return strings.Join(res, ",")
            }
          }
          l.swap(s1[0], s1[1])
          l.swap(s2[0], s2[1])
          l.swap(s3[0], s3[1])
          l.swap(s4[0], s4[1])
          // end
        }
      }
    }
  }
  return ""
}

func readInput(sample bool) logic {
  res := logic{states: map[string]int{}}
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
  if err != nil {
	  log.Fatalf("Reading input: %v", err)
  }
  parts := strings.Split(string(data), "\n\n")
  for _, ln := range strings.Split(parts[0], "\n") {
    var name string
    var val int
    if n, _ := fmt.Sscanf(ln, "%s %d", &name, &val); n > 0 {
      name = strings.TrimSuffix(name, ":")
      res.states[name] = val
    }
  }
  
  for _, ln := range strings.Split(parts[1], "\n") {
    var g gate
    if n, _ := fmt.Sscanf(ln, "%s %s %s -> %s", &g.in1, &g.gateType, &g.in2, &g.out); n > 0 {
      res.gates = append(res.gates, g)
    }
  }
  return res
}

func (l *logic) check() (res int) {
  for i := 0; i<=44; i++ {
    a, b := 0, 1<<i
    l.set(a, b)
    l.sim()
    if l.output() != (a+b) {
      res++
    }
    
    a, b = 1<<i, 0
    l.set(a, b)
    l.sim()
    if l.output() != (a+b) {
      res++
    }
    
    
    a, b = 1<<i, 1<<i
    l.set(a, b)
    l.sim()
    if l.output() != (a+b) {
      res++
    }
    
    if i == 0 {
      continue
    }
    a, b = 1<<i, 1<<(i-1)
    l.set(a, b)
    l.sim()
    if l.output() != (a+b) {
      res++
    }
  }
  return res
}

func (l *logic) sim() {
  for {
    changed := false
    for _, g := range l.gates {
      in1, ok := l.states[g.in1]
      if !ok {
        continue
      }
      in2, ok := l.states[g.in2]
      if !ok {
        continue
      }
      if _, ok := l.states[g.out]; ok {
        // already computed
        continue
      }
      var out int
      switch g.gateType {
      case "AND":
        out = in1 & in2
      case "OR":
        out = in1 | in2
      case "XOR":
        out = in1 ^ in2
      default:
        log.Fatalf("Unknown gate %s", g.gateType)
      }
      l.states[g.out] = out
      changed = true
    }
    if !changed {
      return
    }
  }
}

func (l *logic) swap(o1, o2 string) {
  for i := range l.gates {
    if l.gates[i].out == o1 {
      l.gates[i].out = o2
      continue
    }
    if l.gates[i].out == o2 {
      l.gates[i].out = o1
      continue
    }
  }
}

func (l *logic) output() (res int) {
  for i:=0; i<63; i++ {
    wire := fmt.Sprintf("z%02d", i)
    b, ok := l.states[wire]
    if !ok {
      break
    }
    res |= b << i
  }
  return res
}

func (l *logic) set(x, y int) {
  clear(l.states)
  for i:=0; i<=44; i++ {
    v := x & 1
    l.states[fmt.Sprintf("x%02d", i)] = v
    x >>= 1
  }
  for i:=0; i<=44; i++ {
    v := y & 1
    l.states[fmt.Sprintf("y%02d", i)] = v
    y >>= 1
  }
}

func (l *logic) checkRandom() bool {
  for i:=0; i<100; i++ {
    a := rand.Intn(1<<44)
    b := rand.Intn(1<<44)
    l.set(a, b)
    l.sim()
    if out := l.output(); out != a+b {
      return false
    }
  }
  return true
}

func conflicts(ss... [2]string) bool {
  count := map[string]int{}
  for _, s := range ss {
    count[s[0]] = count[s[0]]+1
    count[s[1]] = count[s[1]]+1
  }
  
  for _, c := range count {
    if c > 1 {
      return true
    }
  }
  return false
}
