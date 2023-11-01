package utils

func Duplicate2Darray[T int](data [][]T) [][]T {
	duplicate := make([][]T, len(data))
	for i := range data {
		duplicate[i] = make([]T, len(data[i]))
		copy(duplicate[i], data[i])
	}
	return duplicate
}
