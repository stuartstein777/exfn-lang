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

	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 0, 1)
	vm.AddConstant(&chunk, 456, 1)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 1, 1)
	vm.AddConstant(&chunk, 789, 1)
	vm.WriteNZeroConstants(&chunk, 1022)
	vm.WriteLongConstantToChunk(&chunk, vm.OP_CONSTANT_LONG, 1024, 2)
	vm.AddConstant(&chunk, 912, 2)
	vm.WriteToChunk(&chunk, vm.OP_NEGATE, 3)
	vm.WriteToChunk(&chunk, vm.OP_RETURN, 4)

	curVm := vm.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	vm.Run(&curVm)

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")
}
