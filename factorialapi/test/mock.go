package test

type MockCollection[T any] struct {
}

func (c MockCollection[T]) FindOne(databaseName, collectionName, id string) (T, error) {
	var t T
	return t, nil
}

func (c MockCollection[T]) Find(databaseName, collectionName string, t T) ([]T, error) {
	var list []T
	return list, nil
}

func (c MockCollection[T]) List(databaseName, collectionName string) ([]T, error) {
	var list []T
	return list, nil
}

func (c MockCollection[T]) Add(databaseName, collectionName string, t T) (string, error) {
	return "", nil
}
