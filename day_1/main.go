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

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    count += calculate_fuel(scanner.Text())
  }

  fmt.Println("day 1:", count)
}

func calculate_fuel(string_number string) int {
  number, err := strconv.Atoi(string_number)

  if err != nil {
    log.Fatal(err)
  }

  return number / 3 - 2
}


