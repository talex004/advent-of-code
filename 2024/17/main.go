package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "strconv"
  "bytes"
)

type vm struct {
  prog []byte
  a, b, c int
  ip int
  outputs []byte
}

func main() {
  vm := readInput("input")
  fmt.Println("program", vm.prog)
  fmt.Println(findRegA(vm))
}

func readInput(fn string) *vm {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	vm := &vm{}
	var prog string
	fmt.Sscanf(string(data), "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s\n", &vm.a, &vm.b, &vm.c, &prog)
	ops := strings.Split(prog, ",")
	for _, op := range ops {
	  opi, _ := strconv.Atoi(op)
	  vm.prog = append(vm.prog, byte(opi))
	} 
	return vm
}

func findRegA(o *vm) int {
  //var factors [16]int
  
  check := func(ff []int, depth int) bool {
    v := &vm{
      prog: o.prog,
      a: calc8(ff),
      b: o.b,
      c: o.c,
      outputs: o.outputs[:0],
    }
    v.run()
    if depth >= len(v.outputs) {
      return false
    }
    res := bytes.Equal(v.outputs, o.prog[len(o.prog)-depth-1:])
    //fmt.Println("for ff", ff)
    //fmt.Println("comparing", v.outputs, o.prog[len(o.prog)-depth-1:])
    //fmt.Println(ff, v.outputs, res)
    if res {
    }
    return res
  }
  
  var step func([]int, int)
  step = func(factors []int, depth int) {
    if depth >= 16 {
      fmt.Println(factors, calc8(factors))
      return
    }
    factors2 := append([]int{0}, factors...)
    for i:=0; i<8; i++ {
      factors2[0] = i
      if !check(factors2, depth) {
        continue
      }
      step(factors2, depth+1)
    }
  }
  
  step([]int{}, 0)
  
  return -1
}

func calc8(ff []int) (sum int) {
  for i, f := range ff {
    sum |= f << (3*i)
  }
  return sum
}

func (v *vm) run() {
  for {
    op := v.prog[v.ip]
    var jumped bool
    switch op {
    case 0:
      v.adv()
    case 1:
      v.bxl()
    case 2:
      v.bst()
    case 3:
      jumped = v.jnz()
    case 4:
      v.bxc()
    case 5:
      v.out()
    case 6:
      v.bdv()
    case 7:
      v.cdv()
    default:
      log.Fatalf("invalid opcode %d", op)
    }
    if !jumped {
      v.ip += 2
    }
    if v.ip >= len(v.prog) - 1 {
      break
    } 
  }
}

func (v *vm) adv() {
  v.a /= 1 << v.combo()
}

func (v *vm) bxl() {
  v.b ^= v.literal()  
}

func (v *vm) bst() {
  v.b = v.combo() & 7
}

func (v *vm) jnz() bool {
  if v.a == 0 {
    return false
  }
  v.ip = v.literal()
  return true
}

func (v *vm) bxc() {
  v.b ^= v.c
}

func (v *vm) out() {
  v.outputs = append(v.outputs, byte(v.combo() & 7))
}

func (v *vm) bdv() {
  v.b = v.a / (1 << v.combo())
}

func (v *vm) cdv() {
  v.c = v.a / (1 << v.combo())
}

func (v *vm) combo() int {
  switch val := v.prog[v.ip+1]; val {
  case 0, 1, 2, 3:
    return int(val)
  case 4:
    return v.a
  case 5:
    return v.b
  case 6:
    return v.c
  default:
    log.Fatalf("Invalid combo value %d", val)
    return 0
  }
}

func (v *vm) literal() int {
  return int(v.prog[v.ip+1])
}
