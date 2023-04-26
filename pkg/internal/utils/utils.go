package utils

func EmptyOrValue[T any](i *T) T {
	if i == nil {
		return Default[T]()
	}
	return *i
}

func Default[T any]() T {
	var i T
	return i
}

func Ref[T any](i T) *T {
	return &i
}
