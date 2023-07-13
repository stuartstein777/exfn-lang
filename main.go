package main

import (
	"os"

	fe "github.com/stuartstein777/exfnlang/frontend"
	h "github.com/stuartstein777/exfnlang/helpers"
	vm "github.com/stuartstein777/exfnlang/vm"
)

func main() {
	//DebugTesting1()
	//DebugTesting2()

	if len(os.Args) > 1 {
		fe.RunFile(os.Args[1])
	} else {
		fe.Repl()
	}
}

func DebugTesting2() {
	chunk := vm.Chunk{
		Code:        []byte{},
		LineNumbers: []int{},
		Constants:   []vm.Value{},
	}

	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 0, 1)
	vm.AddConstant(&chunk, 3, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 1, 1)
	vm.AddConstant(&chunk, 2, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 2, 1)
	vm.AddConstant(&chunk, 1, 1)
	vm.WriteNZeroConstants(&chunk, 1022)
	vm.WriteLongConstantToChunk(&chunk, h.OP_CONSTANT_LONG, 1024, 2)
	vm.AddConstant(&chunk, 912, 2)
	vm.WriteToChunk(&chunk, h.OP_NEGATE, 3)
	vm.WriteToChunk(&chunk, h.OP_ADD, 2)
	vm.WriteToChunk(&chunk, h.OP_RETURN, 4)

	curVm := vm.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")

	vm.Run(&curVm)
}

func DebugTesting1() {
	chunk := vm.Chunk{
		Code:        []byte{},
		LineNumbers: []int{},
		Constants:   []vm.Value{},
	}

	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 0, 1)
	vm.AddConstant(&chunk, 3, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 1, 1)
	vm.AddConstant(&chunk, 2, 1)
	vm.WriteToChunk(&chunk, h.OP_MULTIPLY, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 2, 1)
	vm.AddConstant(&chunk, 4, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 3, 1)
	vm.AddConstant(&chunk, 5, 1)
	vm.WriteToChunk(&chunk, h.OP_DIVIDE, 1)
	vm.WriteToChunk(&chunk, h.OP_SUBTRACT, 1)
	vm.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 4, 1)
	vm.AddConstant(&chunk, 1, 1)
	vm.WriteToChunk(&chunk, h.OP_ADD, 1)
	vm.WriteToChunk(&chunk, h.OP_RETURN, 2)

	curVm := vm.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")

	vm.Run(&curVm)
}
