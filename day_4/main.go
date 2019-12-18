package main

import (
  "strconv"
  "fmt"
)

func main() {
  start := 367479
  end := 893698
  var string_values []string
  var double_ints []string
  var possibles []string


  for i := start; i <= end; i++ {
      string_value := strconv.Itoa(i)
      string_values = append(string_values, string_value)
  }

  double := false
  for _, i := range string_values {
    if i[0] == i[1] && i[0] != i[2] {
      double = true
    } else if i[1] == i[2] && i[1] != i[0] && i[1] != i[3]{
      double = true
    } else if i[2] == i[3] && i[2] != i[1] && i[2] != i[4]{
      double = true
    } else if i[3] == i[4] && i[3] != i[2] && i[3] != i[5]{
      double = true
    } else if i[4] == i[5] && i[4] != i[3] {
      double = true
    } else {
      double = false
    }
      if double {
        double_ints = append(double_ints, i)
      }
    }

    for _, i := range double_ints {
      j := 0
      count := 0
      for j < len(i) - 1 {
        if i[j+1] >= i[j] {
          count++
        }
      j++
      if count == 5 {
        possibles = append(possibles, i)
      }
    }
  }

  fmt.Println("part 1:", len(possibles))
}
