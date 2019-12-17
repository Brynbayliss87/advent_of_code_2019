package main

import (
  "strconv"
  "fmt"
  "strings"
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

  for _, i := range string_values {
    j := 0
    for j < len(i) - 1 {
      if i[j] == i[j+1] {
        double_ints = append(double_ints, i)
        break
      }
      j++
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

  substrings := []string{ "111", "222", "333", "444", "555", "666", "777",
    "888", "999", "000"}


  count := len(possibles)

  for _, possible := range possibles {
    for _, substr := range substrings {
      if strings.Contains(possible, substr) {
        count--
      }
    }
  }

  fmt.Println("part 2:", count)

}
