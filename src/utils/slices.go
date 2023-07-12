package utils

func MapSlice[Input any, Output any](input []Input, transformer func(Input) Output) []Output {
	var result []Output
	for _, item := range input {
		result = append(result, transformer(item))
	}

	return result
}
