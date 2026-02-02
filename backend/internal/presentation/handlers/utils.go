package handlers

import "strconv"

// parseInt преобразует строку в int
func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

