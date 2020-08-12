# From Nand to Tetris assembler implementation
This project provides an assembler implementation for the [From Nand to Tetris](https://www.nand2tetris.org/) course.
In that course you are building a computer - called Hack computer - from scratch with various tools provided by the course
like Hardware Simulator, CPU Emulator, VM Emulator, Assembler, etc.

## Usage
1. You need to have go installed

2. From the assembler directory

3. go build command will create the executable

4. ./assembler asm file name with path, without asm extension - Example: ./assembler testdata/test

5. The hack file with the same name will be created in the asm files directory

## About the assembler
The assembler translates **Hack assembly code** into 16bit length **Hack binary code**

### Assembly program elements:
- White space
  - Empty lines 
  - Line comments 
  - In-line comments
- Instructions
  - A-instructions
  - C-instructions
- Symbols
  - References
  - Label declarations
 
#### comments
everything what is after **//** is a comment, there are no block comments in 
**Hack assembly** language.

#### A-instructions
Symbolic syntax: @value

Where value is either
- a non-negative decimal constant or
- a symbol referring to such a constant

Examples: @21, @foo
Translation to binary:
• If value is a decimal constant, generate the equivalent binary constant
• If value is a symbol.

#### C-instructions
Symbolic syntax: dest=comp;jump

### More information
[From Nant to Tetris Chapter 6](https://b1391bd6-da3d-477d-8c01-38cdf774495a.filesusr.com/ugd/56440f_65a2d8eef0ed4e0ea2471030206269b5.pdf)
