package main

import (
	vm "github.com/stuartstein777/exfnlang/vm"
)

func main() {
	chunk := vm.Chunk{
		Code:        []byte{},
		LineNumbers: []int{},
		Constants:   []vm.Value{},
	}

	vm.WriteToChunk(&chunk, vm.OP_RETURN, 123)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 0, 123)
	vm.AddConstant(&chunk, 456, 123)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 1, 123)
	vm.AddConstant(&chunk, 789, 123)
	vm.WriteToChunk(&chunk, vm.OP_RETURN, 124)
	vm.WriteNZeroConstants(&chunk, 1022)
	vm.WriteLongConstantToChunk(&chunk, vm.OP_CONSTANT_LONG, 1024, 125)
	vm.AddConstant(&chunk, 912, 125)
	vm.WriteToChunk(&chunk, vm.OP_RETURN, 126)

	//fmt.Printf("%v\n", chunk)
	vm.DisassembleChunk(chunk, "Test chunk")
}
