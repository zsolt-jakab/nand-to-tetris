package memaccess

/*
A Static command is providing the main struct for static segment commands
*/
type Static struct {
	*Default
	fileName string
}

/*
NewStatic creates a new memory access command with a stack
*/
func NewStatic(base *Base, fileName string) *Static {
	spec := &Static{}
	spec.Default = &Default{base, spec}
	spec.fileName = fileName
	return spec
}

func (command *Static) push() []string {
	return []string{
		// select static segment i and store D
		"@" + command.fileName + "." + command.i,
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

func (command *Static) pop() []string {
	return []string{
		// pop stack into D
		"@SP",
		"AM=M-1",
		"D=M",

		// select static segment i and store D
		"@" + command.fileName + "." + command.i,
		"M=D",
	}
}
