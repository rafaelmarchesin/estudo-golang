package utils

func SetNoDuplicatedValue(s []string) []string {
	// Solução encontrada em: https://souvikhaldar.medium.com/removing-duplicate-elements-from-an-array-slice-in-golang-e30d93d7e66
	check := make(map[string]int)
	res := make([]string, 0)

	for _, val := range s {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}
