package main

import (
  "os"
  "bufio"
  "log"
  "strings"
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
}
