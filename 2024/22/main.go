package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "strconv"
)

var seqs = []map[[4]int8]int8{}

func main() {
  sens := readInput(false)
  fmt.Println(part1(sens))
  fmt.Println(part2(sens))
}

func readInput(sample bool) (res []int) {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	for _, s := range strings.Split(string(data), "\n") {
	  if s == "" {
	    continue
    }
	  i, err := strconv.Atoi(s)
	  if err != nil {
	    log.Panicln("atoi", err)
	  }
	  res = append(res, i)
	  seqs = append(seqs, map[[4]int8]int8{})
	}
	return res
}

func part1(sens []int) int {
  res := 0
  for i, sen := range sens {
    v, seq := next2000(sen)
    res += v
    seqs[i] = seq
  }
  return res
}

func part2(sens []int) (res int) {
  fmt.Println(seqs[0])
  for _, seq := range seqs {
    for k := range seq {
      v := numBananas(k)
      if v > res {
        res = v
        fmt.Println(v)
      }
    }
  }
  return res
}

func numBananas(k [4]int8) (res int) {
  for _, seq := range seqs {  
    res += int(seq[k])
  }
  return res
}

func next2000(a int) (int, map[[4]int8]int8) {
  seq := map[[4]int8]int8{}
  var last, d3, d2, d1 int8
  last = int8(a % 10);
  for i:=0; i<2000; i++ {
    a = next(a)
    crt := int8(a % 10)
    d0 := crt - last
    if i >= 3 {
      k := [4]int8{d3, d2, d1, d0}
      if _, ok := seq[k]; !ok {
        seq[k] = crt
      }
    }
    last = crt
    d3, d2, d1 = d2, d1, d0
  }
  return a, seq
}

func next(i int) int {
  i = prune(mix(i, i*64))
  i = prune(mix(i, i/32))
  i = prune(mix(i, i*2048))
  return i
}

func mix(a, b int) int {
  return a ^ b
}

func prune(a int) int {
  return a % 16777216
}
