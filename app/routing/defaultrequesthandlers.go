package routing

import (
	"strconv"
)

func parseRequestIntParam(requestParam string) (int, error) {
	value, err := strconv.Atoi(requestParam)
	if err != nil {
		return 0, err
	}
	return value, err
}
