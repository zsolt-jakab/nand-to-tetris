package memaccess

import (
	"fmt"
	"strconv"
)

/*
Specification interface provides methods for different Memory access commands supporting the template method pattern
*/
type Specification interface {
	validate() error
	assembly() []string
	push() []string
	pop() []string
}

/*
A Base command is providing the main struct for all memory access commands
*/
type Base struct {
	stackOperation string
	segment        string
	i              string
}

/*
NewDefault creates a new memory access command with a stack operation(push/pop) using the ith local/argument/this/that segment as source or destination
*/
func NewBase(stackOperation string, segment string, i string) *Base {
	return &Base{stackOperation, segment, i}
}

/*
A Default command is providing the main struct for local, argument, this and that segment commands
*/
type Default struct {
	*Base
	specification Specification
}

var defaultSegmentMap = map[string]string{
	local:    localSegmentPointer,
	argument: argumentSegmentPointer,
	this:     thisSegmentPointer,
	that:     thatSegmentPointer,
}

/*
NewDefault creates a new memory access command with a stack operation(push/pop) using the ith local/argument/this/that segment as source or destination
*/
func NewDefault(base *Base) *Default {
	def := &Default{}
	def.Base = base
	def.specification = def
	return def
}

/*
Assembly the assembly translation of the VM local/argument/this/that memory segment command.
*/
func (command *Default) Translate() ([]string, error) {
	err := command.specification.validate()
	if err != nil {
		return nil, err
	}
	return command.specification.assembly(), nil

}

func (command *Default) validate() error {
	if isNotInt(command.i) || isNegativeInt(command.i) {
		return fmt.Errorf("Memory access command's i part must refer to a non-negative integer value but it was: [%s] ", command.i)
	} else if command.stackOperation != push && command.stackOperation != pop {
		return fmt.Errorf("Memory access command's stack stackOperation part possible values are push or pop but it was : [%s] ", command.stackOperation)
	}
	return nil
}

func (command *Default) assembly() []string {
	//in this point bc of validation only push or pop are possible values
	if command.stackOperation == push {
		return command.specification.push()
	} else {
		return command.specification.pop()
	}
}

/*func newDefault(stackOperation string, segment string, i string) *Default {
	return &Default{stackOperation, segment, i}
}*/

func (command *Default) push() []string {
	segmentAssemblyName := defaultSegmentMap[command.segment]
	return []string{
		// store the value of RAM[segment+i] in D
		"@" + command.i,
		"D=A",
		"@" + segmentAssemblyName,
		"A=D+M",
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

func (command *Default) pop() []string {
	segmentAssemblyName := defaultSegmentMap[command.segment]
	return []string{
		// select *segmentName + i(the address we want to pop to) and store it in D
		"@" + command.i,
		"D=A",
		"@" + segmentAssemblyName,
		"D=D+M",

		// store the address we want to pop to (RAM[addr_to_pop] = segmentName + i)
		"@addr_to_pop",
		"M=D",

		// pop and store it in D
		"@SP",
		"AM=M-1",
		"D=M",

		// select memory segmentName we want to pop to
		"@addr_to_pop",
		"A=M",

		// save top of SP into selected memory segmentName
		"M=D",
	}
}

func isNotInt(text string) bool {
	if _, err := strconv.Atoi(text); err != nil {
		return true
	}
	return false
}

func isNegativeInt(text string) bool {
	if value, err := strconv.Atoi(text); err == nil && value < 0 {
		return true
	}
	return false
}
