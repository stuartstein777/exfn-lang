package types

func IsNumber(value Value) bool {
	_, ok := value.(NumberValue)
	return ok
}

func IsBool(value Value) bool {
	_, ok := value.(BoolValue)
	return ok
}

func IsNil(value Value) bool {
	_, ok := value.(NilValue)
	return ok
}

func IsFalsey(value Value) bool {

	if IsNil(value) {
		return true
	}

	if IsBool(value) {
		_, res := value.(BoolValue)
		return !res
	}

	return false
}
