package utils

import (
	"github.com/forbole/bookkeeper/types"
)

// ConvertAttributeToMap turn attribute into a map so that it is easy to find attributes
// Make Denom as a key
func ConvertDenomToMap(array []types.Denom)map[string]types.Denom{
	resultingMap := map[string]types.Denom{}
	for _, m := range array {
		resultingMap[m.Denom] = m
	}
	return resultingMap
}