package test

type MockApi[T any] struct {
}

func (MockApi[T]) Get(resource string) (T, error) {
	var t T
	return t, nil
}

func (MockApi[T]) List(resource string) ([]T, error) {
	var list []T
	return list, nil
}

func (MockApi[T]) Post(resource string, payload any) (T, error) {
	var t T
	return t, nil
}
