package helpers

import (
	"strconv"
	"strings"
)

func FormatIdFromUintToString(id uint) string {
	stringId := strconv.Itoa(int(id))
	return strings.Repeat("0", 6-len(stringId)) + stringId
}

func FormatIdFromStringToUint(id string) (uint, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(intId), nil
}
