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

/*
TODO: Move instruction handlers to their own file.
if it gets too big.
*/
func Add(vm *VM) {
	b := vm.Pop()
	a := vm.Pop()
	vm.Push(a + b)
}

func Subtract(vm *VM) {
	b := vm.Pop()
	a := vm.Pop()
	vm.Push(a - b)
}

func Divide(vm *VM) {
	b := vm.Pop()
	a := vm.Pop()
	vm.Push(a / b)
}

func Multiply(vm *VM) {
	b := vm.Pop()
	a := vm.Pop()
	vm.Push(a * b)
}

func Negate(vm *VM) {
	vm.Stack[vm.StackPtr-1] = -vm.Stack[vm.StackPtr-1]
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
			Negate(vm)
		case OP_ADD:
			Add(vm)
		case OP_SUBTRACT:
			Subtract(vm)
		case OP_MULTIPLY:
			Multiply(vm)
		case OP_DIVIDE:
			Divide(vm)
		}
	}
}
