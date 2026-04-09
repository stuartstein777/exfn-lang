package types

type Value interface {
	isValue() // Private method to "seal" the interface
}

type BoolValue bool

func (BoolValue) isValue() {}

type NilValue struct{}

func (NilValue) isValue() {}

type NumberValue float64

func (NumberValue) isValue() {}
