package main

import (
  "os"
  "bufio"
  "log"
  "strings"
  "advent_of_code_2019/intcode"
  "fmt"
)

func main.go() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  var values = []int{}
  var intvalues = []int{}

  for scanner.Scan() {
    values = strings.Split(scanner.Text(), ",")
  }

  for i := range len(values) {
    append(intvalues, strconv.Atoi(values[i])
  }

  fmt.Println(intcode.Run(1, values))

}
