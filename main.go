package main

import (
	fe "github.com/stuartstein777/exfnlang/frontend"
	h "github.com/stuartstein777/exfnlang/helpers"
	t "github.com/stuartstein777/exfnlang/types"
	vmn "github.com/stuartstein777/exfnlang/vm"
)

func main() {
	//DebugTesting1()
	//DebugTesting2()
	fe.Repl()

	// if len(os.Args) > 1 {
	// 	fmt.Printf("Running file")
	// 	fe.RunFile(os.Args[1])
	// } else {
	// 	fe.Repl()
	// }
}

func DebugTesting2() {
	chunk := t.Chunk{
		Code:        []byte{},
		LineNumbers: []int{},
		Constants:   []t.Value{},
	}

	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 0, 1)
	t.AddConstant(&chunk, 3, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 1, 1)
	t.AddConstant(&chunk, 2, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 2, 1)
	t.AddConstant(&chunk, 1, 1)
	t.WriteNZeroConstants(&chunk, 1022)
	t.WriteLongConstantToChunk(&chunk, h.OP_CONSTANT_LONG, 1024, 2)
	t.AddConstant(&chunk, 912, 2)
	t.WriteToChunk(&chunk, h.OP_NEGATE, 3)
	t.WriteToChunk(&chunk, h.OP_ADD, 2)
	t.WriteToChunk(&chunk, h.OP_RETURN, 4)

	curVm := vmn.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")

	vmn.Run(&curVm)
}

func DebugTesting1() {
	chunk := t.Chunk{
		Code:        []byte{},
		LineNumbers: []int{},
		Constants:   []t.Value{},
	}

	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 0, 1)
	t.AddConstant(&chunk, 3, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 1, 1)
	t.AddConstant(&chunk, 2, 1)
	t.WriteToChunk(&chunk, h.OP_MULTIPLY, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 2, 1)
	t.AddConstant(&chunk, 4, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 3, 1)
	t.AddConstant(&chunk, 5, 1)
	t.WriteToChunk(&chunk, h.OP_DIVIDE, 1)
	t.WriteToChunk(&chunk, h.OP_SUBTRACT, 1)
	t.WriteConstantToChunk(&chunk, h.OP_CONSTANT, 4, 1)
	t.AddConstant(&chunk, 1, 1)
	t.WriteToChunk(&chunk, h.OP_ADD, 1)
	t.WriteToChunk(&chunk, h.OP_RETURN, 2)

	curVm := vmn.VM{
		Chunk: &chunk,
		IP:    0,
	}

	curVm.Chunk = &chunk

	//fmt.Printf("%v\n", chunk)
	//vm.DisassembleChunk(chunk, "Test chunk")

	vmn.Run(&curVm)
}
