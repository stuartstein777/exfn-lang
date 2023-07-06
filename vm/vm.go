package vm

import (
	"encoding/binary"
)

type Value = float64

type Chunk struct {
	Code        []byte
	LineNumbers []int
	Constants   []Value
}

func WriteToChunk(chunk *Chunk, opCode byte, line int) {
	chunk.Code = append(chunk.Code, opCode)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
}

func WriteConstantToChunk(chunk *Chunk, opCode byte, constantIndex float32, line int) {
	codeAndConstant := []byte{opCode, byte(constantIndex)}
	chunk.Code = append(chunk.Code, codeAndConstant...)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
}

func AddConstant(chunk *Chunk, value Value, line int) int {
	chunk.Constants = append(chunk.Constants, value)
	chunk.LineNumbers = append(chunk.LineNumbers, line)
	return len(chunk.Constants) - 1
}

func WriteLongConstantToChunk(chunk *Chunk, opCode byte, constantIndex int, line int) {
	chunk.Code = append(chunk.Code, opCode)
	valueBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(valueBytes, uint32(constantIndex))
	chunk.Code = append(chunk.Code, valueBytes...)
	lines := []int{line, line, line, line}
	chunk.LineNumbers = append(chunk.LineNumbers, lines...)
}

/* =========== Debugging test methods - to remove */

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
