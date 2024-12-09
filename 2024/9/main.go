package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "slices"
  "strconv"
)

type block struct {
  fileID int
  length int
  free bool
  consumed bool
}

func (b *block) String() string {
  s := "."
  if !b.free {
    s = strconv.Itoa(b.fileID)
  }
  return strings.Repeat(s, b.length)
}

func main() {
  in := readInput()
  out := compact2(in)
  for _, bk := range out {
    fmt.Print(bk)
  }
  fmt.Println()
  fmt.Println(checksum(out))
}

func readInput() (res []*block) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	for i, ch := range strings.TrimSpace(string(data)) {
	  bk := &block{length: int(ch-'0')}
	  if i % 2 == 0 {
	    bk.fileID = i/2
	  } else {
	    bk.free = true
    }
    res = append(res, bk)
	}
	return res
}

func compact(in []*block) (out []*block) {
  queue := append([]*block{}, in...)
  slices.Reverse(queue)
  queue = slices.DeleteFunc(queue, func(b *block) bool { return b.free })
  
  for _, bk := range in {
    if !bk.free {
      out = append(out, &block{fileID: bk.fileID, length: bk.length})
      bk.consumed = true
      continue
    }
    for bk.length > 0 && len(queue) > 0 {
      if queue[0].consumed {
        break
      }
      if bk.length >= queue[0].length {
        out = append(out, &block{fileID: queue[0].fileID, length: queue[0].length})
        bk.length -= queue[0].length
        queue[0].free = true
        queue = queue[1:]
        continue
      }
      out = append(out, &block{fileID: queue[0].fileID, length: bk.length})
      queue[0].length -= bk.length
      bk.length = 0
    }
  }
  return out
}

func compact2(in []*block) (out []*block) {
  queue := append([]*block{}, in...)
  slices.Reverse(queue)
  queue = slices.DeleteFunc(queue, func(b *block) bool { return b.free })
  
  for _, bk := range queue {
    j, ok := findSpace(in, bk)
    if !ok {
      continue
    }
    in[j].length -= bk.length
    in = slices.Insert(in, j, &block{fileID: bk.fileID, length: bk.length})
    bk.free = true
  }
  return in
}

func findSpace(disk []*block, target *block) (int, bool) {
  for i, bk := range disk {
    if !bk.free && bk.fileID == target.fileID {
      return 0, false
    }
    if bk.free && bk.length >= target.length {
      return i, true
    }
  }
  return 0, false
}

func checksum(in []*block) (res int) {
  j := 0
  for _, bk := range in {
    for range bk.length {
      if !bk.free {
        res += j*bk.fileID
      }
      j++
    }
  }
  return res
}
