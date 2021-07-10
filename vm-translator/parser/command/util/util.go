package util

import "strconv"

func IdSequence() func() string {
	id := 0
	return func() string {
		id++
		return strconv.Itoa(id)
	}
}
