package vm

import t "github.com/stuartstein777/exfnlang/types"

func Add(vm *VM) {
	if !t.IsNumber(vm.Peek(1)) || !t.IsNumber(vm.Peek(0)) {
		vm.RunTimeError("Operands must be numbers.")
		return
	}
	a := vm.Pop().(t.NumberValue)
	b := vm.Pop().(t.NumberValue)
	vm.Push(a + b)
}

func Subtract(vm *VM) {
	if !t.IsNumber(vm.Peek(1)) || !t.IsNumber(vm.Peek(0)) {
		vm.RunTimeError("Operands must be numbers.")
		return
	}
	b := vm.Pop().(t.NumberValue)
	a := vm.Pop().(t.NumberValue)
	vm.Push(a - b)
}

func Divide(vm *VM) {
	if !t.IsNumber(vm.Peek(1)) || !t.IsNumber(vm.Peek(0)) {
		vm.RunTimeError("Operands must be numbers.")
		return
	}
	b := vm.Pop().(t.NumberValue)
	a := vm.Pop().(t.NumberValue)
	vm.Push(a / b)
}

func Multiply(vm *VM) {
	if !t.IsNumber(vm.Peek(1)) || !t.IsNumber(vm.Peek(0)) {
		vm.RunTimeError("Operands must be numbers.")
		return
	}
	b := vm.Pop().(t.NumberValue)
	a := vm.Pop().(t.NumberValue)
	vm.Push(a * b)
}

func Negate(vm *VM) {
	if !t.IsNumber(vm.Peek(0)) {
		vm.RunTimeError("Operand must be a number.")
		return
	}
	v := vm.Pop().(t.NumberValue)
	vm.Push(-v)
}
