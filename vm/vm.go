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
			vm.Stack[vm.StackPtr] = value
			vm.StackPtr++
		case OP_CONSTANT_LONG:
			constantIndex := binary.BigEndian.Uint32([]byte{ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm), ReadNextByte(vm)})
			value := vm.Chunk.Constants[int(constantIndex)]
			vm.Stack[vm.StackPtr] = value
			vm.StackPtr++
		}
	}
}

// --------------- Chunks -----------------//
// TODO: Move this chunk.go
type Chunk struct {
	Code        []byte
	LineNumbers []int
	Constants   []Value
}

// TODO: Move this chunk.go
func WriteToChunk(chunk *Chunk, opCode byte, line int) {
	chunk.Code = append(chunk.Code, opCode)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
}

// TODO: Move this chunk.go
func WriteConstantToChunk(chunk *Chunk, opCode byte, constantIndex float32, line int) {
	codeAndConstant := []byte{opCode, byte(constantIndex)}
	chunk.Code = append(chunk.Code, codeAndConstant...)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
}

// TODO: Move this chunk.go
func AddConstant(chunk *Chunk, value Value, line int) int {
	chunk.Constants = append(chunk.Constants, value)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
	return len(chunk.Constants) - 1
}

// TODO: Move this chunk.go
func WriteLongConstantToChunk(chunk *Chunk, opCode byte, constantIndex int, line int) {
	chunk.Code = append(chunk.Code, opCode)
	valueBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(valueBytes, uint32(constantIndex))
	chunk.Code = append(chunk.Code, valueBytes...)
	lines := []int{line, line, line, line}
	chunk.LineNumbers = append(chunk.LineNumbers, lines...)
}

/* =========== Debugging test methods - to remove */
//TODO: Move this chunk.go
func WriteNZeroConstants(chunk *Chunk, n int) {
	for i := 0; i < n; i++ {
		chunk.Constants = append(chunk.Constants, 0)
	}
}

/* =========== Debugging test methods - to remove */

/*
OP_RETURN   -> return from a function.
OP_CONSTANT -> 2 bytes, first byte is opcode, 2nd byte is index of constant in chunk's constants array.
*/
const (
	// Single-character tokens.
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
)

const (
	INTERPRET_OK = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)
