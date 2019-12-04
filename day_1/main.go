package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "fmt"
)

func main() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  var count int;
  var count_with_fuel int;

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    count += calculate_fuel(scanner.Text())
    count_with_fuel += recursively_calculate_fuel(scanner.Text())
  }

  fmt.Println("day 1, part 1:", count)
  fmt.Println("day 1, part 2:", count_with_fuel)
}

func calculate_fuel(string_number string) int {
  number, err := strconv.Atoi(string_number)

  if err != nil {
    log.Fatal(err)
  }

  return number / 3 - 2
}

func recursively_calculate_fuel(string_number string) int {
  fuel := calculate_fuel(string_number)
  total := fuel

  if fuel >= 0 {
    total += recursively_calculate_fuel(strconv.Itoa(fuel))
    return total
  } else {
    return 0
  }
}

