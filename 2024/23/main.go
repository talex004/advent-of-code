package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "strings"
  "slices"
  "maps"
)

type edge struct {
  n1, n2 string
}

func main() {
  edges := readInput(false)
  fmt.Println(part1(edges))
  fmt.Println(part2(edges))
}

func readInput(sample bool) (res []edge) {
  fn := "input"
  if sample {
    fn = "input-sample"
  }
  data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Reading input: %v", err)
	}
	for _, ln := range strings.Split(string(data), "\n") {
	  if ln == "" {
	    continue
    }
    pair := strings.Split(ln, "-")
    res = append(res, edge{pair[0], pair[1]})
	}
	return res
}

func part1(edges []edge) (res int) {
  neighbors := neighbors(edges)
  
  nodes := slices.Collect(maps.Keys(neighbors))
  slices.Sort(nodes)
  
  for _, node := range nodes {
    if node[0] != 't' {
      continue
    }
    nbrs := neighbors[node]
    nbrsOrd := slices.Collect(maps.Keys(nbrs))
    slices.Sort(nbrsOrd)
    
    for _, nb1 := range nbrsOrd {
      if nb1[0] == 't' && nb1 <= node {
        continue
      }
      for _, nb2 := range nbrsOrd {
        if nb2 <= nb1 {
          continue
        }
        if nb2[0] == 't' && nb2 <= node {
          continue
        }
        if _, ok := neighbors[nb1][nb2]; ok {
          if nb1[0] == 't' || nb2[0] == 't' {
            //fmt.Println(node, nb1, nb2)
          }
          res++
        }
      } 
    }
  }
  return res
}

func part2(edges []edge) string {
  neighbors := neighbors(edges)
  nodes := slices.Collect(maps.Keys(neighbors))
  slices.Sort(nodes)
  
  var cluster []string 
  
  for _, node := range nodes {
    cdt := connected(node, neighbors)
    if len(cdt) > len(cluster) {
      cluster = cdt
    }
  }
  slices.Sort(cluster)
  return strings.Join(cluster, ",")
}

func connected(node string, neighbors map[string]map[string]struct{}) []string {
  res := map[string]struct{}{node: {}}
  // assume all neighbors are connected
  for nb := range neighbors[node] {
    res[nb] = struct{}{}
  }
  counts := nbrCounts(res, neighbors)
  fmt.Println(counts)
  
  for !isConnected(res, neighbors) {
    low := lowest(counts)
    delete(counts, low)
    delete(res, low)
  }
  
  return slices.Collect(maps.Keys(res))
}

func lowest(counts map[string]int) (res string) { 
  val := 0
  for node, count := range counts {
    if res == "" ||  count < val {
      res = node
      val = count
    }
  }
  return res
}

func isConnected(nodes map[string]struct{}, neighbors map[string]map[string]struct{}) bool {
  for n1 := range nodes {
    for n2 := range nodes {
      if n1 == n2 {
        continue
      }
      if _, ok := neighbors[n1][n2]; !ok {
        return false
      }
    }
  }
  return true
}

func nbrCounts(cluster map[string]struct{}, neighbors map[string]map[string]struct{}) map[string]int {
  res := map[string]int{}
  for node := range cluster {
    cnt := 0 
    for nbr := range neighbors[node] {
      if _, ok := cluster[nbr]; ok {
        cnt++
      }
    }
    res[node] = cnt
  }
  return res
}

func neighbors(edges []edge) map[string]map[string]struct{} {
  res := map[string]map[string]struct{}{}
  for _, e := range edges {
    m := res[e.n1]
    if m == nil {
      m = map[string]struct{}{}
      res[e.n1] = m
    }
    m[e.n2] = struct{}{}
    
    m = res[e.n2]
    if m == nil {
      m = map[string]struct{}{}
      res[e.n2] = m
    }
    m[e.n1] = struct{}{}
  }
  return res
}
