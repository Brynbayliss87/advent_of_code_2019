package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
  "strings"
)

func main() {
  file, err := os.Open("input.txt")

  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)

  var values = []string{}
  var line_1 = []string{}
  var line_2 = []string{}

  for scanner.Scan() {
    line := scanner.Text()
    values = append(values, line)
  }

  line_1 = append(line_1, values[0])
  line_2 = append(line_2, values[1])

  line_1 = strings.Split(line_1[0], ",")
  line_2 = strings.Split(line_2[0], ",")


}


func create_vectors(directions []string)  [][]string {
  var vectors = []string{}
  for i := range directions {
    if i[0] == "L" || "R" {
      vector := ["0", i[

