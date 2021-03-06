package memaccess

import "strconv"

const TempBaseAddress = 5

/*
A Temp command is providing the struct for temp segment commands
*/
type Temp struct {
	*Default
}

/*
NewTemp creates a new memory access command with a with a temp segment
*/
func NewTemp(base *Base) *Temp {
	spec := &Temp{}
	spec.Default = &Default{base, spec}
	return spec
}

func (command *Temp) push() []string {
	addr := command.calculateAddress()

	return []string{
		// store the *addr in D
		"@" + addr,
		"D=M",

		// store D to *SP
		"@SP",
		"A=M",
		"M=D",

		// SP++
		"@SP",
		"M=M+1",
	}
}

func (command *Temp) pop() []string {
	addr := command.calculateAddress()

	return []string{
		// pop stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		//select temp address and store D
		"@" + addr,
		"M=D",
	}
}

func (command *Temp) calculateAddress() string {
	i, _ := strconv.Atoi(command.i)
	return strconv.Itoa(TempBaseAddress + i)
}
