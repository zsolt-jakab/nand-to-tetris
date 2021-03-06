package memaccess

import (
	"fmt"
)

const push = "push"
const pop = "pop"

const local = "local"
const argument = "argument"
const this = "this"
const that = "that"
const constant = "constant"
const static = "static"
const temp = "temp"
const pointer = "pointer"

const localSegmentPointer = "LCL"
const argumentSegmentPointer = "ARG"
const thisSegmentPointer = "THIS"
const thatSegmentPointer = "THAT"

func Translate(base *Base, fileName string) ([]string, error) {
	return getTranslation(base, fileName)
}

func getTranslation(base *Base, fileName string) ([]string, error) {
	switch base.segment {
	case local, argument, this, that:
		return NewDefault(base).Translate()
	case constant:
		return NewConstant(base).Translate()
	case static:
		return NewStatic(base, fileName).Translate()
	case pointer:
		return NewPointer(base).Translate()
	case temp:
		return NewTemp(base).Translate()
	default:
		return nil, fmt.Errorf("segment is not a valid memory access command with segment: [%s] ", base.segment)
	}
}
