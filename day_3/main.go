package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
  "strings"
  "strconv"
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

  line_1_vectors := create_vectors(line_1)
  line_2_vectors := create_vectors(line_2)

  fmt.Println("line 1:", line_1_vectors)
  fmt.Println("line 2:", line_2_vectors)

}

type vector struct {
  x int
  y int
}

func create_vectors(directions []string)  []vector {
  var all_vectors = []vector{}
  var a_vector = vector{}
  var last_vector = vector{}
  all_vectors = append(all_vectors, last_vector)
  for _, i := range directions {
    axis := string(i[0])
    last_vector = all_vectors[len(all_vectors) -1]
    if axis == "L" || axis == "R" {
      y, err := strconv.Atoi(i[1:len(i)])
      if err != nil {
        log.Fatal("err")
      }
      a_vector = vector{x: last_vector.x, y: last_vector.y + y}
    } else {
      x, err := strconv.Atoi(i[1:len(i)])
      if err != nil {
        log.Fatal("err")
      }

      a_vector = vector{x: last_vector.x + x, y: last_vector.y}
    }
    all_vectors = append(all_vectors, a_vector)
  }
  return all_vectors
}


