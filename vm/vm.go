package vm

import (
	"encoding/binary"
	"fmt"

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

func (vm *VM) Peek(distance int) t.Value {
	var res = vm.Stack[distance]
	//fmt.Printf("Peeking %d, got %v\n", distance, res)
	return res
}

func (vm *VM) RunTimeError(message string) {
	// get the current chunk and line number, we advance before executing
	// so ip is already at the next instruction, so we need to look back one.
	chunk := vm.Chunk
	line := chunk.LineNumbers[vm.IP-1]
	fmt.Printf("[line %d] Runtime Error: %s\n", line, message)
	vm.ResetStack()
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
	//fmt.Printf("In vm.run()\n")

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
			if !t.IsNumber(vm.Peek(1)) {
				vm.RunTimeError("Operand must be a number.")
				return h.INTERPRET_RUNTIME_ERROR
			}
			Negate(vm)
		case h.OP_ADD:
			Add(vm)
		case h.OP_SUBTRACT:
			Subtract(vm)
		case h.OP_MULTIPLY:
			Multiply(vm)
		case h.OP_DIVIDE:
			Divide(vm)
		case h.OP_TRUE:
			vm.Push(t.BoolValue(true))
		case h.OP_FALSE:
			vm.Push(t.BoolValue(false))
		case h.OP_NIL:
			vm.Push(t.NilValue{})
		case h.OP_NOT:
			if !t.IsBool(vm.Peek(0)) {
				vm.RunTimeError("Operand must be a boolean or nil.")
				return h.INTERPRET_RUNTIME_ERROR
			}
			val := vm.Pop()
			b, _ := val.(t.BoolValue)
			if b == true {
				vm.Push(t.BoolValue(false))
			} else {
				vm.Push(t.BoolValue(true))
			}
		}
	}
}
