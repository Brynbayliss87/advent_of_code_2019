package intcode

// Opcode 1 adds together numbers read from two positions
// and stores the result in a third position.
// The three integers immediately after the opcode tell you these three positions
// the first two indicate the positions from which you should read the input values,
// and the third indicates the position at which the output should be stored.

// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
// instead of adding them. Again, the three integers after the opcode indicate
// where the inputs and outputs are, not their values.


// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter.
// For example, the instruction 3,50 would take an input value and store it at address 50.

// Opcode 4 outputs the value of its only parameter.
// For example, the instruction 4,50 would output the value at address 50.

// Opcode 99 halts the program


func Run(input, program int []int) []int {
  var pointer int

  for pointer < len(program) {
    instruction := program[pointer]
    mode1 := (instruction / 100) % 10
    mode2 := (instruction / 1000) % 10
    opcode := instruction % 100
    output := []int{}

    var param1 int
    var param2 int

    switch opcode {
    case 1:
      if mode1 = 1 {
        param1 = program[pointer + 1]
      } else {
        param1 = program[program[pointer + 1]]
      }

      if mode2 == 1 {
        param2 = program[pointer + 2]
      } else {
        param2 = program[program[pointer + 2]]
      }

      program[program[pointer + 3]] = param1 + param2
      pointer += 4

    case 2:
      if mode1 == 1 {
        param1 = program[pointer + 1]
      } else {
        param1 = program[program[pointer + 1]]
      }

      if mode2 == 1 {
        param2 = program[pointer + 2]
      } else {
        param2 =  program[program[pointer + 2]]
      }

      program[program[pointer + 3]] = param1 * param2
      pointer += 4

    case 3:
      program[program[pointer + 1]] = input
      pointer += 2

    case 4:
      if mode1 == 1 {
        param1 = program[pointer + 1]
      } else {
        param1 = program[program[pointer + 1]]
      }

      if mode2 == 1 {
        param2 = program[pointer + 2]
      } else {
        param2 =  program[program[pointer + 2]]
      }

      output = append(output, param1)
      pointer += 2

    case 99:
      return output
  }
}
