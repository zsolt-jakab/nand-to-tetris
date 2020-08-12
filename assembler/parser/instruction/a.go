package instruction

import (
	"fmt"
	"strconv"
)

const (
	opcodeA          = "0"
	lastValidAddress = 32767
)

/*
An A instruction contains an address(15 bit long) for setting Address register value with it
It also has a leading zero, what is the operation code of an A instruction
In hack language an A isntruction looks like this: @value
@value will be translated to 0valueInBinary by the assembler
For example @21 -> 0(operation code) + 000000000010101(address) -> 0000000000010101(binary value of A instruction)
*/
type A struct {
	address string
	binary  string
}

/*
NewA function creates a new A instruction, based on an int decimal address value
It will store it as a 15 bit long binary value inside with the nessesarry leading zeros as an address
It will also create the binary representation of the A instructio which is the address and the leading opteration Code
*/
func NewA(adressDec int) (*A, error) {
	if adressDec < 0 {
		return nil, fmt.Errorf("Can not create A instruction with negative address value: [%d]", adressDec)
	} else if lastValidAddress < adressDec {
		return nil, fmt.Errorf("Can not create A instruction with bigger than "+strconv.Itoa(lastValidAddress)+" address value: [%d]", adressDec)
	}

	address := fmt.Sprintf("%015b", adressDec)
	return &A{address, opcodeA + address}, nil
}

/*
Binary returns a binary representation value of an A instruction.
It is 16 bit long binary value, always an additional leading zero for the addres.
That leading zero means for the hack computer that it is an A instruction.
*/
func (inst *A) Binary() string {
	return inst.binary
}
