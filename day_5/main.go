package main

import (
  "os"
  "bufio"
  "log"
  "strings"
  "../intcode"
  "fmt"
)

func main.go() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  var values = []string{}

  for scanner.Scan() {
    values = strings.Split(scanner.Text(), ",")
  }

  fmt.Println(intcode.Run(1, values))

}
