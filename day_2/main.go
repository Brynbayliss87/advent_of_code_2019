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

  part_1_values := make([]int, len(int_values))
  part_2_values := make([]int, len(int_values))

  copy(part_1_values, int_values)
  copy(part_2_values, int_values)

  part_1_values[1] = 12
  part_1_values[2] = 2

  fmt.Println("day 2, part 1: ", calculate_values(part_1_values))

  part_2_values[1] = 0
  part_2_values[2] = 0

  var result []int
  target := 19690720

  index_1 := 0
  index_2 := 0

  out:
  for part_2_values[1] <= 99 {
    part_2_values[2] = 0
    index_2 = 0
    for part_2_values[2] <= 99 {
      result = calculate_values(part_2_values)

      if result[0] == target {
        fmt.Println("result:", 100 * part_2_values[1] + part_2_values[2])
        break out
      }
      index_2++
      part_2_values = make([]int, len(int_values))
      copy(part_2_values, int_values)
      part_2_values[2] = index_2
      part_2_values[1] = index_1
    }

    result = calculate_values(part_2_values)
    if result[0] == target {
      fmt.Println("result:", 100 * part_2_values[1] + part_2_values[2])
      break
    }
    index_1++
    part_2_values = make([]int, len(int_values))
    copy(part_2_values, int_values)
    part_2_values[1] = index_1
  }
}

func calculate_values(int_values []int) []int {
  position := 0
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
      return int_values
    } else {
      position += 4
    }
  }
  return int_values
}
