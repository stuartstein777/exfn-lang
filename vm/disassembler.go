package vm

import (
	"encoding/binary"
	"fmt"

	h "github.com/stuartstein777/exfnlang/helpers"
	t "github.com/stuartstein777/exfnlang/types"
)

func printStack(vm *VM) {
	fmt.Printf("          ")
	for i := 0; i < vm.StackPtr; i++ {
		fmt.Printf("[ %g ]", vm.Stack[i])
	}
	fmt.Printf("\n")
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, chunk t.Chunk, offset int) int {
	constantIdx := chunk.Code[offset+1]
	fmt.Printf("%-16s %4d '", name, constantIdx)
	fmt.Printf("%g", chunk.Constants[constantIdx])
	fmt.Printf("'\n")
	return offset + 2
}

func constantLongInstruction(name string, chunk t.Chunk, offset int) int {
	constantIdx := binary.BigEndian.Uint32(chunk.Code[offset+1 : offset+5])
	fmt.Printf("%-16s %4d '", name, constantIdx)
	fmt.Printf("%g", chunk.Constants[constantIdx])
	fmt.Printf("'\n")
	return offset + 5
}

func DisassembleChunk(chunk t.Chunk, name string) {
	fmt.Printf("== %s ==\n", name)
	fmt.Printf("Code:%v\n", chunk.Code)
	codeLength := len(chunk.Code)
	fmt.Printf("Code length: %d\n", codeLength)
	for offset := 0; offset < codeLength; {
		offset = disassembleInstruction(chunk, offset)
	}
	fmt.Printf("\n")
}

func disassembleInstruction(chunk t.Chunk, offset int) int {
	fmt.Printf("%04d ", offset)
	instruction := chunk.Code[offset]

	// print line information, if the line is the same as the previous
	// print pipe instead of line number

	if offset > 0 &&
		chunk.LineNumbers[offset] == chunk.LineNumbers[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", chunk.LineNumbers[offset])
	}

	switch instruction {
	case h.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case h.OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", chunk, offset)
	case h.OP_CONSTANT_LONG:
		return constantLongInstruction("OP_CONSTANT_LONG", chunk, offset)
	case h.OP_NEGATE:
		return simpleInstruction("OP_NEGATE", offset)
	case h.OP_ADD:
		return simpleInstruction("OP_ADD", offset)
	case h.OP_SUBTRACT:
		return simpleInstruction("OP_SUBTRACT", offset)
	case h.OP_MULTIPLY:
		return simpleInstruction("OP_MULTIPLY", offset)
	case h.OP_DIVIDE:
		return simpleInstruction("OP_DIVIDE", offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}
