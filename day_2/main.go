package main

import (
  "bufio"
  "os"
  "log"
  "strconv"
  "strings"
  "fmt"
)

func main() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  var values = []string{}

  for scanner.Scan() {
    values = strings.Split(scanner.Text(), ",")
  }

  var int_values = []int{}

  for _, i := range values {
      integer, err := strconv.Atoi(i)
      if err != nil {
          panic(err)
      }
      int_values = append(int_values, integer)
  }

  position := 0

  int_values[1] = 12
  int_values[2] = 2

  for position <= len(int_values) {
    if int_values[position] == 1 {
      result := int_values[int_values[position + 1]] + int_values[int_values[position + 2]]
      int_values[int_values[position + 3]] = result
      position += 4
    } else if int_values[position] == 2 {
      result := int_values[int_values[position + 1]] * int_values[int_values[position + 2]]
      int_values[int_values[position + 3]] = result
      position += 4
    } else if int_values[position] == 99 {
      fmt.Println("day 2, part 1: ", int_values)
      break
    } else {
      position += 4
    }
  }
  fmt.Println("day 2, part 1: ", int_values)
}
