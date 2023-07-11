package vm

/*
OP_RETURN        -> return from a function.
OP_CONSTANT      -> 2 bytes, first byte is opcode, 2nd byte is index of constant in chunk's constants array.
OP_CONSTANT_LONG -> 2 bytes, first byte is opcode, 2nd byte is index of constant in chunk's constants array.
OP_NEGATE        -> negate the top of the stack.
OP_ADD           -> add the top two values on the stack.
OP_SUBTRACT      -> subtract the top two values on the stack.
*/
const (
	// Single-character tokens.
	OP_RETURN byte = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
	OP_NEGATE
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
)

/*
Possible results of running the VM with a chunk.
*/
const (
	INTERPRET_OK = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)
