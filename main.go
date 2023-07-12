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

	// vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 0, 1)
	// vm.AddConstant(&chunk, 3, 1)
	// vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 1, 1)
	// vm.AddConstant(&chunk, 2, 1)
	// vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 2, 1)
	// vm.AddConstant(&chunk, 1, 1)
	// vm.WriteNZeroConstants(&chunk, 1022)
	// vm.WriteLongConstantToChunk(&chunk, vm.OP_CONSTANT_LONG, 1024, 2)
	// vm.AddConstant(&chunk, 912, 2)
	// vm.WriteToChunk(&chunk, vm.OP_NEGATE, 3)
	// vm.WriteToChunk(&chunk, vm.OP_ADD, 2)
	// vm.WriteToChunk(&chunk, vm.OP_RETURN, 4)

	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 0, 1)
	vm.AddConstant(&chunk, 3, 1)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 1, 1)
	vm.AddConstant(&chunk, 2, 1)
	vm.WriteToChunk(&chunk, vm.OP_MULTIPLY, 1)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 2, 1)
	vm.AddConstant(&chunk, 4, 1)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 3, 1)
	vm.AddConstant(&chunk, 5, 1)
	vm.WriteToChunk(&chunk, vm.OP_DIVIDE, 1)
	vm.WriteToChunk(&chunk, vm.OP_SUBTRACT, 1)
	vm.WriteConstantToChunk(&chunk, vm.OP_CONSTANT, 4, 1)
	vm.AddConstant(&chunk, 1, 1)
	vm.WriteToChunk(&chunk, vm.OP_ADD, 1)
	vm.WriteToChunk(&chunk, vm.OP_RETURN, 2)

	curVm := vm.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	vm.Run(&curVm)

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")
}
