package pkg

func ArrChunk[T any](arr []T, size int) [][]T {
	chunks := make([][]T, 0)
	chunk := make([]T, 0, size)
	for i := 0; i < len(arr); i++ {
		chunk = append(chunk, arr[i])
		if len(chunk) >= size {
			chunks = append(chunks, chunk)
			chunk = make([]T, 0, size)
		}
	}
	return chunks
}
