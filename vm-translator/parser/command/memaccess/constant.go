package memaccess

import (
	"fmt"
)

/*
A Constant command is providing the struct constant segment commands
*/
type Constant struct {
	*Default
}

func (command *Constant) validate() error {
	if command.stackOperation != push {
		return fmt.Errorf("Constant memory access command's stack only valid stackOperation is push but it was : [%s] ", command.stackOperation)
	}
	return command.Default.validate()
}

/*
NewConstant creates a new memory access command with the push stack operation using the i param as value for push
*/
func NewConstant(base *Base) *Constant {
	spec := &Constant{}
	spec.Default = &Default{base, spec}
	return spec
}

func (command *Constant) push() []string {
	return []string{
		// store the value of i in D
		"@" + command.i,
		"D=A",

		// store D(i) to *SP
		"@SP",
		"A=M",
		"M=D",

		// SP++
		"@SP",
		"M=M+1",
	}
}
