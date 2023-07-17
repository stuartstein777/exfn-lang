package vm

import (
	"encoding/binary"

	h "github.com/stuartstein777/exfnlang/helpers"
	t "github.com/stuartstein777/exfnlang/types"
)

const Debugging = true
const StackMax = 256

// --------------- The VM -----------------//
type VM struct {
	Chunk *t.Chunk
	/* TODO Want the IP as a pointer so that we can increment it.
	and deref straight into the chunk.*/
	IP       int
	Stack    [StackMax]t.Value
	StackPtr int
}

func (vm *VM) Push(value t.Value) {
	vm.Stack[vm.StackPtr] = value
	vm.StackPtr++
}

func (vm *VM) Pop() t.Value {
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
		case h.OP_RETURN:
			return h.INTERPRET_OK
		case h.OP_CONSTANT:
			constantIndex := ReadNextByte(vm)
			value := vm.Chunk.Constants[int(constantIndex)]
			vm.Push(value)
		case h.OP_CONSTANT_LONG:
			constantIndex := binary.BigEndian.Uint32([]byte{ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm)})
			value := vm.Chunk.Constants[int(constantIndex)]
			vm.Push(value)
		case h.OP_NEGATE:
			Negate(vm)
		case h.OP_ADD:
			Add(vm)
		case h.OP_SUBTRACT:
			Subtract(vm)
		case h.OP_MULTIPLY:
			Multiply(vm)
		case h.OP_DIVIDE:
			Divide(vm)
		}
	}
}
