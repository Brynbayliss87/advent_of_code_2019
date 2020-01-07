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


func Run(input, program int []int) int {
  var pointer = int

  for pointer < len(program) {
    opcode := program[pointer]
  }
}
