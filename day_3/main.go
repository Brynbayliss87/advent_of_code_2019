package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
  "strings"
  "strconv"
  "math"
  "sort"
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

  hash := make(map[vector]bool)
  var inter = []vector{}

  for _, e := range line_1_vectors {
    hash[e] = true
  }

  for _, e := range line_2_vectors {
    if hash[e] && (e.x != 0 && e.y != 0) {
      inter = append(inter, e)
     }
  }

  shortest_journeys := shortest_journey(line_1_vectors, line_2_vectors, inter)
  sort.Sort(sort.IntSlice(shortest_journeys))
  fmt.Println("part 2:", shortest_journeys[0])
  result := find_shortest(inter)
  fmt.Println("result:", result)
}

type vector struct {
  x int
  y int
}

func create_vectors(directions []string)  []vector {
  var all_vectors = []vector{}
  var last_vector = vector{}
  all_vectors = append(all_vectors, last_vector)
  for _, i := range directions {
    direction := string(i[0])
    last_vector = all_vectors[len(all_vectors) -1]
    distance, err := strconv.Atoi(i[1:len(i)])
    if err != nil {
      log.Fatal(err)
    }
    distances := full_distance(distance, direction, last_vector)
    for _, i := range distances {
      all_vectors = append(all_vectors, i)
    }
  }
  return all_vectors
}

func full_distance(distance int, direction string, last_vector vector) []vector {
  var result = []vector{}

  for i := 0; i < distance; i++ {
    if direction == "R" {
      a_vector := vector{x: last_vector.x, y: last_vector.y + 1}
        result = append(result, a_vector)
        last_vector = a_vector
    } else if direction == "L" {
      a_vector := vector{x: last_vector.x, y: last_vector.y -1}
      result = append(result, a_vector)
      last_vector = a_vector
    } else if direction == "D" {
      a_vector := vector{x: last_vector.x - 1, y: last_vector.y}
      result = append(result, a_vector)
      last_vector = a_vector
    } else if direction == "U" {
      a_vector := vector{x: last_vector.x + 1, y: last_vector.y}
      result = append(result, a_vector)
      last_vector = a_vector
    }
  }
  return result
}

func find_shortest(intersects []vector) float64 {
  var result float64
  for i, j := range intersects {
    distance := math.Abs(float64(j.x)) + math.Abs(float64(j.y))
    if i == 0 {
      result = distance
    } else if distance < result {
      result = distance
    }
  }
  return result
}

type lengths struct {
  length_1 int
  length_2 int
}

func shortest_journey(line_1, line_2, intersects []vector) []int {
  var result []int
  for _, i := range intersects {
    for line_1_index, line_1_value := range line_1 {
      if line_1_value == i {
        for line_2_index, line_2_value := range line_2 {
          if line_2_value == i {
            result = append(result, line_1_index + line_2_index)
          }
        }
      }
    }
  }
  return result
}

