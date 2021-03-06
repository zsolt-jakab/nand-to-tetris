package memaccess

import "fmt"

const thisPointerIndex = "0"
const thatPointerIndex = "1"

var pointerMap = map[string]string{
	thisPointerIndex: thisSegmentPointer,
	thatPointerIndex: thatSegmentPointer,
}

type Pointer struct {
	*Default
}

/*
NewPointer creates a new memory access command with a stack
*/
func NewPointer(base *Base) *Pointer {
	spec := &Pointer{}
	spec.Default = &Default{base, spec}
	return spec
}

func (command *Pointer) validate() error {
	if isValidPointerIndex(command.i) {
		return command.Default.validate()
	}
	return fmt.Errorf("Pointer memory access command's valid indices are 0 and 1 but it was : [%s] ", command.i)
}

func isValidPointerIndex(index string) bool {
	_, isPresent := pointerMap[index]
	return isPresent
}

func (command *Pointer) push() []string {
	pointer := command.calculatePointer()
	return []string{
		// select pointer segment and store in D
		"@" + pointer,
		"D=M",

		// store D (RAM[segment+i]) in *SP
		"@SP",
		"A=M",
		"M=D",

		// SP++
		"@SP",
		"M=M+1",
	}
}

func (command *Pointer) pop() []string {
	pointer := command.calculatePointer()
	return []string{
		// pop stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		//select pointer and store D
		"@" + pointer,
		"M=D",
	}
}

func (command *Pointer) calculatePointer() string {
	return pointerMap[command.i]
}
