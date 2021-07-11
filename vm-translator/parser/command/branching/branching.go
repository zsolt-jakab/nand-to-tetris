package branching

import (
	"fmt"
)

var operationMap = map[string]func(label string) []string{
	"label":   labelCommand(),
	"goto":    gotoCommand(),
	"if-goto": ifGotoCommand(),
}

func Translate(command string, label string) ([]string, error) {
	if translation, isPresent := operationMap[command]; isPresent {
		return translation(label), nil
	}
	return nil, fmt.Errorf("This is not a branching command : [%s] ", command)
}

func labelCommand() func(label string) []string {
	return func(label string) []string {
		return []string{
			"(" + label + ")",
		}
	}
}

func gotoCommand() func(label string) []string {
	return func(label string) []string {
		return []string{
			"@" + label,
			"0;JMP",
		}
	}
}

func ifGotoCommand() func(label string) []string {
	return func(label string) []string {
		return []string{
			"@SP",
			"AM=M-1",
			"D=M",
			"@" + label,
			"D;JNE",
		}
	}
}
