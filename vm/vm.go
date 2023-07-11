package vm

import (
	"encoding/binary"
)

const Debugging = true
const StackMax = 256

type Value = float64

// --------------- The VM -----------------//
type VM struct {
	Chunk *Chunk
	// Want the IP as a pointer so that we can increment it.
	// and deref straight into the chunk.
	IP       int
	Stack    [StackMax]Value
	StackPtr int
}

func (vm *VM) Push(value Value) {
	vm.Stack[vm.StackPtr] = value
	vm.StackPtr++
}

func (vm *VM) Pop() Value {
	vm.StackPtr--
	return vm.Stack[vm.StackPtr]
}

func (vm *VM) ResetStack() {
	vm.StackPtr = 0
}

func ReadNextByte(vm *VM) byte {
	opCode := vm.Chunk.Code[vm.IP]
	vm.IP++
	return opCode
}

func Run(vm *VM) int {
	for {
		if Debugging {
			disassembleInstruction(*vm.Chunk, vm.IP)
			printStack(vm)
		}
		opCode := ReadNextByte(vm)
		switch opCode {
		case OP_RETURN:
			return INTERPRET_OK
		case OP_CONSTANT:
			constantIndex := ReadNextByte(vm)
			value := vm.Chunk.Constants[int(constantIndex)]
			vm.Push(value)
		case OP_CONSTANT_LONG:
			constantIndex := binary.BigEndian.Uint32([]byte{ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm)})
			value := vm.Chunk.Constants[int(constantIndex)]
			vm.Push(value)
		case OP_NEGATE:
			vm.Push(-vm.Pop())
		}
	}
}

/*
OP_RETURN   -> return from a function.
OP_CONSTANT -> 2 bytes, first byte is opcode, 2nd byte is index of constant in chunk's constants array.
*/
const (
	// Single-character tokens.
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
	OP_NEGATE
	OP_ADD
	OP_SUBTRACT
)

/*
Possible results of running the VM with a chunk.
*/
const (
	INTERPRET_OK = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)
