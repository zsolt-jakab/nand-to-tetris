package instruction

import (
	"fmt"
)

const (
	opcodeC   = "111"
	emptyJump = ""
	emptyDest = ""
)

var compMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

var jumpMap = map[string]string{
	"JGT":     "001",
	"JEQ":     "010",
	"JGE":     "011",
	"JLT":     "100",
	"JNE":     "101",
	"JLE":     "110",
	"JMP":     "111",
	emptyJump: "000",
}

var destMap = map[string]string{
	"M":       "001",
	"D":       "010",
	"A":       "100",
	"MD":      "011",
	"AM":      "101",
	"AD":      "110",
	"AMD":     "111",
	emptyDest: "000",
}

/*
A C instruction contains a destination a computaion and a jump part
They are stored as they are digested, in a human readable form
*/
type C struct {
	dest   string
	comp   string
	jump   string
	binary string
}

/*
NewC function is the preferable way to create a C instruction, it will check if the parts of the instructions are
valid and raise an exception otherwise.
dest represents the destination register(s), it is optional(zero string is allowed)
jump represents the jump part of the instruction, it is optional(zero string is allowed)
comp represents the computation and it is mandatory
*/
func NewC(dest, comp, jump string) (*C, error) {
	if _, existsDest := destMap[dest]; !existsDest {
		return nil, fmt.Errorf("Can not create C instruction with unknown destination: [%s]", dest)
	} else if _, existsComp := compMap[comp]; !existsComp {
		return nil, fmt.Errorf("Can not create C instruction with unknown computation: [%s]", comp)
	} else if _, eexistsJump := jumpMap[jump]; !eexistsJump {
		return nil, fmt.Errorf("Can not create C instruction with unknown jump: [%s]", jump)
	}

	binary := createBinary(comp, dest, jump)
	return &C{dest, comp, jump, binary}, nil
}

/*
Binary is a function what will return the 16 bit long binary representation for the instruction as a string
Every instruction will have the same structure:
1 1 1 a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3
The first three bit means it is a C instruction
a c1 c2 c3 c4 c5 c6 bits are for the computation
d1 d2 d3 are the 3 destination bits
j1 j2 j3 are the 3 jump bits
*/
func (inst *C) Binary() string {
	return inst.binary
}

func createBinary(comp string, dest string, jump string) string {
	return opcodeC + compMap[comp] + destMap[dest] + jumpMap[jump]
}
